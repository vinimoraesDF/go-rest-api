package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"fmt"
	"github.com/gorilla/mux"
)

func main() {


	httpPort := 8081

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/test", Test)
	router.HandleFunc("/funci/{name}", Hola)
	router.HandleFunc("/info", Info)

	fmt.Printf("listening on %v\n", httpPort)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(router))
	if err != nil {
		log.Fatal(err)
	}
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HipChatResponse{Area: "Developer Support", Identificacao: "Guru Joao", Matricula: "Z123456", Funcao: "Especialista"}
	json.NewEncoder(w).Encode(response)
}

func Hola(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Header().Set("Content-Type", "application/json")
	response := HipChatResponse{Area: "Developer Environment", Identificacao: "Guri " + name, Matricula: "X123456", Funcao: "Trainee"}
	json.NewEncoder(w).Encode(response)
}

type HipChatResponse struct {
	Area         string `json:"area"`
	Identificacao       string `json:"identificacao"`
	Matricula        string `json:"matricula"`
	Funcao string `json:"funcao"`
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go API</h1><div>Welcome to whereever you are</div>")
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
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
