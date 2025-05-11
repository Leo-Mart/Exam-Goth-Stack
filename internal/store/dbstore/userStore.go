package dbstore

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Leo-Mart/goth-test/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userStore struct {
	logger *log.Logger
	db     *mongo.Client
}

func NewUserStore(logger *log.Logger, db *mongo.Client) *userStore {
	return &userStore{
		logger: logger,
		db:     db,
	}
}

func (us *userStore) GetUserByEmail(email string) (models.User, error) {
	var u models.User

	email = strings.ToLower(email)
	filter := bson.D{{"email", email}}
	coll := us.db.Database("goth-exam").Collection("users")
	err := coll.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return models.User{}, fmt.Errorf("error when fetching user by email: %v", err)
	}
	return u, nil
}

func (us *userStore) Authenticate(email, testPassword string) (string, string, error) {
	var id primitive.ObjectID
	var hashedPassword string

	user, err := us.GetUserByEmail(email)
	if err != nil {
		us.logger.Printf("could  not fetch user from db: %v", err)
		return primitive.ObjectID.Hex(id), "", err
	}

	id = user.ID
	hashedPassword = user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		us.logger.Printf("invalid password")
		return "", "", errors.New("incorrect password")
	} else if err != nil {
		return "", "", err
	}

	return primitive.ObjectID.Hex(id), hashedPassword, nil
}
