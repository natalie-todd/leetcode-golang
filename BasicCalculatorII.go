package main

import (
	"strconv"
)

//break the string into a slice of strings(chars)
//have a total int
//loop through looking for * or /. if index is * or / use that on the surrounding indices, then remove all three and replace with the total
//loop through and look for + and - if present add the sides or minus them and replace with the number
//factor in spaces
//return slice[0]
func calculate(s string) int {
	var stringSlice []string
	multiplicationCounter := 0
	divisionCounter := 0
	additionCounter := 0
	subtractionCounter := 0
	result := 0
	newString := ""

	for i := range s {
		if string(s[i]) != " " {
			stringSlice = append(stringSlice, string(s[i]))
			newString = newString + string(s[i])
		}
	}

	for len(stringSlice) > 1 {
		for i := 0; i < len(stringSlice); i++ {
			if stringSlice[i] != "*" && stringSlice[i] != "/" && stringSlice[i] != "+" && stringSlice[i] != "-" {
				continue
			}
			if stringSlice[i] == "*" {
				multiplicationCounter++
				continue
			}
			if stringSlice[i] == "/" {
				divisionCounter++
				continue
			}
			if stringSlice[i] == "+" {
				additionCounter++
				continue
			}
			if stringSlice[i] == "-" {
				subtractionCounter++
				continue
			}
		}

		if multiplicationCounter == 0 && divisionCounter == 0 && additionCounter == 0 && subtractionCounter == 0 {
			stringSlice = []string{"a"}
		}

		for multiplicationCounter > 0 {
			for i := 0; i < len(stringSlice); i++ {
				if stringSlice[i] == "*" {
					var numLess, _ = strconv.Atoi(stringSlice[i-1])
					var numMore, _ = strconv.Atoi(stringSlice[i+1])
					var result = numLess * numMore

					stringSlice[i-1] = strconv.Itoa(result)
					stringSlice = append(stringSlice[:i], stringSlice[i+2:]...)
					multiplicationCounter--
				}
			}
		}
	for divisionCounter > 0 {
		for i := 0; i < len(stringSlice); i++ {
		if stringSlice[i] == "/" {
			var numLess, _ = strconv.Atoi(stringSlice[i-1])
			var numMore, _ = strconv.Atoi(stringSlice[i+1])
			var result = numLess / numMore

			stringSlice[i-1] = strconv.Itoa(result)
			stringSlice = append(stringSlice[:i], stringSlice[i+2:]...)
			divisionCounter--
		}
	}
	}
	for additionCounter > 0 {
		for i := 0; i < len(stringSlice); i++ {
		if stringSlice[i] == "+" {
			var numLess, _ = strconv.Atoi(stringSlice[i-1])
			var numMore, _ = strconv.Atoi(stringSlice[i+1])
			var result = numLess + numMore

			stringSlice[i-1] = strconv.Itoa(result)
			stringSlice = append(stringSlice[:i], stringSlice[i+2:]...)
			additionCounter--
		}
	}
	}
	for subtractionCounter > 0 {
		for i := 0; i < len(stringSlice); i++ {
		if stringSlice[i] == "-" {
			var numLess, _ = strconv.Atoi(stringSlice[i-1])
			var numMore, _ = strconv.Atoi(stringSlice[i+1])
			var result = numLess - numMore

			stringSlice[i-1] = strconv.Itoa(result)
			stringSlice = append(stringSlice[:i], stringSlice[i+2:]...)
			subtractionCounter--
		}
	}
	}
	}

	if stringSlice[0] == "a" {
		result, _ = strconv.Atoi(newString)
	} else {
		result, _ = strconv.Atoi(stringSlice[0])
	}
	return result
}

func main() {
	println(calculate("0-2147483647"))
}
