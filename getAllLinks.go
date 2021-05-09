package main

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// get all links from current page
func getAllLinks(respbody io.Reader, root string) []string {
	var links []string
	z := html.NewTokenizer(respbody)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			return links
		}
		token := z.Token()

		if token.Type == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					if strings.Contains(attr.Val, "#") {
						break
					}

					var link string
					if strings.HasPrefix(attr.Val, "/") {
						link = root + attr.Val
					} else {
						link = attr.Val
					}

					if isUnique(links, link) {
						links = append(links, link)
					}

				}
			}
		}
	}
}

func isUnique(links []string, link string) bool {
	for _, lk := range links {
		if link == lk {
			return false
		}
	}

	return true
}
