package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("[YourName][Positive|Negative]")
		return
	}

	mood := [...][3]string{
		{"good", "perfect", "amazing"},
		{"terible", "bad", "sad"},
	}

	name, status := args[0], args[1]
	var mood_index int
	if strings.ToLower(status) == "negative" {
		mood_index = 1
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(mood[0][0]) - 1)

	fmt.Printf("%s's mood is %q\n", name, mood[mood_index][n])

}
