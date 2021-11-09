package main

import "fmt"

type employee struct {
	name   string
	salary int
}

//From the above result, we can verify that even we have mutated the receiver object, it did not impact the original object on which the method was called.

//This proves that the method changeName received just a copy of the actual struct e(from main method). Hence any changes made to the copy inside the method
//did not affect the original struct.

//func (e Employee) changeName(newName string) {
//	e.name = newName
//}

//In the below definition, we instructed Go that this method will belong to the pointer of the Type instead of the value of the Type.
//When a method belongs to the pointer of a type, its receiver will receive the pointer to the object instead of a copy of the object.
func (e *employee) changeName(newName string) {
	(*e).name = newName
}

func main() {
	e := employee{
		name:   "Ross Geller",
		salary: 1200,
	}

	// e before name change
	fmt.Println("e before name change =", e)

	// change name
	//e.changeName("Monica Geller")

	// create pointer to `e`
	ep := &e
	// change name
	ep.changeName("Monica Geller")

	// e after name change
	fmt.Println("e after name change =", e)
}
