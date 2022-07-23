package handlers

import (
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
)

// his implementation
func _MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, //
		r *http.Request) {
		// if we can match a path to the paths ToUrls else call fallback
		path := r.URL.Path // path is the attribute URL attribute Path
		// new syntax, can declare variables but separate with ; to the ok check
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func _YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		log.Fatalf("error")
	}
	// convert array to map
	pathsToUrls := make(map[string]string)
	for _, val := range pathUrls {
		pathsToUrls[val.Path] = val.URL
	}
	return MapHandler(pathsToUrls, fallback), nil
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	newFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		found, pass := pathsToUrls[r.URL.Path]
		if pass {
			http.Redirect(w, r, found, 301)
		} else {
			fallback.ServeHTTP(w, r)
		}
	})
	return newFunc
}

type T struct {
	paths []struct {
		path string `yaml:"path"`
		url  string `yaml:"url"`
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	data := T{}
	// marshal is turning data into yaml, unmarshling is taking it out
	err := yaml.Unmarshal(yml, &data)
	if err != nil {
		log.Fatalf("error")
	}
	pathsToUrls := make(map[string]string)
	for _, val := range data.paths {
		pathsToUrls[val.path] = val.url
	}

	handlerFunc := MapHandler(pathsToUrls, fallback)

	return handlerFunc, nil
}
