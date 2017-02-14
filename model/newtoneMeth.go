package model

import "fmt"
import "math"



func CalcNewtone(choice int, x0 float64) (float64, int) {

	var xPrev = x0

	var xNext = xPrev - F(choice, xPrev)/FDer(choice,xPrev)
	var i = 0

	for math.Abs(xNext-xPrev) > Eps && i < IterationCount {
		xPrev = xNext
		xNext = xPrev - F(choice, xPrev)/FDer(choice,xPrev)

		fmt.Println("***** ", xPrev, "*****: ", xNext ) 
		i++
	}

	if i == IterationCount {
		fmt.Println("Bad choice")
	} else {
		fmt.Println("root is ", xPrev, "value of function: ", F(choice, xPrev), ", count Of Iteration", i)
	}
	return xPrev, i
}
