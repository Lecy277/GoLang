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
	value := 1.0
	value = (x * (1 / math.Pow(math.Cos(x), 2))) / math.Tan(x)
	return value
}
func main() {
	fmt.Println("For x: Pi/2 + 0.1 (Pi/2)")
	arg := (math.Pi / 2) + (0.1 * (math.Pi / 2))
	fmt.Println("Conditional number ", equation(arg))

	fmt.Println("\nFor x: Pi/2 + 0.01 (Pi/2)")
	arg = (math.Pi / 2) + (0.01 * (math.Pi / 2))
	fmt.Println("Conditional number ", equation(arg))
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	var x,x1,x2,x3,x4,x5,x6,x7,x9 float64=-.50,-.25,0,.55,.75,.85,1.00,1.25,1.45
	p.Title.Text = "Error Propagation Figure 4.7"
	p.X.Label.Text = "<------Tri x'--------->"
	p.Y.Label.Text = "<--------|f'(x)|tri x   Estimated error------> "

	err = plotutil.AddLinePoints(p,
		"Dotted line True error1", Points(2, x,x6, x9, x6),
		"Dotted line True error2", Points(2, x6,x5, x9, x5),
                "Dotted line Down1", Points(2, x6,x5, x6, x2),
                "Dotted line Down2", Points(2, x2,x6, x2, x2),
                "Estimated error line", Points(2, x,x4, x6,x4),
		"Error line 1", TruePoints(4,x1,x7,x2,x6,x6,x4,x7,x3),
		"Error line 2", TruePoints(4,x1,x9,x2,x6,x6,x5,x7,x4))

	
	if err != nil {
		panic(err)
	}

	// Saves our plotted file to a png file
	if err := p.Save(10*vg.Inch, 10*vg.Inch, "plots2.png"); err != nil {
		panic(err)
	}


}

//Plot out points
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