package helpers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

type optionArc struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string      `json:"title"`
	Story   []string    `json:"story"`
	Options []optionArc `json:"options"`
}

type Page struct {
	Title string
	Body  []string
}

type StoryArcs map[string]StoryArc

func EditHandler(pageData StoryArc, fallback http.Handler) http.HandlerFunc {
	newFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p Page
		p.Title = pageData.Title
		p.Body = pageData.Story
		t, _ := template.ParseFiles("template.html")
		err := t.Execute(w, p)
		if err != nil {
			return
		}
	})
	return newFunc
}

func LoadStory(p string) StoryArcs {
	file, _ := ioutil.ReadFile(p)
	var m StoryArcs
	json.Unmarshal(file, &m)
	return m
}
