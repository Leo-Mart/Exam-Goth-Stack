package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func Open() *sql.DB {
	connStr := os.Getenv("DB_CONN_STRING")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price NUMERIC(6,2) NOT NULL,
  available BOOLEAN,
  created timestamp DEFAULT NOW()
  )`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available)
  VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}

	return pk
}
