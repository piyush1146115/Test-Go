package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
//func circleArea(c Circle) float64 {
//	return math.Pi * c.r*c.r
//}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r*c.r
}


type Person struct {
	Name string
}
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main7(){
	c := Circle{x: 0, y: 0, r: 5}

	//c := new(Circle)
	//var c Circle

	fmt.Println(c.x, c.y, c.r)

	c.x = 10
	c.y = 6

	fmt.Println(c.x, c.y, c.r)

	//fmt.Println(circleArea(c))
	fmt.Println(circleArea(&c))

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())


	b := new(Android)
	b.Person.Talk()

	a := new(Android)
	a.Talk()
}