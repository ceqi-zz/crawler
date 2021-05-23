package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestIsUnique(t *testing.T) {
	links := []string{"https://www.example1.com", "https://www.example2.com", "https://www.example3.com"}
	link := "https://www.example4.com"
	got := isUnique(links, link)
	if got != true {
		t.Errorf("isUnique(%v,%v ) = %v ; want true", links, link, got)
	}
}

func TestResolveLink(t *testing.T) {
	base, _ := url.Parse("https://www.example.com")
	link := "subpage"
	got := resolveLink(link, base)
	uri := "https://www.example.com/subpage"
	if got != uri {
		t.Errorf("resolveLink(%v, %v) = %v ; want %v", link, base, got, uri)
	}
}

func TestGetAllLinks(t *testing.T) {
	t.Run("get all links", func(t *testing.T) {
		respbody := strings.NewReader(`<a href="/subpage">subpage</a> <a href="https://www.example2.com">example2</a>`)
		rawurl := "https://www.example.com/abc"
		got := getAllLinks(respbody, rawurl)
		expected := []string{"https://www.example.com/subpage", "https://www.example2.com"}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("getAllLinks(%v, %v) = %v ; want %v", respbody, rawurl, got, expected)
		}
	})
	t.Run("remove duplicate links", func(t *testing.T) {
		respbody := strings.NewReader(`<a href="/subpage">subpage</a> <a href="/subpage">subpage</a>`)
		rawurl := "https://www.example.com/abc"
		got := getAllLinks(respbody, rawurl)
		expected := "https://www.example.com/subpage"
		if got == nil || got[0] != expected {
			t.Errorf("getAllLinks(%v, %v) = %v ; want %v", respbody, rawurl, got, expected)
		}
	})
	t.Run("parse links in the form of text", func(t *testing.T) {
		respbody := strings.NewReader(`<html> <body> <p> linkintext.io </p> </body> </html>`)
		rawurl := "https://www.example.com/abc"
		got := getAllLinks(respbody, rawurl)
		expected := "https://linkintext.io"
		if got == nil || got[0] != expected {
			t.Errorf("getAllLinks(%v, %v) = %v ; want %v", respbody, rawurl, got, expected)
		}
	})
	t.Run("links are not followed when meta tag contains nofollow", func(t *testing.T) {
		respbody := strings.NewReader(`<html> 
		<head> <meta name="robots" content="NOINDEX, nofollow" /> </head> 
		<body> <a href="https://www.anotherexample.com">anotherexample</a> </body> </html>`)
		rawurl := "https://www.example.com"
		got := getAllLinks(respbody, rawurl)

		if got != nil {
			t.Errorf("getAllLinks(%v, %v) = %v ; want nil", respbody, rawurl, got)
		}

	})

}
