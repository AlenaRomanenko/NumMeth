package model

import "fmt"
import "math"
//import "gowiki/model"



func CalcNewtonePol(choice int, x0 float64, c float64, d float64) (float64, int) {
	
   

	var xPrev = x0
	if FDer(choice,xPrev) == 0 {
		fmt.Println("Bad choice")
		return 0,0
	}
	var xNext = xPrev - F(choice,xPrev)/(FDer(choice,xPrev)-d/(c-xPrev))

    // fmt.Println("root is ", xPrev, "value of function: ", xNext)
	var i = 0

	for math.Abs(xNext-xPrev) > Eps && i < IterationCount {
		xPrev = xNext
		xNext = xPrev - F(choice,xPrev)/(FDer(choice, xPrev)-d/(c-xPrev))
       fmt.Println("x_k", xPrev, "x_k+1: ", xNext)
		i++
	}
   
   if i == IterationCount {
		fmt.Println("Bad choice")
	} else {
		fmt.Println("root is ", xPrev, "value of function: ", F(choice,xPrev), ", count Of Iteration", i)
	}
	return xNext, i
}
