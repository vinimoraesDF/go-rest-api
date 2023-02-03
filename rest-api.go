package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", Test)
	router.HandleFunc("/funci/{name}", Hola)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func Test(w http.ResponseWriter, r *http.Request) {
	response := HipChatResponse{Area: "Developer Support", Identificacao: "Guru Joao", Matricula: "Z123456", Funcao: "Especialista"}
	json.NewEncoder(w).Encode(response)
}

func Hola(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	response := HipChatResponse{Area: "Developer Environment", Identificacao: "Guri " + name, Matricula: "X123456", Funcao: "Trainee"}
	json.NewEncoder(w).Encode(response)
}

type HipChatResponse struct {
	Area         string `json:"area"`
	Identificacao       string `json:"identificacao"`
	Matricula        string `json:"matricula"`
	Funcao string `json:"funcao"`
}

type HipChatrequest struct {
	Event string `json:"event"`
	Item  struct {
		Message struct {
			Date time.Time `json:"date"`
			From struct {
				ID          int    `json:"id"`
				MentionName string `json:"mention_name"`
				Name        string `json:"name"`
			} `json:"from"`
			ID       string        `json:"id"`
			Mentions []interface{} `json:"mentions"`
			Message  string        `json:"message"`
			Type     string        `json:"type"`
		} `json:"message"`
		Room struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"room"`
	} `json:"item"`
	WebhookID int `json:"webhook_id"`
}
