package server

import (
	"encoding/json"
	"net/http"
)

func Fetch(url string, w http.ResponseWriter) *http.Response {
	data1, err := http.Get(url)
	if err != nil {
		data := map[string]any{"code": http.StatusInternalServerError, "msg": "failed fetshing data"}
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "error.html", data)
	}
	return data1
}

func DecodeByUs(db *http.Response, pointer any, w http.ResponseWriter) {
	if err := json.NewDecoder(db.Body).Decode(&pointer); err != nil {
		data := map[string]any{"code": http.StatusInternalServerError, "msg": "error decoding"}
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "error.html", data)
	}
}
