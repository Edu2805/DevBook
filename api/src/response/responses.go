package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if jsonError := json.NewEncoder(w).Encode(data); jsonError != nil {
			log.Fatal(jsonError)
		}
	}
}

// ResponseError retorna um erro no formato JSON
func ResponseError(w http.ResponseWriter, statusCode int, responseError error) {
	JSON(w, statusCode, struct {
		JsonError string `json:"jsonError"`
	}{
		JsonError: responseError.Error(),
	})
}
