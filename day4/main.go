package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type slope struct {
	right int
	down  int
}

func main() {

	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	input := string(content)

	total := countValidPassprts(input)

	fmt.Printf("Valid passports: %d\n", total)
}

func countValidPassprts(input string) int {
	passports := strings.Split(input, "\n\n")

	validCount := 0
	for _, p := range passports {

		f := strings.Fields(p)

		if len(f) == 8 {
			fmt.Println(p)
			fmt.Println("Valid! All eight fields.")
			validCount = validCount + 1
			continue
		}

		if len(f) == 7 {
			found := false
			for _, v := range f {
				if strings.Contains(v, "cid:") {
					found = true
				}
			}

			if !found {
				fmt.Println(p)
				fmt.Println("Valid! Only missing cid.")
				validCount = validCount + 1
				continue
			}
		}
	}

	return validCount
}
