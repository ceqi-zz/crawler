package main

import (
	"net/url"
	"testing"
)

func TestEnqueue(t *testing.T) {
	t.Run("processing invalid urls throw url.Error", func(t *testing.T) {
		q := make(chan string)
		uri := "inva.lid"
		got := enqueue(uri, q)

		if got != got.(*url.Error) {
			t.Errorf("enquue(%v, %v) = %v ; want *url.Error", uri, q, got)
		}

	})
}
