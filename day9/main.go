package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	value := run(input, 25)

	fmt.Printf("value %s\n", value)
}

func run(input string, preambleSize int) string {
	xmas := strings.Split(input, "\n")
	for i := preambleSize; i < len(xmas); i++ {

		currentValue := xmas[i]
		preamble := xmas[i-preambleSize : i+preambleSize]

		if !findInvalid(preamble, currentValue) {
			return currentValue
		}
	}

	return ""
}

func findInvalid(preamble []string, num string) bool {
	//brute force
	for _, a := range preamble {
		for _, b := range preamble {
			if a == b {
				continue
			}

			aNum, _ := strconv.Atoi(a)
			bNum, _ := strconv.Atoi(b)
			numNum, _ := strconv.Atoi(num)

			if (aNum + bNum) == numNum {
				return true
			}
		}
	}

	return false
}
