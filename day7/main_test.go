package main

import "testing"

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
			got, got1 := findPossibleBags(tt.args.input)
			if got != tt.total {
				t.Errorf("findPossibleBags() got total = %v, want %v", got, tt.total)
			}
			if got1 != tt.possible {
				t.Errorf("findPossibleBags() got possible = %v, want %v", got1, tt.possible)
			}
		})
	}
}
