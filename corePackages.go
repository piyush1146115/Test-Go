package main

import (
	"container/list"
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age int
}

type ByName []person

func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	if this[i].Name == this[j].Name{
		return this[i].Age < this[j].Age
	}

	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}


type ByAge []person
func (this ByAge) Len() int {
	return len(this)
}
func (this ByAge) Less(i, j int) bool {
	if this[i].Age == this[j].Age{
		return this[i].Name < this[j].Name
	}

	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value.(int))
	}

	kids := []person{
		{"Jill",9},
		{"Jack",10},
		{"Jack",8},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}