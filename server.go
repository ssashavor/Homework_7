package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Host        string   `json:"host"`
	UserAgent  string   `json:"userAgent"`
	RequestUri string   `json:"requestUri"`
	Headers     *Headers `json:"headers"`
}

type Headers struct {
	UserAgent string `json:"UserAgent"`
	Accept     string `json:"Accept"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Host: r.Host,
		UserAgent: r.Header.Get("User-agent"),
		RequestUri: r.URL.Path,
		Headers: &Headers{r.Header.Get("User-agent"),
			r.Header.Get("Accept")},
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "err", http.StatusInternalServerError)
		return
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
