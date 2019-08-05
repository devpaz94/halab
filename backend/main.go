package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/halab/backend/handlers"
	"github.com/halab/backend/models"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "paz"
	password = "password"
	dbname   = "postgres"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	models.InitDB(psqlInfo)

	handler.HandleRequests()
	log.Fatal(http.ListenAndServe(":8881", nil))
}
