package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		panic("Please specify start page")
	}

	validRoot := regexp.MustCompile(`(https://[\w.]+).*`)
	root := validRoot.FindStringSubmatch(args[0])[0]

	q := make(chan string)

	go func() {
		q <- args[0]
	}()

	for url := range q {
		retrieve(url, root, q)
	}

}

func retrieve(url, root string, q chan string) {
	fmt.Println("retrieve url: " + url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error retrieving page: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	links := getAllLinks(resp.Body, root)

	for _, link := range links {
		go func() {
			q <- link
		}()
	}
}
