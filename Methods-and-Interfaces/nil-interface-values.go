//A nil interface value holds neither value nor concrete type.
//
//Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
package main

type In interface {
	M()
}

func main() {
	var i In
	describe(i)
	i.M()
}
