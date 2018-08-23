package main

import (
	"fmt"
	"math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func equation(x float64) float64 {
	return -0.1*math.Pow(x, 4) - 0.15*math.Pow(x, 3) - 0.5*math.Pow(x, 2) - 0.25*(x) + 1.2
}
func approximation(x float64, y float64) float64 {
	return x - y
}
func main() {
	var  x,y2 float64=0,.45
	Et := approximation(0.2, 1.2)
	fmt.Println("Et: ", Et)

	fmt.Println("At x=0 for n=1:")
	value := -0.4*math.Pow(0.0, 3) - 0.45*math.Pow(0.0, 2) - 1.0*(0.0) - 0.25
	fmt.Println("For n=1: ", value)
	fmt.Println("First order approximation equation: ", "1.2", value, "h")
	f1 := 1.2 + (value)
	fmt.Println("Correct value at x=1 is: ", f1)
	fmt.Println("First order truncation error: ", approximation(0.2, f1))

	value2 := -1.2*math.Pow(0.0, 2) - 0.9*0.0 - 1.0
	fmt.Println("\nFor n=2: ", value2)
	fmt.Println("First order approximation equation: ", "1.2", value, "h", value2, "h^2")
	f2 := 1.2 + (value) + value2
	fmt.Println("Correct value at x=1 is: ", f2)
	fmt.Println("First order truncation error: ", approximation(0.2, f2))

	// value3 := 1.2 - 0.25 - 0.5*math.Pow(1, 2) - 0.15*math.Pow(1, 3) - 0.1*math.Pow(1, 4) - 1.0
	value3 := equation(1)
	fmt.Println("For x=1: ", value3)
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	

	p.Title.Text = "The Taylor Series!"
	p.X.Label.Text = "H==(Xi=0, Xi+1=1)"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"Zero order", Points(2, x, math.Abs(value+value2)+f2,math.Abs(value2),math.Abs(value+value2)+f2),
		"First Order", Points(2, x, math.Abs(value+value2)+f2, math.Abs(value2), f1),
		"Second Order", Points(2, x, math.Abs(value+value2)+f2, math.Abs(value2), y2),
		"True", TruePoints(4,value, math.Abs(value+value2), x, math.Abs(value+value2)+f2,math.Abs(value2), value3, math.Abs(value+value2)+f2, x ))
		
	
	if err != nil {
		panic(err)
	}

	// Saves our plotted file to a png file
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "plots.png"); err != nil {
		panic(err)
	}

}
///Plot out points
func Points(n int, x, y,x1,y1 float64) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
	       if i==0{
			pts[i].X = x
			pts[i].Y = y
		}else{
			pts[i].X = x1
			pts[i].Y = y1
		}
	}
	return pts
}

//plot points for the true line
func TruePoints(n int, x, y,x1,y1, x2, y2,x3,y3 float64) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i==0{
			pts[i].X = x
			pts[i].Y = y
		}else if i==1{
			pts[i].X = x1
			pts[i].Y = y1
		}else if i==2{
			pts[i].X = x2
			pts[i].Y = y2
		}else{
			pts[i].X = x3
			pts[i].Y = y3
		}
	}
	return pts
}



