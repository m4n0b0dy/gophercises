package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var quitMessage = "q"
var filePath = "story.json"

type optionArc struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type storyArc struct {
	Title   string      `json:"title"`
	Story   []string    `json:"story"`
	Options []optionArc `json:"options"`
}

type storyArcs map[string]storyArc

func main() {
	loadedStory := load_story(filePath)
	currentArc := loadedStory["intro"]
	for {

		print_story(currentArc)

		if len(currentArc.Options) < 1 {
			break
		}

		userResponse := get_response(currentArc)

		if userResponse == quitMessage {
			break
		} else {
			currentArc = loadedStory[userResponse]
		}

	}
}

func load_story(p string) storyArcs {
	file, _ := ioutil.ReadFile(p)
	var m storyArcs
	json.Unmarshal(file, &m)
	return m
}

func print_story(a storyArc) {
	for _, line := range a.Story {
		fmt.Println(line)
	}
}

func get_response(a storyArc) string {
	fmt.Println("Choose your path:")

	for i, option := range a.Options {
		fmt.Printf("Option %v: %v ----> %v\n", i+1, option.Text, option.Arc)
	}

	var userResponse string
	fmt.Scanln(&userResponse)
	return userResponse
}
