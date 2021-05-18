package main

import "testing"

func TestIsUnique(t *testing.T) {
	links := []string{"https://www.example1.com", "https://www.example2.com", "https://www.example3.com"}
	link := "https://www.example4.com"
	got := isUnique(links, link)
	if got != true {
		t.Errorf("isUnique(%v,%v ) = %v ; want true", links, link, got)
	}
}
