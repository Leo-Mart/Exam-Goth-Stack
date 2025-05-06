package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Leo-Mart/goth-test/internal/models"
	"github.com/Leo-Mart/goth-test/templates"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CharacterStore interface {
	AddCharacter(character models.Character) (primitive.ObjectID, error)
	UpdateCharacter(charId primitive.ObjectID, character models.Character) error
	GetCharacters() ([]models.Character, error)
	GetCharacterByID(charId primitive.ObjectID) (models.Character, error)
	GetCharacterByName(charName string) (models.Character, error)
	DeleteCharacterById(charId primitive.ObjectID) error
}

type server struct {
	logger      *log.Logger
	port        int
	httpServer  *http.Server
	characterDb CharacterStore
}

func NewServer(logger *log.Logger, port int, characterDb CharacterStore) (*server, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is required")
	}
	if characterDb == nil {
		return nil, fmt.Errorf("characterDb is required")
	}
	return &server{
		logger:      logger,
		port:        port,
		characterDb: characterDb,
	}, nil
}

func (s *server) Start() error {
	s.logger.Printf("Starting server on port %d", s.port)
	var stopChan chan os.Signal

	// define router
	r := http.NewServeMux()
	// serving static files
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	// routes
	r.HandleFunc("GET /", s.homeHandler)
	r.HandleFunc("GET /import", s.importPageHandler)
	r.HandleFunc("GET /import-home", s.importHomeHandler)
	r.HandleFunc("GET /about", s.aboutHandler)
	r.HandleFunc("GET /characters", s.getCharactersHandler)
	r.HandleFunc("GET /{id}", s.getCharacterDetailsHandler)
	r.HandleFunc("PUT /update/{id}", s.updateCharacterHandler)
	r.HandleFunc("POST /character/add", s.addCharacterHandler)
	r.HandleFunc("DELETE /delete/{id}", s.deleteCharactersHandler)

	// define server
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: r,
	}

	// create a channel to listen for signals
	stopChan = make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error when running server: %s", err)
		}
	}()

	<-stopChan

	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		log.Fatalf("error when shutting down server: %v", err)
		return err
	}

	return nil
}

func (s *server) homeHandler(w http.ResponseWriter, r *http.Request) {
	homeTemplate := templates.Home()
	err := templates.Layout(homeTemplate, "WoW Tracker", "/").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (s *server) importPageHandler(w http.ResponseWriter, r *http.Request) {
	newCharTemplate := templates.ImportNewCharacter()
	err := templates.Layout(newCharTemplate, "Import new character", "/import").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (s *server) importHomeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ImportNewCharacterFromHome().Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering tempalte", http.StatusInternalServerError)
		return
	}
}

func (s *server) getKeystoneProfile(charProfile models.CharacterProfile, accessToken string) models.KeystoneProfile {
	profileRequestURL := fmt.Sprintf("%s&locale=en_GB", charProfile.KeystoneProfileURL.URL)
	body := strings.NewReader("")

	request, err := http.NewRequest(http.MethodGet, profileRequestURL, body)
	if err != nil {
		s.logger.Printf("could not create request: %v", err)
		return models.KeystoneProfile{}

	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	httpClient := new(http.Client)
	resp, err := httpClient.Do(request)
	if err != nil {
		s.logger.Printf("could not make request: %v", err)
		return models.KeystoneProfile{}
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Printf("could not read response body: %v", err)
		return models.KeystoneProfile{}
	}
	var keystoneProfile models.KeystoneProfile
	err = json.Unmarshal(respData, &keystoneProfile)
	if err != nil {
		s.logger.Printf("Could not unmarshal json: %v", err)
		return models.KeystoneProfile{}
	}

	return keystoneProfile
}

func (s *server) getGear(charProfile models.CharacterProfile, accessToken string) models.CharacterGear {
	requestURL := fmt.Sprintf("%s&locale=en_GB", charProfile.EquipmentURL.URL)
	body := strings.NewReader("")

	request, err := http.NewRequest(http.MethodGet, requestURL, body)
	if err != nil {
		s.logger.Printf("Could not create request: %v", err)
		return models.CharacterGear{}
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	httpClient := new(http.Client)
	resp, err := httpClient.Do(request)
	if err != nil {
		s.logger.Printf("Could not make request: %v", err)
		return models.CharacterGear{}
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Printf("Could not read response body: %v", err)
		return models.CharacterGear{}
	}
	var characterGear models.CharacterGear
	err = json.Unmarshal(respData, &characterGear)
	if err != nil {
		s.logger.Printf("Could not unmarshal json: %v", err)
		return models.CharacterGear{}
	}

	return characterGear
}

func (s *server) addCharacterHandler(w http.ResponseWriter, r *http.Request) {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	realm := strings.ToLower(r.FormValue("realm"))
	name := strings.ToLower(r.FormValue("name"))

	accessToken := s.getAPIToken(clientID, clientSecret, "eu")

	charProfile := s.getCharacterProfile(realm, name, accessToken)
	characterGear := s.getGear(charProfile, accessToken)
	keystoneProfile := s.getKeystoneProfile(charProfile, accessToken)
	characterMedia := s.getCharacterMedia(charProfile, accessToken)

	var character models.Character
	character.CharacterProfile = charProfile
	character.KeystoneProfile = keystoneProfile
	character.Gear = characterGear
	character.Media = characterMedia

	oid, err := s.characterDb.AddCharacter(character)
	if err != nil {
		s.logger.Printf("could not add character to DB: %v", err)
	}

	importedChar, err := s.characterDb.GetCharacterByID(oid)
	if err != nil {
		s.logger.Printf("error while fetching characters from DB")
		return
	}

	charCardTemplate := templates.CharacterCard(importedChar)
	err = charCardTemplate.Render(r.Context(), w)
	if err != nil {
		s.logger.Printf("Error when getting character-list: %v", err)
		return
	}
}

func (s *server) getCharacterProfile(characterRealm string, characterName string, accessToken string) models.CharacterProfile {
	requestURL := fmt.Sprintf("https://%s.api.blizzard.com/profile/wow/character/%s/%s?namespace=profile-eu&locale=en_GB", "eu", characterRealm, characterName)
	body := strings.NewReader("")

	request, err := http.NewRequest(http.MethodGet, requestURL, body)
	if err != nil {
		s.logger.Printf("Could not create request: %v", err)
		return models.CharacterProfile{}
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	httpClient := new(http.Client)
	resp, err := httpClient.Do(request)
	if err != nil {
		s.logger.Printf("Could not make request: %v", err)
		return models.CharacterProfile{}
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Printf("Could not read response body: %v", err)
		return models.CharacterProfile{}
	}
	var charProfile models.CharacterProfile
	err = json.Unmarshal(respData, &charProfile)
	if err != nil {
		s.logger.Printf("Could not unmarshal json: %v", err)
		return models.CharacterProfile{}
	}

	return charProfile
}

func (s *server) getCharacterMedia(charProfile models.CharacterProfile, accessToken string) models.CharacterMedia {
	requestURL := fmt.Sprintf("%s&locale=en_GB", charProfile.MediaURL.URL)
	body := strings.NewReader("")
	request, err := http.NewRequest(http.MethodGet, requestURL, body)
	if err != nil {
		s.logger.Printf("Could not create request :%v", err)
		return models.CharacterMedia{}
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	httpClient := new(http.Client)
	resp, err := httpClient.Do(request)
	if err != nil {
		s.logger.Printf("could not make request: %v", err)
		return models.CharacterMedia{}
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Printf("could not read response body: %v", err)
		return models.CharacterMedia{}
	}
	var characterMedia models.CharacterMedia
	err = json.Unmarshal(respData, &characterMedia)
	if err != nil {
		s.logger.Printf("error while unmarshalling json: %v", err)
		return models.CharacterMedia{}
	}

	return characterMedia
}

func (s *server) getAPIToken(clientID string, clientSecret string, region string) string {
	requestURL := fmt.Sprintf("https://%s.battle.net/oauth/token", region)
	body := strings.NewReader("grant_type=client_credentials")

	request, err := http.NewRequest(http.MethodPost, requestURL, body)
	if err != nil {
		s.logger.Printf("could not create request: %v", err)
		return ""
	}

	request.SetBasicAuth(clientID, clientSecret)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := new(http.Client)
	response, err := httpClient.Do(request)
	if err != nil {
		s.logger.Printf("Could not make request: %s", err)
		return ""
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		s.logger.Printf("Error while reading response body: %v", err)
		return ""
	}

	var tokenData models.Token
	err = json.Unmarshal(responseData, &tokenData)
	if err != nil {
		s.logger.Printf("Could not unmarshal token json: %v", err)
		return ""
	}

	return tokenData.AccessToken
}

func (s *server) getCharactersHandler(w http.ResponseWriter, r *http.Request) {
	characters, err := s.characterDb.GetCharacters()
	if err != nil {
		s.logger.Printf("Error when getting characters: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	characterListTemplate := templates.CharacterList(characters)
	err = templates.Layout(characterListTemplate, "Imported Characters", "/characters").Render(r.Context(), w)
	if err != nil {
		s.logger.Printf("Error when rendering character-list")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) updateCharacterHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		s.logger.Printf("error converting id to hex: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	oldCharacter, err := s.characterDb.GetCharacterByID(oid)
	if err != nil {
		s.logger.Printf("error fetching character from db: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	name := strings.ToLower(oldCharacter.CharacterProfile.Name)
	realm := strings.ToLower("shadowsong")

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	accessToken := s.getAPIToken(clientID, clientSecret, "eu")

	charProfile := s.getCharacterProfile(realm, name, accessToken)
	characterGear := s.getGear(charProfile, accessToken)
	keystoneProfile := s.getKeystoneProfile(charProfile, accessToken)
	characterMedia := s.getCharacterMedia(charProfile, accessToken)

	var character models.Character
	character.CharacterProfile = charProfile
	character.KeystoneProfile = keystoneProfile
	character.Gear = characterGear
	character.Media = characterMedia

	err = s.characterDb.UpdateCharacter(oid, character)
	if err != nil {
		s.logger.Printf("error when updating character: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	character.ID = primitive.ObjectID(oid)
	w.WriteHeader(http.StatusOK)
	characterDetailsTemplate := templates.CharacterDetails(character)
	err = characterDetailsTemplate.Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func (s *server) deleteCharactersHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		s.logger.Printf("error converting objectID: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.characterDb.DeleteCharacterById(oid)
	if err != nil {
		s.logger.Printf("error when deleting character: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *server) getCharacterDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		s.logger.Printf("error when converting id: %v", err)
		return
	}

	character, err := s.characterDb.GetCharacterByID(oid)
	if err != nil {
		s.logger.Printf("error when fetching character: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	characterDetailsTemplate := templates.CharacterDetails(character)
	err = templates.Layout(characterDetailsTemplate, "Details page", "").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func (s *server) aboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutTemplate := templates.About()
	err := templates.Layout(aboutTemplate, "About page", "/about").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error when rendering template", http.StatusInternalServerError)
	}
}
