package dbstore

import (
	"context"
	"fmt"
	"log"

	"github.com/Leo-Mart/goth-test/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (cs *characterStore) AddCharacter(char models.Character) (primitive.ObjectID, error) {
	coll := cs.db.Database("goth-exam").Collection("characters")

	inserted, err := coll.InsertOne(context.TODO(), char)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("could not insert character into db: %v", err)
	}

	fmt.Printf("Added new character %v \n", inserted.InsertedID)
	return inserted.InsertedID.(primitive.ObjectID), nil
}

func (cs *characterStore) UpdateCharacter(id primitive.ObjectID, character models.Character) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"characterprofile": character.CharacterProfile, "keystoneprofile": character.KeystoneProfile, "gear": character.Gear, "media": character.Media}}

	coll := cs.db.Database("goth-exam").Collection("characters")
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("could not update character: %v", err)
	}
	fmt.Printf("Results of put: %v", result)
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

func (cs *characterStore) GetCharacterByID(id primitive.ObjectID) (models.Character, error) {
	var foundCharacter models.Character

	filter := bson.M{"_id": id}
	coll := cs.db.Database("goth-exam").Collection("characters")

	err := coll.FindOne(context.TODO(), filter).Decode(&foundCharacter)
	if err != nil {
		return models.Character{}, fmt.Errorf("could not find character with id: %s. Reason: %v", id, err)
	}
	return foundCharacter, nil
}

func (cs *characterStore) DeleteCharacterById(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	coll := cs.db.Database("goth-exam").Collection("characters")

	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("could not find character with id: %s. Reason: %v", id, err)
	}
	return nil
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
