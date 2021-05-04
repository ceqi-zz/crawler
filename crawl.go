package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
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
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("network error: ", err)
		return
	}
	fmt.Println(string(body))
}
