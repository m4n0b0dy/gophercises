package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type customLink string

type pageLink struct {
	Href string
	Text string
}

type pageData []pageLink

func main() {
	pagestring, err := ioutil.ReadFile("4.html")
	if err != nil {
		log.Fatal(err)
	}

	text := customLink(pagestring)

	for _, result := range text.linkExtract() {
		fmt.Printf("Link: %v | Text: %v\n", result.Href, result.Text)
	}

}

func (s customLink) linkExtract() pageData {
	doc, err := html.Parse(strings.NewReader(string(s)))
	if err != nil {
		log.Fatal(err)
	}
	foundLinks := pageData{}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					currentPage := pageLink{
						Href: a.Val,
						Text: (*n.FirstChild).Data,
					}
					foundLinks = append(foundLinks, currentPage)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

	}
	f(doc)
	return foundLinks

}
