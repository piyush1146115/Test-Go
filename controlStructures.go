package main

import "fmt"

func main3(){
	//i := 1
	//for i <= 10 {
	//	fmt.Println(i)
	//	i = i + 1
	//}

	//for i := 1; i <= 10; i++ {
	//	if i % 2 == 0 {
	//		fmt.Println(i, "even")
	//	} else {
	//		fmt.Println(i, "odd")
	//	}
	//}

	//for i := 1; i <= 10; i++ {
	//	if i % 2 == 0 {
	//		fmt.Println(i, "even")
	//	} else {
	//		fmt.Println(i, "odd")
	//	}
	//}

	var i int
	fmt.Print("Enter a number:")
	fmt.Scanf("%d", &i)

	switch i {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	default: fmt.Println("Unknown Number")
	}
}
