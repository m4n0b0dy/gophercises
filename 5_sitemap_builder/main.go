package main

import (
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

type Page struct {
	node     *html.Node
	linkText string
	depth    int
	children []Page
}

var base = "https://webscraper.io/test-sites"
var maxDepth = 4
var maxChildren = 5
var alreadySeen []string

func (p *Page) setPageChildren() {
	depth := p.depth + 1
	childCounter := 0
	for c := p.node.FirstChild; c != nil; c = c.NextSibling {
		res := extractLinks(c, depth, &childCounter)
		p.children = append(p.children, res...)
	}

}

func loadPage(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		panic(1)
	}
	n, err := html.Parse(resp.Body)
	if err != nil {
		panic(1)
	}
	return n
}

func main() {
	start := "https://webscraper.io/test-sites"
	n := loadPage(start)
	basePage := Page{node: n, depth: 0}
	basePage.setPageChildren()
	childPages := basePage.children
	var childPage Page
	for len(childPages) > 0 && childPage.depth < maxDepth {
		childPage, childPages = childPages[0], childPages[1:]
		childPage.setPageChildren()
		childPages = append(childPages, childPage.children...)
	}

}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func extractLinks(c *html.Node, depth int, children *int) []Page {
	if *children >= maxChildren {
		return []Page{}
	}

	if c.Type == html.ElementNode && c.Data == "a" {
		for _, attr := range c.Attr {
			if attr.Key == "href" {
				con := base + attr.Val
				if !strings.HasPrefix(attr.Val, "http") && !contains(alreadySeen, con) {
					*children++
					alreadySeen = append(alreadySeen, con)
					return []Page{{node: loadPage(con),
						linkText: con,
						depth:    depth}}

				}

			}
		}

	}
	var ret []Page
	for n := c.FirstChild; n != nil && *children < maxChildren; n = n.NextSibling {
		res := extractLinks(n, depth, children)
		ret = append(ret, res...) // variatic parameters, feeding in multiple values to append
	}
	return ret
}
