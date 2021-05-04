package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		panic("Please specify start page")
	}

	retrieve(args[0])
}

func retrieve(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error retrieving page: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	links := getAllLinks(resp.Body)

	for _, link := range links {
		fmt.Println(link)
	}
}

func getAllLinks(respbody io.Reader) []string {
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
					links = append(links, attr.Val)
				}
			}
		}
	}
}
