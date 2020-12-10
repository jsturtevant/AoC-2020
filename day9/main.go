package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	input := string(content)
	xmas := parseInput(input)

	value := run(xmas, 25)
	fmt.Printf("invalid value %d\n", value)
	crackedValue := crack(value, xmas)
	fmt.Printf("cracked value: %d\n", crackedValue)
}

func crack(value int, xmas []int) int {
	for i, _ := range xmas {
		v, success := addNumbers(i, xmas, value)
		if success {
			return v
		}
	}

	return 0
}

func addNumbers(start int, xmas []int, value int) (int, bool) {
	total := 0
	var end int
	for end = start; end < len(xmas); end++ {
		total = total + xmas[end]
		if total == value {
			break
		}

		if total > value {
			return -1, false
		}
	}

	contiguous := xmas[start:end]
	sort.Ints(contiguous)

	return contiguous[0] + contiguous[len(contiguous)-1], true
}

func parseInput(input string) []int {
	xmas := strings.Split(input, "\n")
	xmasNums := make([]int, len(xmas))
	for i, n := range xmas {
		v, _ := strconv.Atoi(n)
		xmasNums[i] = v
	}

	return xmasNums
}

func run(xmas []int, preambleSize int) int {

	for i := preambleSize; i < len(xmas); i++ {

		currentValue := xmas[i]
		preamble := xmas[i-preambleSize : i+preambleSize]

		if !findInvalid(preamble, currentValue) {
			return currentValue
		}
	}

	return -1
}

func findInvalid(preamble []int, num int) bool {
	//brute force
	for _, a := range preamble {
		for _, b := range preamble {
			if a == b {
				continue
			}

			if (a + b) == num {
				return true
			}
		}
	}

	return false
}
