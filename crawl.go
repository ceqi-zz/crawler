package main

import (
	"flag"
	"fmt"
	"net/http"
	"regexp"
)

var visited = make(map[string]bool)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		panic("Please specify start page")
	}

	validRoot := regexp.MustCompile(`(https://[\w.]+).*`)
	root := validRoot.FindStringSubmatch(args[0])[0]

	q := make(chan string)

	enqueue(q, args[0])

	for url := range q {
		retrieve(url, root, q)
	}

}

func retrieve(url, root string, q chan string) {
	fmt.Println("retrieve url: " + url)

	visited[url] = true
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error retrieving page: ", err)
		return
	}

	defer resp.Body.Close()

	links := getAllLinks(resp.Body, root)

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
