package server

import (
	"encoding/json"
	"net/http"
)

// feth mean bring the data from the source -> in this case  by api
// http.get => returns the whole response and error  if exist
func Fetch(url string, w http.ResponseWriter) *http.Response {
	data1, err := http.Get(url)
	if err != nil {
		data := map[string]any{"code": http.StatusInternalServerError, "msg": "failed fetshing data"}
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "error.html", data)
	}
	return data1
}

// decode parameters => res , pointer  type any to work with all daata types , and http resp
// newdecoder transform the body of response from json form to go struct the decode store it into the struct
func DecodeByUs(db *http.Response, pointer any, w http.ResponseWriter) {
	if err := json.NewDecoder(db.Body).Decode(&pointer); err != nil {
		data := map[string]any{"code": http.StatusInternalServerError, "msg": "error decoding"}
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "error.html", data)
	}
}
