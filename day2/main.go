package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}
	input := strings.Split(string(content), "\n")

	valid := countValid(input)

	fmt.Printf("total valid: %d\n", valid)
}

func countValid(rawInput []string) int {
	totalValid := 0
	for _, raw := range rawInput {
		split := strings.Split(raw, ":")

		if len(split) != 2 {
			fmt.Printf("error parsing values: %s\n", raw)
			continue
		}

		policy := strings.TrimSpace(split[0])
		pw := strings.TrimSpace(split[1])
		valid := validate(policy, pw)

		if valid {
			fmt.Printf("VALID policy: %s, pw: %s\n", policy, pw)
			totalValid = totalValid + 1
		}
	}

	return totalValid
}

func validate(rawPolicy, pw string) bool {
	p := newPolicy(rawPolicy)

	return p.validate(pw)
}

type policy struct {
	min           int
	max           int
	requiredValue string
}

func newPolicy(rawPolicy string) policy {
	re := regexp.MustCompile(`([0-9]*)-([0-9]*) (.)`)
	matches := re.FindAllStringSubmatch(rawPolicy, -1)
	//fmt.Printf("%q\n", re.FindAllStringSubmatch(rawPolicy, -1))

	if len(matches) <= 0 {
		fmt.Printf("error parsing values: %s\n", rawPolicy)
		panic("cannot read")
	}

	if len(matches[0]) != 4 {
		fmt.Printf("error parsing values: %s\n", rawPolicy)
	}

	min, err := strconv.Atoi(matches[0][1])
	if err != nil {
		fmt.Printf("error converting min: %s\n", matches[1])
	}
	max, err := strconv.Atoi(matches[0][2])
	if err != nil {
		fmt.Printf("error converting min: %s\n", matches[2])
	}

	return policy{
		min:           min,
		max:           max,
		requiredValue: matches[0][3],
	}
}

func (p *policy) validate(pw string) bool {
	content := []byte(pw)
	re := regexp.MustCompile(p.requiredValue)
	found := re.FindAllIndex(content, -1)

	if len(found) > 0 {
		matches := 0
		for _, idx := range found {
			i := idx[0] + 1

			if i == p.min && i == p.max {
				return false
			}

			if i == p.min || i == p.max {
				matches = matches + 1
			}
		}

		if matches == 1 {
			return true
		}
	}

	return false
}
