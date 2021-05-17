package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		panic("Please specify start page")
	}

	q := make(chan string)
	filteredq := make(chan string)

	go enqueue(args[0], q)

	go filterq(q, filteredq)

	for link := range filteredq {
		go enqueue(link, q)
	}

}

func filterq(q, filteredq chan string) {
	var visited = make(map[string]bool)
	for link := range q {
		if !visited[link] {
			visited[link] = true
			filteredq <- link
		}
	}
}

func enqueue(rawurl string, q chan string) {
	fmt.Println("retrieve url: " + rawurl)

	resp, err := http.Get(rawurl)
	if err != nil {
		fmt.Println("error retrieving page: ", err)
		return
	}

	defer resp.Body.Close()

	links := getAllLinks(resp.Body) // can have relative links

	baseuri, _ := url.Parse(rawurl)

	for _, link := range links {
		u, _ := url.Parse(link)
		absuri := baseuri.ResolveReference(u)
		link = absuri.String()
		q <- link
	}
}
