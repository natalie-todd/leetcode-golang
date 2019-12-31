package main

import (
	"fmt"
)


//break the string into a slice of strings(chars)
//have a total int
//loop through looking for * or /. if index is * or / use that on the surrounding indices, then remove all three and replace with the total
//loop through and look for + and - if present add the sides or minus them and replace with the number
//factor in spaces
//return slice[0]
func calculate(s string) int {
	var stringSlice []string

	for i := range s {
		if string(s[i]) != " " {
			stringSlice= append(stringSlice, string(s[i]))
		}
	}
	fmt.Println(stringSlice)
	return len(stringSlice)
}

func main() {
	println(calculate("test Test"))
}
