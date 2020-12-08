package main

import (
	"testing"
)

const input = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

func Test_findPossibleBags(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		total    int
		possible int
	}{
		{
			name: "given example",
			args: args{
				input: input,
			},
			total:    9,
			possible: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := parseRules(tt.args.input)
			got, got1 := findPossibleBags(rules)
			if got != tt.total {
				t.Errorf("findPossibleBags() got total = %v, want %v", got, tt.total)
			}
			if got1 != tt.possible {
				t.Errorf("findPossibleBags() got possible = %v, want %v", got1, tt.possible)
			}
		})
	}
}

func Test_countTotalBags(t *testing.T) {
	type args struct {
		r     rule
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example",
			args: args{
				input: input,
				r: rule{
					color: "shiny gold",
					bags: map[string]int{
						"dark olive":   1,
						"vibrant plum": 2,
					},
				},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := parseRules(tt.args.input)
			if got := countTotalBags(tt.args.r, rules); got != tt.want {
				t.Errorf("countTotalBags() = %v, want %v", got, tt.want)
			}
		})
	}
}
