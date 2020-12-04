package main

import "testing"

const trees = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func Test_countTrees(t *testing.T) {
	type args struct {
		input string
		slope slope
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example",
			args: args{
				input: trees,
				slope: slope{right: 3, down: 1},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTrees(tt.args.input, tt.args.slope); got != tt.want {
				t.Errorf("countTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
