package models

type Character struct {
	ID               string
	CharacterProfile CharacterProfile
	KeystoneProfile  KeystoneProfile
	Gear             CharacterGear
	Media            CharacterMedia
}
