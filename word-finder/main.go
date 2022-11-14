package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "the lazy cat jumps or smile again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]
queries:
	for _, q := range query {
	search:
		for i, w := range words {
			switch q { // unwanted query words
			case "and", "or", "the":
				break search
			}
			if q == w {
				fmt.Printf("#%-2d: %q\n",
					i+1, w)
				break queries
			}
		}
	}
}
