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
	s := slope{right: 3, down: 1}
	trees := countTrees(input, s)

	fmt.Printf("total trees: %d\n", trees)
}

func countTrees(input string, slope slope) int {
	m := generateMap(input)
	printmap(m)

	trees := traverse(m, slope)
	return trees
}

func printmap(m [][]string) {
	// debug output
	fmt.Println("Current Map:")
	for _, r := range m {
		fmt.Println(r)
	}
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
		fmt.Println("Found tree!")
		m[t.y][t.x] = "X"
		return true
	}

	return false
}
