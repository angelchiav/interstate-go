package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/angelchiav/interstate-go/internal/database"
)

type apiConfig struct {
	database *database.Queries
	env      string
}

func main() {
	dbURL := os.Getenv("DB_URL")
	fmt.Println("DB_URL in runtime: ", dbURL)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("error opening the database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("error with ping in the database (check DB_URL in .env) ", err)
	}

	dbQueries := database.New(db)

	mux := http.NewServeMux()

	apiCfg := &apiConfig{
		database: dbQueries,
		env:	os.Getenv("PORT")
	}
}
