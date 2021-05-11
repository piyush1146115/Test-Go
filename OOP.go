package main

import "fmt"

//type Rect struct {
//	Name 	string
//	Width, Height float64
//}
//
//type base struct {
//	a string
//	b int
//}
//
//type derived struct {
//	base //Embedding
//	d int
//	a float32
//}


//Multiple embedding
//type Nameobj struct{
//	name string
//}
//
//type Shape struct{
//	Nameobj //inheritance
//	color int32
//	isRegular bool
//}
//
//type Point struct{
//	x, y float64
//}
//
//type RectAngle struct{
//	Nameobj //multiple inheritance
//	Shape
//	center	Point  //Standard composition
//	Width, Height float64
//}



//Method Shadowing
//Since all golang-struct methods are non-virtual, you cannot override methods (you need interfaces for that)
type base struct {
	a string
	b int
}

//method xyz
func (this base) xyz(){
	fmt.Println("xyz, a is :", this.a)
}

//method display
func (this base) display(){
	fmt.Println("base a is: ", this.a)
}

type derived struct{
	base // embedding
	d int
	a float32
}

//method display - Shadowed
func (this derived) display(){
	fmt.Println("Derived a is : ", this.a)
}

func main(){
	//var a Rect
	//var b = Rect{"I am B", 10, 20}
	//var c = Rect{Height: 12, Width: 14}
	//
	////println("Hello")
	//Println(a)
	//Println(b)
	//Println(c)

	//var x derived
	//fmt.Printf("%T\n", x.a)
	//fmt.Printf("%T\n", x.base.a)


	//Multiple embedding
	//var aRect = RectAngle{
	//	Nameobj{"name1"},
	//	Shape{Nameobj{"name2"}, 0, true},
	//	Point{0, 0},
	//	20, 2.5,
	//}
	//
	//fmt.Println(aRect.name)
	//fmt.Println(aRect.Shape)
	//fmt.Println(aRect.Shape.name)


	//Method Shadowing
	var a derived = derived{base{"base-a", 10}, 20, 2.5}
	a.display() // calls Derived/display(a)
	// => "derived, a is: 2.5"
	a.base.display() // calls Base/display(a.base), the base implementation
	// => "base, a is: base-a"
	a.xyz() // "xyz" was not shadowed, calls Base/xyz(a.base)
	// => "xyz, a is: base-a"

}
