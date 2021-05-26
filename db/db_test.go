package db

import "testing"

func TestConnectdb(t *testing.T) {
	got, ctx, err := CreatedbClient()
	if err != nil {
		t.Errorf("connectiondb() = %v, %v, %v ; want no error", got, ctx, err)
	}

}
