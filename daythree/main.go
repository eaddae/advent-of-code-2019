package main

import (
	"fmt"
	"errors"
	"strings"
	"strconv"
)

type Direction int8
const (
	Left Direction = iota
	Up
	Right
	Down
)

type Point struct {
	x int
	y int
}

type Vector struct {
	dir Direction
	len int 
}

type Wire []Point

func main() {

	input := `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`
	wire_strings := strings.Split(input, "\n")
	var wires [][]string
	wires = append(wires, strings.Split(wire_strings[0], ","))
	wires = append(wires, strings.Split(wire_strings[1], ","))
	var w1, w2 Wire
	w1 = append(w1, Point{x:0,y:0})
	w2 = append(w2, Point{x:0,y:0})
	var vectors [2][]Vector
	for _, s := range wires[0] {
		vectors[0] = append(vectors[0], vectorize(s))
	}
	for _, s := range wires[1] {
		vectors[1] = append(vectors[1], vectorize(s))
	}
	w1, err := Draw(w1, vectors[0]...)
	w2, err = Draw(w2, vectors[1]...)
	if err != nil {
		fmt.Println(err)
	}
	intersections, err := w1.Intersects(w2)
	if err != nil {
		fmt.Println(err)
	}
	for _, x := range intersections {
		fmt.Println(DistOrigin(x))
	}
}

func vectorize(l string) Vector {
	var v Vector
	var len int
	var dir Direction
	switch l[0] {
	case 'D':
		dir = Down
	case 'U':
		dir = Up
	case 'L':
		dir = Left
	case 'R':
		dir = Right
	default:
		dir = -1
	}
	len, err := strconv.Atoi(l[1:])
	if err != nil {
		fmt.Println(err)
		return v
	}
	v.dir = dir
	v.len = len
	return v
}

func Draw(w Wire, vectors ...Vector) (result Wire, err error)  {
	new_wire := w[:]
	for _, v := range vectors {
		start := new_wire[len(new_wire) - 1]
		new_segment, err := Pointilize(start, v)
		if err != nil {
			return nil, err
		}
		new_wire = append(new_wire, new_segment...)
	}
	return new_wire, nil
}

func Pointilize(p Point, v Vector) (result Wire, err error) {
	horizontal, negative := false, false
	switch v.dir {
		case Left:
			horizontal = true
		case Up:
		case Right:
			horizontal, negative = true, true
		case Down:
			negative = true
		default:
			return nil, errors.New(fmt.Sprintf("Invalid value for vector direction:", v.dir))
	}
	result = append(result, p)
	for i := 0; i < v.len; i++ {
		start := result[len(result) - 1]
		var newp Point
		newp.x = start.x
		newp.y = start.y
		if horizontal {
			if negative {
				newp.x -= 1
			} else {
				newp.x += 1
			}
		} else {
			if negative {
				newp.y -= 1
			} else {
				newp.y += 1
			}
		}
		result = append(result, newp)
	}
	return result, nil
}

func (w1 Wire) Intersects(w2 Wire) (intersections []Point, err error) {
	
	for _, p := range w1 {
		for _, v := range w2 {
			if p.x == v.x && p.y == v.y {
				intersections = append(intersections, p)
			}
		}
	}
	return intersections, nil
}

func Distance(point1 Point, point2 Point) int {
	return abs(abs(point1.x - point2.x) + abs(point1.y - point2.y))
}

func DistOrigin(p Point) int {
	return abs(p.x) + abs(p.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
