package utils

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func CalculationResultPos(posArray []Point, length int, begin int, end int) []Point {
	textLen := posArray[1].X - posArray[0].X
	bx := begin * textLen / length
	ex := (length - end) * textLen / length
	pos := []Point{
		{X: posArray[0].X + bx, Y: posArray[0].Y},
		{X: posArray[1].X - ex, Y: posArray[0].Y},
		{X: posArray[1].X - ex, Y: posArray[3].Y},
		{X: posArray[0].X + bx, Y: posArray[3].Y},
	}
	return pos
}
