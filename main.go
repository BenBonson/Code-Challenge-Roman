package main

import (
	"fmt"
	"strings"
)

// collection of values as global variables
var out = 0
var current = 0
var previous = 0
var total = 0

// program starts here
func main() {
	var roman string
	// creation of a map of values
	var m map[rune]int
	m = make(map[rune]int)
	m['i'] = 1
	m['v'] = 5
	m['x'] = 10
	m['l'] = 50
	m['c'] = 100
	m['m'] = 1000
	// get input from user
	fmt.Println("Roman numerals, Please enter your symbols")
	fmt.Scanln(&roman)

	// make user input lowercase so it works with capital or lowercase
	var lroman = strings.ToLower(roman)

	// iterate through user input one letter at a time
	for i, letter := range lroman {
		// get value from map
		out = m[letter]
		// make current value previous value
		previous = current
		// reset current
		current = 0 + out
		// if it is not the first letter
		if i > 0 {
			// check if it is something like IV or XL
			if previous < current {
				//translate() old code left for refrence
				// subtract previous value twice
				total = total - previous
				total = total - previous
				// add current to total
				total = total + current
			} else {
				// add current to total
				total = total + current
			}
		} else {
			total = total + current
		}
	}
	// give total
	fmt.Println("Here is your number", total)
}

// Attempt with If statements

// func translate() {
// for i, letter := range lroman {
// 	letter2 := i + 1
// 	letter3 := i - 1
// 	if letter == "i" && letter3 != "v" && letter3 != "x" {
// 		out = out + 1
// 	} else if letter == "v" && letter2 == "i" {
// 		out = out + 4
// 	} else if letter == "v" && letter2 != "i" {
// 		out = out + 5
// 	} else if letter == "x" && letter2 == "i" {
// 		out = out + 9
// 	} else if letter == "x" && letter2 != "i" && letter3 != "l" && letter3 != "c" {
// 		out = out + 10
// 	} else if letter == "l" && letter2 == "x" {
// 		out = out + 40
// 	} else if letter == "l" && letter2 != "x" {
// 		out = out + 50
// 	} else if letter == "c" && letter2 == "x" {
// 		out = out + 90
// 	} else if letter == "c" && letter2 != "x" {
// 		out = out + 100
// 	} else if letter == "d" && letter2 == "c" {
// 		out = out + 400
// 	} else if letter == "d" && letter2 != "c" {
// 		out = out + 500
// 	} else if letter == "m" && letter2 == "c" {
// 		out = out + 900
// 	} else if letter == "m" && letter2 != "c" {
// 		out = out + 1000
// 	} else {
// 		fmt.Println("invalid input")
// 	}
// }
// 	return out
// }
