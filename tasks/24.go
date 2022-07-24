package tasks

import "math"

func Task24() {
	var p1 Point = *NewPoint(5, 5)
	var p2 Point = *NewPoint(10, 10)

	var result float64 = p1.distance(p2)

	PrintResult(24, "Расстояние между точками", result)
}

// структура с 2 инкапсулированными параметрами
type Point struct {
	x int
	y int
}

// Конструктор
func NewPoint(x int, y int) *Point {
	return &Point{x, y}
}

// Формула расчета расстояния 2 точек
func (p1 *Point) distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x)-float64(p1.x), 2) + math.Pow(float64(p2.y)-float64(p1.y), 2))
}
