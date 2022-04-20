package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/kienit/be_brankas_test/database"
	"github.com/kienit/be_brankas_test/upload"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Create a new token object, specifying signing method and the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": "kiennguyen"})
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN")))

		// Using template to generate HTML
		// Get string from environment variables to set for token
		t := template.New("formUpload.gotmpl")
		t, err = t.ParseFiles(path.Join("templates", "formUpload.gotmpl"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		templateData := struct {
			UploadPath string
			Token      string
		}{os.Getenv("UploadPath"), tokenString}

		t.Execute(w, templateData)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	// Setup fix schema
	os.Setenv("schema", "brankas_test")
	os.Setenv("Token", "TestingToken")
	os.Setenv("UploadPath", "/upload")

	database.SetupDatabase()
	http.HandleFunc("/", rootHandler)
	upload.SetupRoutes()
	http.ListenAndServe(":5000", nil)
}
