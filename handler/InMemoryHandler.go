package handler

import (
	"YS_TestCase/service"
	"encoding/json"
	"net/http"
)

type InMemoryHandler struct {
	inMemoryService service.InMemoryService
}

func (handler *InMemoryHandler) GetAllValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := handler.inMemoryService.FindAllValues()

	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func (handler *InMemoryHandler) AddValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	response := handler.inMemoryService.AddValue(key, value)

	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func (handler *InMemoryHandler) DeleteValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := r.URL.Query().Get("key")

	response := handler.inMemoryService.DeleteValue(key)
	w.WriteHeader(response.Code)
	json.NewEncoder(w).Encode(response)
}

func (handler *InMemoryHandler) FlushValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := handler.inMemoryService.Flush()

	w.WriteHeader(response.Code)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}