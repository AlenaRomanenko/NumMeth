package service

import "fmt"
import "math"
import "NumMeth/model"



func CalcNewtonePol(choice int, x0 float64, c float64, d float64) (float64, int, float64) {

	var xPrev = x0
	if xPrev == c || model.FDer(choice,xPrev)-d/(c-xPrev) == 0  {
		fmt.Println("Bad choice")
		return 0,-1,0
	}
	var xNext = xPrev - model.F(choice,xPrev)/(model.FDer(choice,xPrev)-d/(c-xPrev))

    // fmt.Println("root is ", xPrev, "value of function: ", xNext)
	var i = 0

	for math.Abs(xNext-xPrev) > model.Eps && i < model.IterationCount {
		xPrev = xNext
		xNext = xPrev - model.F(choice,xPrev)/(model.FDer(choice, xPrev)-d/(c-xPrev))
       fmt.Println("x_k", xPrev, "x_k+1: ", xNext)
		i++
	}
   
   if i == model.IterationCount {
		fmt.Println("Bad choice")
	} else {
		fmt.Println("root is ", xPrev, "value of function: ", model.F(choice,xPrev), ", count Of Iteration", i)
	}
	return xNext, i, model.F(choice, xPrev)
}


func CalcFlyNewtonePol(choice int, x0 float64) (float64, int, float64) {

	var xPrev = x0
	if model.FDer(choice,xPrev) == 0 {
		fmt.Println("Bad choice")
		return 0,0,0
	}

	d :=-model.F(choice, xPrev)
	c:= xPrev - (model.F(choice, xPrev) +2)/ model.FDer(choice, xPrev)

	var xNext = xPrev - model.F(choice,xPrev)/(model.FDer(choice,xPrev)-d/(c-xPrev))

	// fmt.Println("root is ", xPrev, "value of function: ", xNext)
	var i = 0

	for math.Abs(xNext-xPrev) > model.Eps && i < model.IterationCount {
		xPrev = xNext

		d :=-model.F(choice, xPrev)
		c:= xPrev - (model.F(choice, xPrev) +2)/ model.FDer(choice, xPrev)

		xNext = xPrev - model.F(choice,xPrev)/(model.FDer(choice, xPrev)-d/(c-xPrev))
		fmt.Println("x_k", xPrev, "x_k+1: ", xNext)
		i++
	}

	if i == model.IterationCount {
		fmt.Println("Bad choice")
	} else {
		fmt.Println("root is ", xPrev, "value of function: ", model.F(choice,xPrev), ", count Of Iteration", i)
	}
	return xNext, i, model.F(choice, xPrev)
}
