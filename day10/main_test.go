package main

import (
	"fmt"
	"testing"
)

const input = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

const input2 = `16
10
15
5
1
11
7
19
6
12
4`

func Test_jolter_run(t *testing.T) {
	type fields struct {
		diffs map[int]int
	}
	type args struct {
		joltAdapters []int
		startRating  int
		ratingDiff   int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		diffCount1   int
		diffCount3   int
		combinations int
	}{
		{
			name:   "given",
			fields: fields{diffs: make(map[int]int)},
			args: args{
				joltAdapters: parseInput(input),
				startRating:  0,
				ratingDiff:   3,
			},
			diffCount1:   22,
			diffCount3:   10,
			combinations: 19208,
		},
		{
			name:   "given2",
			fields: fields{diffs: make(map[int]int)},
			args: args{
				joltAdapters: parseInput(input2),
				startRating:  0,
				ratingDiff:   3,
			},
			diffCount1:   7,
			diffCount3:   5,
			combinations: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jolter{
				diffs: tt.fields.diffs,
			}
			j.run(tt.args.joltAdapters, tt.args.startRating, tt.args.ratingDiff)
			if tt.diffCount1 != j.Diffs1() || tt.diffCount3 != j.Diffs3() {
				t.Errorf("got diff1 %d want %d, got diff3 %d want %d", j.Diffs1(), tt.diffCount1, j.Diffs3(), tt.diffCount3)
			}
			j.combos(tt.args.joltAdapters, 0)
			fmt.Println(j.combinations)
			if tt.combinations != j.Combinations() {
				t.Errorf("got combinations %d want %d", j.Combinations(), tt.combinations)
			}
		})
	}
}
