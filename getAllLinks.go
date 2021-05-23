package main

import (
	"io"
	"net/url"
	"regexp"
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

		if token.Type == html.SelfClosingTagToken && token.DataAtom.String() == "meta" {
			for _, attr := range token.Attr {
				if attr.Key == "content" && strings.ToLower(strings.ReplaceAll(attr.Val, " ", "")) == "noindex,nofollow" {
					return links
				}
			}
		}

		if token.Type == html.TextToken {
			re := regexp.MustCompile(`\b\w+\.[\w.]+\b`)
			link := re.Find([]byte(token.Data))
			if link == nil {
				continue
			}
			if isUnique(links, string(link)) {
				links = append(links, "https://"+string(link))
			}
		}

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
