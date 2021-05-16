package main

import (
	"flag"
	"fmt"
	"net/http"
)

//TODO: resolve relative links

var visited = make(map[string]bool)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		panic("Please specify start page")
	}

	q := make(chan string)

	enqueue(q, args[0])

	for url := range q {
		retrieve(url, q)
	}

}

func retrieve(url string, q chan string) {
	fmt.Println("retrieve url: " + url)

	visited[url] = true
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error retrieving page: ", err)
		return
	}

	defer resp.Body.Close()

	links := getAllLinks(resp.Body)

	for _, link := range links {
		if !visited[link] {
			enqueue(q, link)
		}

	}
}

func enqueue(c chan string, url string) {
	go func() {
		c <- url
	}()
}
