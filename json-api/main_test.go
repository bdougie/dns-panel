package main

import (
	"net/http"
	"testing"
)

func TestGet200(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:3130/", nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestGet200ForRecordRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:3130/records", nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
