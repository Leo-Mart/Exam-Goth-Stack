package dbstore

import (
	"context"
	"fmt"
	"log"

	"github.com/Leo-Mart/goth-test/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type characterStore struct {
	logger *log.Logger
	db     *mongo.Client
}

func NewCharacterStore(logger *log.Logger, db *mongo.Client) *characterStore {
	return &characterStore{
		logger: logger,
		db:     db,
	}
}

func (cs *characterStore) AddCharacter(char models.Character) error {
	coll := cs.db.Database("goth-exam").Collection("characters")

	inserted, err := coll.InsertOne(context.TODO(), char)
	if err != nil {
		return fmt.Errorf("could not insert character into db: %v", err)
	}

	fmt.Printf("Added new character %v \n", inserted.InsertedID)
	return nil
}

func (cs *characterStore) GetCharacters() ([]models.Character, error) {
	coll := cs.db.Database("goth-exam").Collection("characters")

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("could not fetch characters: %v", err)
	}

	results := []models.Character{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, fmt.Errorf("could not decode? characters: %v", err)
	}

	return results, nil
}

func (cs *characterStore) GetCharacterByID(ID string) (models.Character, error) {
	// TODO: make this do stuff.
	return models.Character{}, nil
}

func (cs *characterStore) GetCharacterByName(charName string) (models.Character, error) {
	var result models.Character

	filter := bson.D{{"characterprofile.name", charName}}
	coll := cs.db.Database("goth-exam").Collection("characters")
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return models.Character{}, fmt.Errorf("could not find character with name: %v", err)
	}
	return result, nil
}

func (cs *characterStore) DeleteCharacterByName(charName string) error {
	filter := bson.D{{"characterprofile.name", charName}}

	coll := cs.db.Database("goth-exam").Collection("characters")
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error when deleting character with name: %s: %v", charName, err)
	}

	return nil
}
