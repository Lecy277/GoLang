package main

import (
	"fmt"
	"math"
)

func equation(x float64) float64 {
	return -0.1*math.Pow(x, 4) - 0.15*math.Pow(x, 3) - 0.5*math.Pow(x, 2) - 0.25*(x) + 1.2
}
func et(x float64) float64 {
	v := -0.9125
	value := (v - x)
	if math.Signbit(value) {
		return value * -1
	}
	return value
}
func main() {
	correctValue := -0.4*math.Pow(0.5, 3) - 0.45*math.Pow(0.5, 2) - 1.0*(0.5) - 0.25
	fmt.Println("The correct of f`(0.5): ", correctValue)

	h := 10.0
	fmt.Println("\n Centered divided difference")
	for i := 2; i < 10; i++ {
		x := 0.5
		h = h / 10
		fmt.Println("h: ", h)
		value := (equation(x+h) - equation(x-h)) / (2 * h)
		fmt.Println("Difference ", value)
		fmt.Println("Et: ", et(value), "\n")
	}

}
