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

	rules := parseRules(input)

	total, possibleBags := findPossibleBags(rules)
	fmt.Printf("Total # of Bags: %d\n", total)
	fmt.Printf("Total possible Bags: %d\n", possibleBags)

	sgr := rules["shiny gold"]
	sgr.Print()

	bagsNeeded := countTotalBags(sgr, rules)
	fmt.Printf("bags needed: %d\n", bagsNeeded)
}

func countTotalBags(r rule, rules map[string]rule) int {
	return r.CountBags(rules)
}

type rule struct {
	color string
	bags  map[string]int
}

func (r *rule) AddBag(color string, numberCanHold int) {
	r.bags[color] = numberCanHold
}

func (r *rule) CountBags(allRules map[string]rule) int {
	count := 0
	for color, num := range r.bags {
		b := allRules[color]
		count = count + num
		if len(b.bags) != 0 {
			count = count + num*b.CountBags(allRules)
		}
	}

	return count
}

func (r *rule) CanHold(color string, allRules map[string]rule) bool {
	if _, ok := r.bags[color]; ok {
		//fmt.Printf("found %s\n", r.color)
		return true
	}

	for c := range r.bags {
		b := allRules[c]
		if b.CanHold(color, allRules) {
			return true
		}
	}

	return false
}

func (r *rule) Print() {
	fmt.Println(r)
}

func parseRules(input string) map[string]rule {
	lines := strings.Split(input, "\n")

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
			//fmt.Printf("%s has %s\n", r.color, bags[0])
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
	return rules
}

func findPossibleBags(rules map[string]rule) (int, int) {
	possibleColors := []string{}
	for _, r := range rules {
		if r.CanHold("shiny gold", rules) {
			possibleColors = append(possibleColors, r.color)
		}
	}

	return len(rules), len(possibleColors)
}
