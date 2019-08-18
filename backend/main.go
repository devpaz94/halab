package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/halab/backend/handlers"
	"github.com/halab/backend/models"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://locahost:8881",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string("https://www.googleapis.com/auth/userinfo.email"),
		Endpoint:     google.Endpoint,
	}
)

const (
	host     = "localhost"
	port     = 5432
	user     = "paz"
	password = "password"
	dbname   = "halab"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	models.InitDB(psqlInfo)

	handler.HandleRequests()
	log.Fatal(http.ListenAndServe(":8881", nil))
}
