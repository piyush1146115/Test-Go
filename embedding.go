package main

import "fmt"

type Par struct{
value int64
}

func (i *Par) Value() int64{
	return i.value
}

type Child struct{
	Par
	multiplier int64
}

func (i Child) Value() int64{
	return i.value * i.multiplier
}

type Valueable interface {
	Value() int64
}


type Bitmap struct{
	data [4][4]bool
}

type Renderer struct{
	*Bitmap //Embed by pointer
	on uint8
	off uint8
}

func (r *Renderer)render() {
	for i := range(r.data){
		for j := range(r.data[i]){
			if r.data[i][j] {
				fmt.Printf("%c",r.on)
			} else {
				fmt.Printf("%c",r.off)
			}
		}
		fmt.Printf("\n")
	}
}

func main(){

	var renderA,renderB Renderer
	renderA.on = 'X'
	renderA.off = 'O'
	renderB.on = '@'
	renderB.off = '.'

	var pic Bitmap
	pic.data[0][1] = true
	pic.data[0][2] = true
	pic.data[1][1] = true
	pic.data[2][1] = true
	pic.data[3][1] = true

	renderA.Bitmap = &pic
	renderB.Bitmap = &pic

	renderA.render()
	renderB.render()
}