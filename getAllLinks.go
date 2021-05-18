package main

import (
	"io"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// get all links from current page
func getAllLinks(respbody io.Reader, rawurl string) []string {
	var links []string
	baseuri, _ := url.Parse(rawurl)

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
					link := attr.Val
					if strings.Contains(link, "#") {
						break
					}

					link = resolveLink(link, baseuri)

					if link == "" {
						break
					}

					if isUnique(links, link) {
						links = append(links, link)
					}

				}
			}
		}
	}
}

// relative links eg. "/subpage/sectionx" need to be resolved
func resolveLink(link string, baseuri *url.URL) string {
	u, _ := url.Parse(link)
	if u != nil {
		absuri := baseuri.ResolveReference(u)
		return absuri.String()
	}

	return ""
}

func isUnique(links []string, link string) bool {
	for _, lk := range links {
		if link == lk {
			return false
		}
	}

	return true
}
