package main

import (
	"errors"
	"fmt"
	"math"
)

type NegError struct {
	num float64 // negative number
}
type Div0Error struct {
}

func rootDivN(num float64, n int) (res float64, err error) {
	if num < 0.0 {
		return 0.0, errors.New("= Main NegError: Negative number ")
	}
	if n == 0 {
		return 0.0, errors.New("Main Div0Error: Division by 0")
	}
	res = math.Sqrt(num) / float64(n)
	return
}

func main() {
	divs := []int{2, 10, 3, 0}
	nums := []float64{511.8, 0.65, -3.0, 2.123}
	for i, num := range nums {
		fmt.Printf("%d) sqrt(%f)/%d = ", i, num, divs[i])
		res, err := rootDivN(num, divs[i])
		if err == nil {
			fmt.Printf("%f\n", res)
		} else {
			if num < 0 {
				fmt.Printf("%s,%f", err, num)
				println()
			} else {
				fmt.Println(err)
			}

		}
	}
}
