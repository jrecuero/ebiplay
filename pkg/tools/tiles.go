package tools

import "fmt"

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("%d,%d\n", p.X, p.Y)
}

func GetHorizontalMovableTiles(pos *Point, moveRange int) []*Point {
	var p *Point
	result := []*Point{pos}
	for i := 1; i <= moveRange; i++ {
		p = NewPoint(pos.X+i, pos.Y)
		result = append(result, p)
		p = NewPoint(pos.X-i, pos.Y)
		result = append(result, p)
	}
	return result
}

func GetMovableTiles(pos *Point, moveRange int) []*Point {
	var r []*Point
	result := GetHorizontalMovableTiles(pos, moveRange)
	for i := 1; i <= moveRange; i++ {
		r = GetHorizontalMovableTiles(NewPoint(pos.X, pos.Y+i), moveRange-i)
		result = append(result, r...)
		r = GetHorizontalMovableTiles(NewPoint(pos.X, pos.Y-i), moveRange-i)
		result = append(result, r...)
	}
	return result
}
