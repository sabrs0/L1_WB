package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// почему указатель:
// Если структура большая - то это помогает избежать большого копирования данных

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func getDistBetweenPoints(p1, p2 *Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}
func main() {
	p1 := NewPoint(1.0, 2.0)
	p2 := NewPoint(-1.0, -2.0)
	fmt.Println(getDistBetweenPoints(p1, p2))
}
