package main

import (
	"html/template"
	"net/http"
)

//https://freshman.tech/web-development-with-go/

type optionArc struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string      `json:"title"`
	Story   []string    `json:"story"`
	Options []optionArc `json:"options"`
}

type page struct {
	Title string
	Body  string
}

type StoryArcs map[string]StoryArc

var port = "3000"

func main() {

	mux := http.NewServeMux()
	s := StoryArc{Title: "test", Story: []string{"testing"}, Options: []optionArc{}}
	mux.HandleFunc("/", s.indexHandler)
	http.ListenAndServe(":"+port, mux)
}

var tpl = template.Must(template.ParseFiles("template.html"))

func (s *StoryArc) indexHandler(w http.ResponseWriter, r *http.Request) {
	p := page{Title: s.Title, Body: s.Story[0]}
	tpl.Execute(w, p)

}
