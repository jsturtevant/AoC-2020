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

	// combinations could will start with 1
	joltFinder := jolter{diffs: make(map[int]int), combinations: []int{}}
	joltFinder.run(joltAdapters, 0, 3)
	fmt.Printf("value %d * %d:  %d\n", joltFinder.Diffs1(), joltFinder.Diffs3(), joltFinder.Diffs1()*joltFinder.Diffs3())

	joltFinder.combos(joltAdapters, 0)
	fmt.Println(joltFinder.combinations)
	fmt.Printf("combinations: %d\n", joltFinder.Combinations())
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
	combinations []int
}

func (j *jolter) Diffs1() int {
	return j.diffs[1]
}

func (j *jolter) Diffs3() int {
	return j.diffs[3] + 1
}

func (j *jolter) Combinations() int {
	total := 1
	for _, count := range j.combinations {
		switch count {
		case 1, 2:
			total *= 1
		case 3:
			// 2^1
			total *= 2
		case 4:
			// 2^2
			total *= 4
		case 5:
			// 2^3 - 1
			total *= 7
		default:
			panic("unkown value")
		}
	}
	return total
}

func (j *jolter) run(joltAdapters []int, startRating int, ratingDiff int) {
	//fmt.Printf("%d, ", startRating)
	for i, adapterRating := range joltAdapters {
		if adapterRating <= startRating+ratingDiff {
			//fmt.Printf("%d, ", adapterRating)
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

func (j *jolter) combos(joltAdapters []int, startRating int) {
	splits := 0
	if len(j.combinations) == 0 {
		//always 1 combination
		j.combinations = append(j.combinations, 1)
	}

	if joltAdapters[0]-startRating == 1 {
		j.combinations[splits]++
	}
	for i, adapterRating := range joltAdapters {
		if i+1 == len(joltAdapters) {
			break
		}

		diff := joltAdapters[i+1] - adapterRating

		if diff == 1 {
			j.combinations[splits]++
		}

		if diff == 3 {
			splits++
			j.combinations = append(j.combinations, 1)
		}
	}
}
