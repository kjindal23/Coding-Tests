package main

import "fmt"

type triangle struct {
	i int64
	j int64
}

//func (t triangle) getArea()float64{}
type square struct {
	i int64
	j int64
}

//func(s square ) getArea()float64 {}

type shape interface {
	getArea() float64
}

//func (s shape) printArea()

func main() {
	//fmt.Println("my first program")

	TR := triangle{23, 12}
	SQ := square{10, 10}
	fmt.Println("The area of Triangle == ", printArea(TR))
	fmt.Println("The area of Square == ", printArea(SQ))

}

func printArea(s shape) float64 {
	//fmt.Println(s.getArea())
	return (s.getArea())
}

func (t triangle) getArea() float64 {
	//fmt.Println("the area of triangle ==")
	return float64(t.i * t.j)
}
func (t square) getArea() float64 {
	//fmt.Println("the area of square ==")
	return float64(t.i * t.j)
}
