package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	const (
		username = "kutman"
		password = 1999
	)

	if l == 1 || l == 0 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	p, _ := strconv.Atoi(args[2])

	if username == args[1] && password == p {
		fmt.Printf("Access granted to %q\n", username)
	} else if username != args[1] || password != p {
		fmt.Printf("Access denied for %q\n", args[1])
	}

	fmt.Println(len(args))
	fmt.Println(l)
}
