package main

import "fmt"

func main4(){
	//var x [5]int
	//x[4] = 100
	//fmt.Println(x)

	//var x [5]float64
	//x[0] = 98
	//x[1] = 93
	//x[2] = 77
	//x[3] = 82
	//x[4] = 83

	//var total float64 = 0
	//for i := 0; i < 5; i++ {
	//	total += x[i]
	//}
	//fmt.Println(total / 5)

	//var total float64 = 0
	//for i := 0; i < len(x); i++ {
	//	total += x[i]
	//}
	////fmt.Println(total / len(x))
	//
	//fmt.Println(total/float64(len(x)))

	//var total float64 = 0
	//for i, value := range x {
	//	total += value
	//}
	//fmt.Println(total / float64(len(x)))



	//x := [5]float64{ 98, 93, 77, 82, 83 }
	//var total float64 = 0
	//for _, value := range x {
	//	total += value
	//}
	//fmt.Println(total / float64(len(x)))

	//slice1 := []int{1,2,3}
	//slice2 := make([]int, 2)
	//copy(slice2, slice1)
	//fmt.Println(slice1, slice2)

	//slice1 := []int{1,2,3}
	//slice2 := append(slice1, 4, 5)
	//fmt.Println(slice1, slice2)

	//x := make(map[string]int)
	//x["key"] = 10
	//fmt.Println(x["key"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])

	name, ok := elements["Un"]
	fmt.Println(name, ok)

	name1, ok1 := elements["Ne"]
	fmt.Println(name1, ok1)
}