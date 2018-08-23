package main

import (
	"fmt"
	"math"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func height(x1 float64, x2 float64) float64 {
	return x2 - x1
}
func percentageRelatveError(a float64, b float64) float64 {
	return ((a - b) / a) / 100
}
func printvalues(order int, value float64, correctValue float64) {
	fmt.Println(order, "order approximation:", value)
	fmt.Println("Correct value: ", correctValue)
	fmt.Println("Percent relative error: ", percentageRelatveError(correctValue, value), "%\n")
}
func main() {
	correctValue := math.Cos(math.Pi / 3)
	h := height(math.Pi/3, math.Pi/4)
	fmt.Println("h is", h, "\n")
	value := math.Cos(math.Pi / 4)
	printvalues(0, value, correctValue)

	//1st order
	value = value - math.Sin((math.Pi/4)*(math.Pi/12))
	printvalues(1, value, correctValue)
	// 2nd order
	value = value - ((math.Cos(math.Pi/4) / 2) * math.Pow(math.Pi/12, 2))
	printvalues(2, value, correctValue)
	// 3rd order
	value = value + math.Sin((math.Cos(math.Pi/4)/6)*math.Pow(math.Pi/12, 3))
	printvalues(3, value, correctValue)
	// 4th order
	value = value + math.Cos((math.Cos(math.Pi/4)/24)*(math.Pow(math.Pi/12, 4)))
	printvalues(4, value, correctValue)
	// 5th order
	value = value - math.Sin((math.Cos(math.Pi/4)/120)*math.Pow(math.Pi/12, 5))
	printvalues(5, value, correctValue)
	// 6th order
	value = value - math.Cos((math.Cos(math.Pi/4)/720)*math.Pow(math.Pi/12, 6))
	printvalues(6, value, correctValue)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	var  x,y,x1,xc,yc,xc1,yc1,xc2,yc2,xc3,yc3,xc4,yc4 float64=0,1.48,1,-.03,.37,.25,1.15,.075,.92,.50,1.35,1.15,1.70
	p.Title.Text = "Zero order Taylor Series and remainder!"
	p.X.Label.Text = "H==(Xi=0, Xi+1=1)"
	p.Y.Label.Text = "f(x)"

	err = plotutil.AddLinePoints(p,
		"Exact Prediction", Points(2, x, correctValue,x1,y),
		"First Order", Points(2, x,correctValue, x1, correctValue),
		"Dotted line1", Points(2, x1, x,x1, y),
		"Dotted line2", Points(2, x1, correctValue,y, correctValue),
		"Dotted line3", Points(2, x1, y,y, y),
                "Remainder", Points(2, y, x, y, y),
		"Curve", TruePoints(6,xc, yc, x,correctValue,xc2,yc2, xc1, yc1, xc3,yc3,x1,y,xc4,yc4))
	
	if err != nil {
		panic(err)
	}

	// Saves our plotted file to a png file
	if err := p.Save(12*vg.Inch, 12*vg.Inch, "plots1.png"); err != nil {
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
func TruePoints(n int, x, y,x1,y1, x2, y2,x3,y3,x4, y4,x5,y5, x6, y6 float64) plotter.XYs {
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
		}else if i==3{
			pts[i].X = x3
			pts[i].Y = y3
		}else if i==4{
			pts[i].X = x4
			pts[i].Y = y4
		}else if i==5{
			pts[i].X = x5
			pts[i].Y = y5
		}else{
			pts[i].X = x6
			pts[i].Y = y6
		}
		
	}
	return pts
}
