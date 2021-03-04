package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f() (int, int) {
	return 5, 6
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func main() {
	x := 0
	//var y uint
	//Closure
	increment := func() int {
		x++
		return x
}
	defer second()
	fmt.Println(increment())
	fmt.Println(increment())

	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))

	fmt.Println(f())

	first()

	//fmt.Print("Enter a number: ")
	//fmt.Scanf("%u", y)
	//fmt.Println(factorial(y))

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}