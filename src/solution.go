package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(reverse("A4B5C2"))
}

func reverse(input string) string {
	// Use a regular expression to extract all combinations of letters and numbers
	r := regexp.MustCompile("(?P<Letter>[A-z])(?P<Count>\\d)")

	// Find all matches for the regular expression in the input string
	match := r.FindAllStringSubmatch(input, -1)

	output := ""

	// Interate over the combinations, repeating the letters and adding that to output string
	for _, element := range match {
		count, _ := strconv.Atoi(element[2])

		output += strings.Repeat(element[1], count)
	}

	return output
}
