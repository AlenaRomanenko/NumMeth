package service

import "fmt"
import (
	"math"
	"NumMeth/model"
)


func CalcNewtone(choice int, x0 float64) (float64, int, float64) {

	var xPrev = x0

	if model.FDer(choice,xPrev) == 0  {
		fmt.Println("Bad choice")
		return 0,-1,0
	}

	var xNext = xPrev - model.F(choice, xPrev)/model.FDer(choice,xPrev)
	var i = 0

	for math.Abs(xNext-xPrev) > model.Eps && i < model.IterationCount {
		xPrev = xNext
		xNext = xPrev - model.F(choice, xPrev)/model.FDer(choice,xPrev)

		fmt.Println("***** ", xPrev, "*****: ", xNext ) 
		i++
	}

	if i == model.IterationCount {
		fmt.Println("Bad choice")
	} else {
		fmt.Println("root is ", xPrev, "value of function: ", model.F(choice, xPrev), ", count Of Iteration", i)
	}
	return xPrev, i, model.F(choice,xPrev)
}

