package main

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	Add int = 1
	Multiply int = 2
	Halt int = 99
)

func main() {
	init := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,19,5,23,2,9,23,27,1,5,27,31,1,5,31,35,1,35,13,39,1,39,9,43,1,5,43,47,1,47,6,51,1,51,13,55,1,55,9,59,1,59,13,63,2,63,13,67,1,67,10,71,1,71,6,75,2,10,75,79,2,10,79,83,1,5,83,87,2,6,87,91,1,91,6,95,1,95,13,99,2,99,13,103,1,103,9,107,1,10,107,111,2,111,13,115,1,10,115,119,1,10,119,123,2,13,123,127,2,6,127,131,1,13,131,135,1,135,2,139,1,139,6,0,99,2,0,14,0"
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; i++ {
			arbys, err := MapToIntSlice(init)
			if err != nil {
				fmt.Println(err)
			}
			if TestInputs(arbys, i, j) {
				fmt.Println("BINGO!",i,j)
				break
			}
		}
	}
}

func TestInputs(program []int, first int, second int) bool {
	wanted_result := 19690720
	program[1] = first
	program[2] = second
	return RunIntcode(program) == wanted_result
}

func MapToIntSlice(program string) ([]int, error) {
	arby := strings.Split(program, ",")
	var arbynums []int
	for _, s := range arby {
		si, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		arbynums = append(arbynums, si)
	}
	return arbynums, nil
}

func RunIntcode(program []int) int {
	for i := 0; i < len(program); i += 4 {
		o := program[i]
		if o == Add {
			program[program[i + 3]] = program[program[i + 1]] + program[program[i + 2]]
		} else if o == Multiply {
			program[program[i + 3]] = program[program[i + 1]] * program[program[i + 2]]
		} else if o == Halt {
			return program[0]
		}
	}
	return program[0]
}
