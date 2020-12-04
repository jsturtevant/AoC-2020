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
	s1 := slope{right: 1, down: 1}
	t1 := countTrees(input, s1)
	fmt.Printf("total trees: %d\n", t1)

	s2 := slope{right: 3, down: 1}
	t2 := countTrees(input, s2)
	fmt.Printf("total trees: %d\n", t2)

	s3 := slope{right: 5, down: 1}
	t3 := countTrees(input, s3)
	fmt.Printf("total trees: %d\n", t3)

	s4 := slope{right: 7, down: 1}
	t4 := countTrees(input, s4)
	fmt.Printf("total trees: %d\n", t4)

	s5 := slope{right: 1, down: 2}
	t5 := countTrees(input, s5)
	fmt.Printf("total trees: %d\n", t5)

	total := t1 * t2 * t3 * t4 * t5
	fmt.Printf("mulitple of slopes: %d\n", total)
}

func countTrees(input string, slope slope) int {
	m := generateMap(input)
	printmap(m)

	trees := traverse(m, slope)
	return trees
}

func printmap(m [][]string) {
	// debug output
	// fmt.Println("Current Map:")
	// for _, r := range m {
	// 	fmt.Println(r)
	// }
}

func generateMap(input string) [][]string {
	grid := [][]string{}
	rows := strings.Split(input, "\n")
	for _, r := range rows {
		c := strings.Split(r, "")
		grid = append(grid, c)
	}

	return grid
}

func traverse(m [][]string, slope slope) int {
	trees := 0
	t := tobbogan{x: 0, y: 0}

	for t.y < len(m) {
		istree := t.move(m, slope)

		if istree {
			trees = trees + 1
		}
	}

	return trees
}

type tobbogan struct {
	x int
	y int
}

func (t *tobbogan) move(m [][]string, slope slope) bool {
	t.x = t.x + slope.right
	t.y = t.y + slope.down

	if t.y >= len(m) {
		return false
	}

	if t.x >= len(m[t.y]) {
		t.x = t.x - len(m[t.y])
	}

	c := m[t.y][t.x]

	m[t.y][t.x] = "0"
	printmap(m)

	if c == "#" {
		//fmt.Println("Found tree!")
		m[t.y][t.x] = "X"
		return true
	}

	return false
}
