package main

import "fmt"

func main() {
	message := make(chan string, 2)

	message <- "Hello"
	message <- "hi"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
