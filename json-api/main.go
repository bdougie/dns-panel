package main

func main() {
	Open()        // init DB
	defer Close() // close DB when done

	StartSever()
}
