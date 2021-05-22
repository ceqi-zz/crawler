package db

import "testing"

func TestConnectdb(t *testing.T) {
	got := connectmdb()
	if got != 1 {
		t.Errorf("connectiondb() = %v ; want 1", got)
	}

}
