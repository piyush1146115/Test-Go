package main

import "fmt"

func main1()  {
	fmt.Println("Hello Go!")
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1 + 1 =", 1.0 + 1.0)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
