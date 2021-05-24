package main

import (
	"flag"
	"log"
	"net/http"
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

func enqueue(url string, q chan string) error {
	log.Println("retrieve url: " + url)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("error retrieving page: ", err)
		return err
	}

	defer resp.Body.Close()

	links := getAllLinks(resp.Body, url)

	for _, link := range links {
		q <- link
	}

	return nil
}
