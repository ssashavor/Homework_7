package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Host        string   `json:"host"`
	User_agent  string   `json:"user_agent"`
	Request_uri string   `json:"request_uri"`
	Headers     *Headers `json:"headers"`
}

type Headers struct {
	User_agent string `json:"User_agent"`
	Accept     string `json:"Accept"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	user := User{
		r.Host,
		r.Header.Get("User-agent"),
		r.URL.Path,
		&Headers{r.Header.Get("User-agent"),
			r.Header.Get("Accept")},
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", Handler)
	fmt.Println("Running server")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
