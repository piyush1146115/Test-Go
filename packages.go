package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main10() {
	// Finding the Max of two numbers
	fmt.Println(math.Max(73.15, 92.46))

	// Calculate the square root of a number
	fmt.Println(math.Sqrt(225))

	// Printing the value of `𝜋`
	fmt.Println(math.Pi)

	// MaxInt64 is an exported name
	fmt.Println("Max value of int64: ", int64(math.MaxInt64))

	// Phi is an exported name
	fmt.Println("Value of Phi (ϕ): ", math.Phi)

	// Epoch time in milliseconds
	epoch := time.Now().Unix()
	fmt.Println(epoch)

	// Generating a random integer between 0 to 100
	rand.Seed(epoch)
	fmt.Println(rand.Intn(100))
}