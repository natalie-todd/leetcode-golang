package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var result []string
	counter := 1

	for counter < n + 1 {
		if counter % 15 == 0 {
			result = append(result, "FizzBuzz")
	} else if counter % 5 == 0 {
		result = append(result, "Buzz")
	} else if counter % 3 == 0 {
		result = append(result, "Fizz")
	} else {
		result = append(result, strconv.Itoa(counter))
	}
	counter++
	}
	return result
}

func main() {
	fmt.Print(fizzBuzz(1))
}
