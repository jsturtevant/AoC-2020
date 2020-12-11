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
	joltAdapters := parseInput(input)

	joltFinder := jolter{diffs: make(map[int]int)}
	joltFinder.run(joltAdapters, 0, 3)
	fmt.Printf("value %d\n", joltFinder.Diffs1()*joltFinder.Diffs3())
}

func parseInput(input string) []int {
	joltAdapters := strings.Split(input, "\n")
	joltAdaptersNums := make([]int, len(joltAdapters))
	for i, n := range joltAdapters {
		v, _ := strconv.Atoi(n)
		joltAdaptersNums[i] = v
	}
	sort.Ints(joltAdaptersNums)
	return joltAdaptersNums
}

type jolter struct {
	diffs        map[int]int
	combinations int
}

func (j *jolter) Diffs1() int {
	return j.diffs[1]
}

func (j *jolter) Diffs3() int {
	return j.diffs[3] + 1
}

func (j *jolter) run(joltAdapters []int, startRating int, ratingDiff int) {
	fmt.Printf("start rating: %d, len adapters: %d\n", startRating, len(joltAdapters))
	for i, adapterRating := range joltAdapters {
		if adapterRating <= startRating+ratingDiff {
			fmt.Print(adapterRating)
			diff := adapterRating - startRating
			count, _ := j.diffs[diff]
			j.diffs[diff] = count + 1

			j.run(joltAdapters[i+1:], adapterRating, ratingDiff)
			return
		}

		if adapterRating > startRating+ratingDiff {
			fmt.Println("something went wrong?")
			return
		}
	}
}
