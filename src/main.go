package main

import (
	"fmt"
	"math"
	"os"
)

func sum(c chan float64, x float64) { // Note: the channel is strictly typed, and must match the channel made and the argument's type.
	x = math.Pow(x, 2.0)
	c <- x // Sends the value of x (srgument) to chan c
}

func expect(e error) {
	// Takes in an error, and panics if the error is not nil.
	if e != nil {
		panic(e)
	}
}

func main() {
	c := make(chan float64) // Like maps, channels must use make() before usage
	
	go sum(c, 5.0)          // Goroutine to calculate the sum.
	go sum(c, 6.0)

	dat, err := os.ReadFile("./lament.txt") // Reads from a file
	expect(err)

	fmt.Print(string(dat))

	num1, num2 := <-c, <-c // Receives from channel c, of which the input is from the goroutine of sum()

	fmt.Printf("\n%v, %v", num1, num2)
	fmt.Printf("\n%v", "Hello world!")
}
