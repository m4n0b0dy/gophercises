package main

import (
	"flag"
	"fmt"
	"github.com/jellydator/ttlcache/v2"
	"hn/hn_client"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var cache ttlcache.SimpleCache = ttlcache.NewCache()

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	cache.SetTTL(time.Duration(10 * time.Second))

	tpl := template.Must(template.ParseFiles("index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// https://dev.to/franciscomendes10866/easy-and-simple-in-memory-cache-in-golang-1lpb
type GetItemResponse struct {
	ItemResponse hn_client.Item
	Itemerr      error
	Itemid       int
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var client hn_client.Client
		ids, err := client.TopItems()
		if err != nil {
			http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
			return
		}

		fetchCount := numStories * 2

		ch := make(chan GetItemResponse, fetchCount-1)
		for i := 0; i < fetchCount; i++ {
			go func(newId int, newRnk int) {
				val, err := cache.Get(string(rune(newRnk)))

				if err != ttlcache.ErrNotFound {
					fmt.Println(val.(GetItemResponse))
					ch <- val.(GetItemResponse) // https://go.dev/play/p/lGseg88K1m
				} else {

					hnItem, err := client.GetItem(newId)

					r := GetItemResponse{ItemResponse: hnItem,
						Itemerr: err,
						Itemid:  newRnk}
					ch <- r
					_ = cache.Set(string(rune(newRnk)), r)
				}
			}(ids[i], i)

		}

		stories := make([]item, fetchCount)
		var nonStories []int
		for i := 0; i < fetchCount; i++ {
			res := <-ch
			item := parseHNItem(res.ItemResponse)
			if !isStoryLink(item) {
				nonStories = append(nonStories, i)
			}
			stories[res.Itemid] = item

		}
		for _, i := range nonStories {
			stories = append(stories[:i], stories[i+1:]...)
		}

		data := templateData{
			Stories: stories[:numStories],
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn_client.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn_client.Item, but adds the Host field
type item struct {
	hn_client.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
