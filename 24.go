package main

import (
	"math"
)

type point struct {
	x float64
	y float64
}

func newPoint(x, y float64) point {
	return point{x: x, y: y}
}

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
func distance(first, second point) float64 {
	firstCath := math.Abs(first.x - second.x)
	secondCath := math.Abs(first.y - second.y)

	res := math.Sqrt(math.Pow(firstCath, 2) + math.Pow(secondCath, 2))
	return res
}
