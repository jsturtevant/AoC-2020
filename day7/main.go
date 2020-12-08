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

	input := string(content)

	total, possibleBags := findPossibleBags(input)
	fmt.Printf("Total # of Bags: %d\n", total)
	fmt.Printf("Total possible Bags: %d\n", possibleBags)
}

type rule struct {
	color string
	bags  map[string]int
}

func (r *rule) AddBag(color string, numberCanHold int) {
	r.bags[color] = numberCanHold
}

func (r *rule) CanHold(color string, allRules map[string]rule) bool {
	if _, ok := r.bags[color]; ok {
		fmt.Printf("found %s\n", r.color)
		return true
	}

	for c := range r.bags {
		r := allRules[c]
		if r.CanHold(color, allRules) {
			return true
		}
	}

	return false
}

func (r *rule) Print() {
	fmt.Println(r)
}

func findPossibleBags(input string) (int, int) {
	lines := strings.Split(input, "\n")
	total := len(lines)

	rules := make(map[string]rule)
	for _, l := range lines {
		parts := strings.Split(l, "bags contain")
		if len(parts) != 2 {
			fmt.Println("Not a valid rule")
		}

		r := rule{color: strings.TrimSpace(parts[0]), bags: make(map[string]int)}
		rules[r.color] = r
		//r.Print()

		bags := strings.Split(parts[1], ",")

		if len(bags) == 1 && strings.TrimSpace(bags[0]) == "no other bags." {
			fmt.Printf("%s has %s\n", r.color, bags[0])
			continue
		}

		for _, b := range bags {
			re := regexp.MustCompile(`([1-9]*) ([\S ]*) bag[s.]?`)
			m := re.FindStringSubmatch(strings.TrimSpace(b))

			if len(m) != 3 {
				fmt.Println("Invalid bag rule")
			}

			numOfBags, _ := strconv.Atoi(m[1])
			r.AddBag(m[2], numOfBags)
		}
	}

	possibleColors := []string{}
	for _, r := range rules {
		if r.CanHold("shiny gold", rules) {
			possibleColors = append(possibleColors, r.color)
		}
	}

	return total, len(possibleColors)
}
