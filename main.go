package main

import (
	"email-microservice/email"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env")
	}
	router := mux.NewRouter()
	router.HandleFunc("/email", createMessage).Methods("POST")
	http.ListenAndServe(":8000", router)
}

type Mail struct {
	Sender string `json:"sender"`
	Body   string `json:"body"`
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mail Mail
	_ = json.NewDecoder(r.Body).Decode(&mail)
	email.Send(mail.Sender, mail.Body)
	json.NewEncoder(w).Encode(mail)

}
