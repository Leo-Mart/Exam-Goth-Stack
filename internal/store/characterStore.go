package store

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Leo-Mart/goth-test/internal/models"
)

type characterStore struct {
	logger     *log.Logger
	characters map[string]models.BlizzardCharacter
}

func NewCharacterStore(logger *log.Logger) *characterStore {
	return &characterStore{
		logger:     logger,
		characters: make(map[string]models.BlizzardCharacter),
	}
}

func (cs *characterStore) AddCharacter(char models.BlizzardCharacter) error {
	if char.ID == "" {
		return fmt.Errorf("ID is required")
	}

	if _, ok := cs.characters[char.ID]; ok {
		return fmt.Errorf("a character with the ID: %s already exists", char.ID)
	}

	cs.characters[char.ID] = char
	fmt.Printf("Added new character %v \n", char.CharacterProfile.Name)
	return nil
}

func (cs *characterStore) GetCharacters() ([]models.BlizzardCharacter, error) {
	if cs.characters == nil {
		return nil, fmt.Errorf("no characters found")
	}

	characters := make([]models.BlizzardCharacter, 0, len(cs.characters))
	for _, char := range cs.characters {
		characters = append(characters, char)
	}
	return characters, nil
}

func (cs *characterStore) GetCharacterByID(ID string) (models.BlizzardCharacter, error) {
	character := cs.characters[ID]
	if character.ID == "" {
		cs.logger.Printf("Could not find character with ID: %s", ID)
		return models.BlizzardCharacter{}, fmt.Errorf("could not find character")
	}
	return character, nil
}

func (cs *characterStore) GetCharacterByName(charName string) (models.BlizzardCharacter, error) {
	for _, val := range cs.characters {
		if val.CharacterProfile.Name == charName {
			return val, nil
		}
	}
	return models.BlizzardCharacter{}, fmt.Errorf("could not find a character with name: %s", charName)
}

func DecodeCharacter(payload []byte) (models.RaiderioCharacter, error) {
	var char models.RaiderioCharacter
	if err := json.Unmarshal(payload, &char); err != nil {
		return models.RaiderioCharacter{}, fmt.Errorf("error decoding characters: %v", err)
	}
	return char, nil
}
