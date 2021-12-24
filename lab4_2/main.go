package main

import (
	"fmt"
	"math"
)

func main() {
	CreateTriangle(0, 0, 2, 5, 6, 1)
}

func CreateTriangle(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int) interface{} {
	figure := Triangle{x1: x1, y1: y1, x2: x2, y2: y2, x3: x3, y3: y3}

	figure.PrintPerimeter()
	figure.printSquare()
	figure.containsDot(2, 2)

	return figure
}

type Triangle struct {
	x1, x2, x3, y1, y2, y3 int
	l1, l2, l3, p, s       float64
}

func (fig *Triangle) GetPerimeter() {
	fig.GetParties()
	fig.p = fig.l1 + fig.l2 + fig.l3
}

func (fig *Triangle) PrintPerimeter() {

	fig.GetPerimeter()
	fmt.Println("Периметр треугольника = ", fig.p)
}

func GetPerimeter(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int) float64 {

	l1 := math.Sqrt(math.Pow(float64(y2-y1), 2) + math.Pow(float64(x2-x1), 2))
	l2 := math.Sqrt(math.Pow(float64(y3-y2), 2) + math.Pow(float64(x3-x2), 2))
	l3 := math.Sqrt(math.Pow(float64(y1-y3), 2) + math.Pow(float64(x1-x3), 2))

	p := (l1 + l2 + l3) / 2
	s := math.Sqrt(p * (p - l1) * (p - l2) * (p - l3))

	return s
}

func (fig *Triangle) GetSquare() {

	fig.GetParties()

	p := (fig.l1 + fig.l2 + fig.l3) / 2

	fig.s = math.Sqrt(p * (p - fig.l1) * (p - fig.l2) * (p - fig.l3))
}

func (fig *Triangle) printSquare() {

	fig.GetSquare()
	fmt.Println("Площадь треугольника = ", fig.s)
}

func (fig *Triangle) GetParties() {

	fig.l1 = math.Sqrt(math.Pow(float64(fig.y2-fig.y1), 2) + math.Pow(float64(fig.x2-fig.x1), 2))
	fig.l2 = math.Sqrt(math.Pow(float64(fig.y3-fig.y2), 2) + math.Pow(float64(fig.x3-fig.x2), 2))
	fig.l3 = math.Sqrt(math.Pow(float64(fig.y1-fig.y3), 2) + math.Pow(float64(fig.x1-fig.x3), 2))
}

func (fig *Triangle) containsDot(dotx int, doty int) {
	// от точки к каждой паре вершин строим треугольники. Если их сумма равна пл треугольника, то true, иначе false
	fig.GetParties()
	fig.GetSquare()

	s1 := GetPerimeter(fig.x1, fig.y1, fig.x2, fig.y2, dotx, doty)
	s2 := GetPerimeter(fig.x2, fig.y2, fig.x3, fig.y3, dotx, doty)
	s3 := GetPerimeter(fig.x3, fig.y3, fig.x1, fig.y1, dotx, doty)

	if (math.Ceil(s1*10)/10 + math.Ceil(s2*10)/10 + math.Ceil(s3*10)/10) == math.Ceil(fig.s*10)/10 {
		fmt.Println("Точка (", dotx, ", ", doty, ") лежит в треугольнике")
	} else {
		fmt.Println("Точка (", dotx, ", ", doty, ") не лежит в треугольнике")
	}

}
