package main

import (
	"go-sync/application"
	"go-sync/web"
	"log"
	"net/http"
)

func main() {
	store := application.NewKeyValueStore()

	handlers := web.NewHandlers(store)

	http.HandleFunc("/", handlers.ListKeys)
	http.HandleFunc("/key/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetKey(w, r)
		case http.MethodPut:
			handlers.PutKey(w, r)
		case http.MethodDelete:
			handlers.DeleteKey(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
