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
	input := "1,9,10,3,2,3,11,0,99,30,40,50"
	arbys, err := MapToIntSlice(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arbys)
}

func MapToIntSlice(program string) ([]int, error) {
	arby := strings.Split(program, ",")
	fmt.Println(arby)
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

/*func (program *[]int) RunAsIntcode() []int {
	for i := 0, o := program[i]; i < len(program); i += 4 {
		if o == Add {
			program[program[i + 3]] = program[i + 1] + program[i + 2]
		} else if o == Multiply {
			program[program[i + 3]] = program[i + 1] * program[i + 2]
		} else if o == Halt {
			return program
		}
	}
}
*/