package utils

import (
	"fmt"
	"strconv"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func CalculationResultPos(posArray []Point, length int, begin int, end int) []Point {
	textLen := posArray[1].X - posArray[0].X
	bx := float64(begin) * textLen / float64(length)
	ex := float64(length-end) * textLen / float64(length)
	pos := []Point{
		{X: decimal(posArray[0].X + bx), Y: decimal(posArray[0].Y)},
		{X: decimal(posArray[1].X - ex), Y: decimal(posArray[0].Y)},
		{X: decimal(posArray[1].X - ex), Y: decimal(posArray[3].Y)},
		{X: decimal(posArray[0].X + bx), Y: decimal(posArray[3].Y)},
	}
	return pos
}

func decimal(value float64) float64 {
	float, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return float
}
