package main

import (
	"fmt"
	"math"
)

func sum(c chan float64, x float64) {  // Note: the channel is strictly typed, and must match the channel made and the argument's type.
	x = math.Pow(x, 2.0)
	c <- x  // Sends the value of x (argument) to chan c
}

func main() {
	c := make(chan float64)  // Like maps, channels must use make() before usage
	go sum(c, 5.0)  // Goroutine to calculate the sum.

	num := <-c  // Receives from channel c, of which the input is from the goroutine of sum()

	fmt.Printf("%v", num)
	fmt.Printf("%v", "Hello world!")
}
