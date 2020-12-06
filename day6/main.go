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

		unique := make(map[rune]int)
		for _, a := range groupAns {
			if string(a) == "\n" {
				continue
			}

			if count, ok := unique[a]; ok {
				unique[a] = count + 1
			} else {
				unique[a] = 1
			}
		}

		fmt.Printf("unique: %d\n\n", len(unique))
		total = total + len(unique)
	}
	return total
}
