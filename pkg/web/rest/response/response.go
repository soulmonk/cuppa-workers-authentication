package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		log.Fatal("Error write response: ", err.Error())
	}
}

func Ni(w http.ResponseWriter, r *http.Request) {
	RespondWithJson(w, http.StatusOK, "Not implemented")
}
