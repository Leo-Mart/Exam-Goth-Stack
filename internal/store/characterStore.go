package store

import (
	"fmt"
	"log"

	"github.com/Leo-Mart/goth-test/internal/models"
)

type characterStore struct {
	logger     *log.Logger
	characters map[string]models.Character
}

func NewCharacterStore(logger *log.Logger) *characterStore {
	return &characterStore{
		logger:     logger,
		characters: make(map[string]models.Character),
	}
}

func (cs *characterStore) AddCharacter(char models.Character) error {
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

func (cs *characterStore) GetCharacters() ([]models.Character, error) {
	if cs.characters == nil {
		return nil, fmt.Errorf("no characters found")
	}

	characters := make([]models.Character, 0, len(cs.characters))
	for _, char := range cs.characters {
		characters = append(characters, char)
	}
	return characters, nil
}

func (cs *characterStore) GetCharacterByID(ID string) (models.Character, error) {
	character := cs.characters[ID]
	if character.ID == "" {
		cs.logger.Printf("Could not find character with ID: %s", ID)
		return models.Character{}, fmt.Errorf("could not find character")
	}
	return character, nil
}

func (cs *characterStore) GetCharacterByName(charName string) (models.Character, error) {
	for _, val := range cs.characters {
		if val.CharacterProfile.Name == charName {
			return val, nil
		}
	}
	return models.Character{}, fmt.Errorf("could not find a character with name: %s", charName)
}
