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

	validRoot := regexp.MustCompile(`(https://[\w.]+)/.+`)
	root := validRoot.FindStringSubmatch(args[0])[0]
	retrieve(args[0], root)
}

func retrieve(url, root string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error retrieving page: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	links := getAllLinks(resp.Body, root)

	for _, link := range links {
		fmt.Println(link)
	}
}
