package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var tmpl = template.Must(template.ParseFiles("templates/error.html", "templates/index.html", "templates/secondpage.html", "templates/thirdpage.html"))

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		data := map[string]any{"code": http.StatusNotFound, "msg": "page not found"}
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "error.html", data)
		return
	}
	if r.Method != "GET" {
		data := map[string]any{"code": http.StatusMethodNotAllowed, "msg": "method not allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		tmpl.ExecuteTemplate(w, "error.html", data)
		return
	}

	url := "https://groupietrackers.herokuapp.com/api/artists"
	data := Fetch(url, w)
	var artists []Artist
	DecodeByUs(data, &artists, w)

	if err := tmpl.ExecuteTemplate(w, "index.html", artists); err != nil {
		data := map[string]any{"code": http.StatusInternalServerError, "msg": "failed executing"}
		tmpl.ExecuteTemplate(w, "error.html", data)
		return
	}
}

func SecondPage(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/details/") {
		data := map[string]any{"code": http.StatusNotFound, "msg": "page not found"}
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "error.html", data)
		return
	}
	if r.Method != "GET" {
		data := map[string]any{"code": http.StatusMethodNotAllowed, "msg": "method not allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		tmpl.ExecuteTemplate(w, "error.html", data)
		return
	}

	idStr := r.URL.Path[len("/details/"):]

	ids, _ := strconv.Atoi(idStr)
	if ids <= 0 || ids > 52 {
		data := map[string]any{"code": http.StatusNotFound, "msg": "page not found"}
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "error.html", data)
		return
	}

	url1 := "https://groupietrackers.herokuapp.com/api/artists/" + idStr
	data1 := Fetch(url1, w)

	var artist Artist
	DecodeByUs(data1, &artist, w)

	url2 := "https://groupietrackers.herokuapp.com/api/locations/" + idStr
	data2 := Fetch(url2, w)

	var locationsResponse LocationsResponse
	DecodeByUs(data2, &locationsResponse, w)

	pageData := SecondPageData{
		ID:        locationsResponse.ID,
		Artist:    artist,
		Locations: locationsResponse.Locations,
	}

	if err := tmpl.ExecuteTemplate(w, "secondpage.html", pageData); err != nil {
		data := map[string]any{"code": http.StatusInternalServerError, "msg": "error executing"}
		tmpl.ExecuteTemplate(w, "error.html", data)
		return
	}
}

func LastPage(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/lastpage/") {
		data := map[string]any{"code": http.StatusNotFound, "msg": "page not found"}
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "error.html", data)
		return
	}
	if r.Method != "GET" {
		data := map[string]any{"code": http.StatusMethodNotAllowed, "msg": "method not allowed"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		tmpl.ExecuteTemplate(w, "error.html", data)
		return
	}
	idStr := r.URL.Path[len("/lastpage/"):]
	ids, _ := strconv.Atoi(idStr)
	if ids <= 0 || ids > 52 {
		data := map[string]any{"code": http.StatusNotFound, "msg": "page not found"}
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "error.html", data)
		return
	}

	url := "https://groupietrackers.herokuapp.com/api/relation/" + idStr
	response := Fetch(url, w)

	var relations Relations
	DecodeByUs(response, &relations, w)

	url2 := "https://groupietrackers.herokuapp.com/api/dates/" + idStr
	response2 := Fetch(url2, w)

	var dates Dates
	DecodeByUs(response2, &dates, w)

	pageData := ThirdPageData{
		ID:             relations.ID,
		Dates:          dates.DATES,
		DatesLocations: relations.DATESLOCAT,
	}

	if err := tmpl.ExecuteTemplate(w, "thirdpage.html", pageData); err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}
