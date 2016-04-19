package main

type Record struct {
	ID      string `json:"id"`
	Domain  string `json:"domain"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Records []Record
