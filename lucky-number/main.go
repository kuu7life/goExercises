package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurn = 3

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Cannot convert")
		return
	}

	for turn := 0; turn < maxTurn; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("You win")
			return
		}
	}
	fmt.Println("You lost try again")
}
