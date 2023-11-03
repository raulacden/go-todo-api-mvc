package model

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "123456"
	dbname   = "todo-api-mvc"
)

// Dejando la variable en minuscula, solo permitimos el acceso a los objetos dentro de model
var conn *pgx.Conn

func Connect() *pgx.Conn {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := pgx.Connect(context.Background(), psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	conn = db
	return conn
}
