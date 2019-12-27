package main

import (
	"fmt"
	"errors"
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
	
	/*for _, _ := range w1 {

	}*/
	return nil, errors.New("not implemented")
}

func Distance(point1 Point, point2 Point) int {
	return abs(abs(point1.x - point2.x) + abs(point1.y - point2.y))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {

}
