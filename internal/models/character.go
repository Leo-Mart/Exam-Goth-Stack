package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CharacterProfile CharacterProfile
	KeystoneProfile  KeystoneProfile
	Specializations  Specialization
	Gear             CharacterGear
	Media            CharacterMedia
}
