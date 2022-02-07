package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me the size of the table")
		return
	}

	max, err := strconv.Atoi(args[1])
	if err != nil || max < 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", "X")

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}

	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", j*i)
		}
		fmt.Println()
	}

}
