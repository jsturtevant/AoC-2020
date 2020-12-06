package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	input := string(content)

	sum := sumOfAnswers(input)
	fmt.Printf("Sum of answers: %d\n", sum)
}

func sumOfAnswers(input string) int {
	groupAnswers := strings.Split(input, "\n\n")

	total := 0
	for _, groupAns := range groupAnswers {
		fmt.Println(groupAns)

		individuleAnswers := strings.Split(groupAns, "\n")
		if len(individuleAnswers) == 1 {
			fmt.Printf("Single person in group. %d", len(individuleAnswers[0]))
			total = total + len(individuleAnswers[0])
			fmt.Printf("new total: %d\n\n", total)
			continue
		}

		var matches []rune
		for i, a := range individuleAnswers {
			if matches == nil && i+1 < len(individuleAnswers) {
				// first entry load er up
				matches = intersect(a, individuleAnswers[1])
				if len(matches) == 0 {
					// if the first and second don't have
					// matches then break
					break
				}
			}

			// does this match all previous?
			matches = intersect(a, string(matches))
			if len(matches) == 0 {
				// if no matches break
				break
			}
		}

		fmt.Printf("matches: %d\n\n", len(matches))
		total = total + len(matches)
		fmt.Printf("new total: %d\n\n", total)
	}
	return total
}

func intersect(ans1, ans2 string) []rune {
	matches := make([]rune, 0)
	m := make(map[rune]rune)
	for _, a := range ans1 {
		m[a] = a
	}

	for _, a := range ans2 {
		if _, ok := m[a]; ok {
			matches = append(matches, a)
		}
	}

	return matches
}
