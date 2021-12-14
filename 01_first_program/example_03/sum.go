package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithProg := os.Args

	numA, err := strconv.Atoi(argsWithProg[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	numB, err := strconv.Atoi(argsWithProg[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	result := numA + numB
	fmt.Printf("%d + %d = %d\n", numA, numB, result)
}
