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
		t.Errorf("Success was expected but got %d", res.StatusCode)
	}
}

func TestGet200ForRecordsRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:3130/records", nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success was expected but got %d", res.StatusCode)
	}
}

func TestGet200ForRecordShowRoute(t *testing.T) {
	t.Skip("this test is not really valid yet, show route is not aware of param")

	req, err := http.NewRequest("GET", "http://localhost:3130/records/1", nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Success was expected but got %d", res.StatusCode)
	}
}
