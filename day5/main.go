package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	input := string(content)
	seat := findSeat(input)
	fmt.Printf("Seat '%s' with id: %d\n", seat.Name(), seat.ID())
}

type seat struct {
	row string
	col string
}

func (s *seat) Name() string {
	return s.row + s.col
}

type bsp struct {
	min   int
	max   int
	lower rune
	upper rune
}

func newBsp(max int, lower, upper rune) bsp {
	return bsp{
		min:   0,
		max:   max,
		lower: lower,
		upper: upper,
	}
}

func (b *bsp) Print() {
	fmt.Printf("Min: %d, Max: %d\n", b.min, b.max)
}

func (b *bsp) Split(splitter rune) {
	fmt.Printf("Split on %s\n", string(splitter))
	b.Print()
	switch splitter {
	case b.lower:
		i := b.max - b.min
		b.max = b.min + i/2
	case b.upper:
		i := b.max - b.min
		b.min = b.min + (i / 2) + 1
	default:
		fmt.Printf("Invalid splitter passed: %s", string(splitter))
	}
	b.Print()
}

func (b *bsp) Value() int {
	if b.min == b.max {
		return b.min
	}
	fmt.Println("Didn't complete splitting")
	return -1
}

func (s *seat) Row() int {
	t := newBsp(127, 'F', 'B')
	for _, c := range s.row {
		t.Split(c)
	}
	return t.Value()
}

func (s *seat) Col() int {
	t := newBsp(7, 'L', 'R')
	for _, c := range s.col {
		t.Split(c)
	}
	return t.Value()
}

func (s *seat) ID() int {
	return (s.Row() * 8) + s.Col()
}

func findSeat(input string) seat {
	rawSeats := strings.Split(input, "\n")

	seats := []seat{}
	for _, rs := range rawSeats {
		re := regexp.MustCompile(`^([FB]{7})([RL]{3})$`)
		s := re.FindStringSubmatch(rs)

		if len(s) != 3 {
			fmt.Println("Invalid Seat")
			continue
		}

		seats = append(seats, seat{row: s[1], col: s[2]})
	}

	sort.Slice(seats, func(i, j int) bool {
		if seats[i].row < seats[j].row {
			return true
		}
		if seats[i].row > seats[j].row {
			return false
		}
		return seats[i].col > seats[j].col
	})

	highestSeat := seats[0]

	return highestSeat
}
