package db

import (
	"context"
	"os"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func Open() *mongo.Client {
	connString := os.Getenv("MONGO_DB_CONN_STRING")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connString))
	if err != nil {
		panic(err)
	}
	return client
}
