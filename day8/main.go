package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	instructions := parseInstructions(input)

	acc := runner(instructions)

	fmt.Printf("Number of instructions: %d\n", len(instructions))
	fmt.Printf("acc value %d run\n", acc)
}

type instruction struct {
	operation    string
	signedNumber int
}

func runner(instructions []instruction) int {
	for index, i := range instructions {

		if i.operation == "acc" {
			continue
		}

		if i.operation == "nop" {
			instructions[index] = instruction{
				operation:    "jmp",
				signedNumber: i.signedNumber,
			}
		}

		if i.operation == "jmp" {
			instructions[index] = instruction{
				operation:    "nop",
				signedNumber: i.signedNumber,
			}
		}

		accumulator := 0
		executedInstructions := make(map[int]instruction)
		acc, success := run(instructions, executedInstructions, 0, accumulator)
		if success {
			fmt.Println("success!")
			return acc
		}

		// put back original
		instructions[index] = i
	}

	return -1
}

func parseInstructions(input string) []instruction {
	lines := strings.Split(input, "\n")

	instructions := []instruction{}
	for _, line := range lines {
		p := strings.Split(line, " ")
		if len(p) != 2 {
			fmt.Printf("invalid instruction: %s\n", p)
		}
		n, _ := strconv.Atoi(p[1])
		i := instruction{
			operation:    p[0],
			signedNumber: n,
		}
		//fmt.Println(i)
		instructions = append(instructions, i)
	}

	return instructions
}

func run(instructions []instruction, executedInstructions map[int]instruction, current int, accumulator int) (int, bool) {
	instruction := instructions[current]
	var nextInstruction int
	if instruction.operation == "nop" {
		fmt.Println("noop")
		nextInstruction = current + 1
	}

	if instruction.operation == "acc" {
		fmt.Printf("add %d to accumulator %d.\n", instruction.signedNumber, accumulator)
		accumulator = accumulator + instruction.signedNumber
		fmt.Printf("new value: %d\n", accumulator)
		nextInstruction = current + 1
	}

	if instruction.operation == "jmp" {
		fmt.Printf("jumping %d.\n", instruction.signedNumber)
		nextInstruction = current + instruction.signedNumber
	}

	executedInstructions[current] = instruction
	if _, ok := executedInstructions[nextInstruction]; ok {
		return accumulator, false
	}

	if nextInstruction == len(instructions) {
		return accumulator, true
	}

	return run(instructions, executedInstructions, nextInstruction, accumulator)
}
