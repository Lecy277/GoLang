package main

import (
	"fmt"
	"math"
)

func et(x float64) float64 {
	v := -0.9125
	value := ((v - x) / v) * 100
	if value < 1 {
		return value * -1
	}
	return value
}
func equation1(x float64) float64 {
	return -0.1*math.Pow(x, 4) - 0.15*math.Pow(x, 3) - 0.5*math.Pow(x, 2) - 0.25*(x) + 1.2
}

func approximation(x float64, y float64) float64 {
	return x - y
}
func equation(x float64, m float64) float64 {
	return math.Pow(x, m)
}
func printvalues() {
	fmt.Println("x=0: ", equation(2.4, 3.3))
	fmt.Println("x=0.5: ", equation(5, 43))
	fmt.Println("x=1: ", equation(1.3, 0.4))
}
func main() {
	v := 1.0
	f2 := v + v*(v)
	fmt.Println("For 1: ", f2)
	r1 := f2 / f2 * (v)
	fmt.Println("R1: ", r1)
	f2 = 1 + 3*(v*v)*v
	fmt.Println("For 1: ", f2)
	r1 = (f2*3/f2)*(v*v) + (f2*3/f2*3)*(v*v)
	fmt.Println("R1: ", r1)
	value := r1
	value = value - math.Sin((math.Pi/4)*(math.Pi/12))
	printvalues()
}
