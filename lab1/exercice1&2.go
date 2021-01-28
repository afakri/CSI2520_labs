package main

import (
	"fmt"
	"math"
)

func main() {
}

//exercice1
func round(x float64) (float64, float64) {
	f := math.Floor(x)
	c := math.Ceil(x)
	return f, c
}

//exercice2
func draw(x int) {
	for i := 0; i < x; i++ {
		for j := x; j > i; j-- {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("x")
		}
		fmt.Print("\n")

	}
	for i := 0; i < x; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" ")
		}
		for j := x; j > i; j-- {
			fmt.Print("x")
		}

		fmt.Print("\n")

	}

}
