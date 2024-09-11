package web

import (
	"encoding/json"
	"errors"
	"go-sync/application"
	"io/ioutil"
	"net/http"
)

// Handlers structure holds the key-value store instance
type Handlers struct {
	Store *application.KVStore
}

// NewHandlers creates a new instance of Handlers
func NewHandlers(store *application.KVStore) *Handlers {
	return &Handlers{Store: store}
}

// GetKey handles the GET /key/{key} request
func (h *Handlers) GetKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/key/"):]
	value, err := h.Store.Get(key)
	if err != nil {
		if errors.Is(err, application.ErrKeyNotFound) {
			http.Error(w, "Key not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	w.Write([]byte(value))
}

// PutKey handles the PUT /key/{key} request
func (h *Handlers) PutKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/key/"):]

	if r.Body == nil {
		http.Error(w, "Invalid or missing request body", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Invalid or missing request body", http.StatusBadRequest)
		return
	}

	h.Store.Put(key, string(body))
	w.WriteHeader(http.StatusOK)
}

// DeleteKey handles the DELETE /key/{key} request
func (h *Handlers) DeleteKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/key/"):]
	err := h.Store.Delete(key)
	if err != nil {
		if errors.Is(err, application.ErrKeyNotFound) {
			http.Error(w, "Key not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ListKeys handles the GET / request to list all keys
func (h *Handlers) ListKeys(w http.ResponseWriter, r *http.Request) {
	keys := h.Store.ListKeys()
	jsonResponse, _ := json.Marshal(keys)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
