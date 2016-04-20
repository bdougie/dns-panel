package main

import (
	"brian-dns-panel/json-api/lib"
)

func main() {
	lib.Open()        // init DB
	defer lib.Close() // close DB when done

	lib.StartSever()
}
