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
	value := ((v - x) / v) * 100
	if value < 1 {
		return value * -1
	}
	return value
}
func main() {
	correctValue := -0.4*math.Pow(0.5, 3) - 0.45*math.Pow(0.5, 2) - 1.0*(0.5) - 0.25
	fmt.Println("The correct of f`(0.5): ", correctValue)

	fmt.Println("\nWith step size 0.5: ")
	fmt.Println("x=0: ", equation(0))
	fmt.Println("x=0.5: ", equation(0.5))
	fmt.Println("x=1: ", equation(1))

	fmt.Println("\n Forward divided difference")
	x := 0.5
	value := (equation(1) - equation(0.5)) / x
	fmt.Println("f`(", x, ")= ", value)
	fmt.Println("Et: ", et(value), "%")

	fmt.Println("\n Bakward divided difference")
	x = 0.5
	value = (equation(0.5) - equation(0)) / x
	fmt.Println("f`(", x, ")= ", value)
	fmt.Println("Et: ", et(value), "%")

	fmt.Println("\n Bakward divided difference")
	x = 1.0
	value = (equation(1.0) - equation(0)) / x
	fmt.Println("f`(", x, ")= ", value)
	fmt.Println("Et: ", et(value), "%")

	fmt.Println("\n\nWith step size 0.5: ")
	fmt.Println("x=0.5: ", equation(0.5))
	fmt.Println("x=0.25: ", equation(0.25))
	fmt.Println("x=0.75: ", equation(0.75))

	fmt.Println("\n Forward divided difference")
	x = 0.25
	value = (equation(0.75) - equation(0.5)) / x
	fmt.Println("f`(", x, ")= ", value)
	fmt.Println("Et: ", et(value), "%")

	fmt.Println("\n Bakward divided difference")
	x = 0.25
	value = (equation(0.5) - equation(0.25)) / x
	fmt.Println("f`(", x, ")= ", value)
	fmt.Println("Et: ", et(value), "%")

	fmt.Println("\n Bakward divided difference")
	x = 0.5
	value = (equation(0.75) - equation(0.25)) / x
	fmt.Println("f`(", x, ")= ", value)
	fmt.Println("Et: ", et(value), "%")

}
