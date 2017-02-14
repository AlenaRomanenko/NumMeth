package model

import "fmt"
import "math"


func Calc(x0 float64, n int, a int) (float64, int)  {

	var corr = 1
	if a < 0 {
		a = (-1) * a
		corr = -1
	}

	var xPrev = x0

	var xNext = 1 / float64(n) * ((float64(n) - 1) + xPrev + float64(a)/math.Pow(xPrev, (float64)(n-1)))
	var i = 0

	for math.Abs(xNext-xPrev) > Eps && i < IterationCount {
		xPrev = xNext
		xNext = 1 / float64(n) * ((float64(n)-1)*xPrev + float64(a)/math.Pow(xPrev, (float64)(n-1)))
		fmt.Println("1", xPrev, "2", xNext)
		i++
	}

	if i == IterationCount {
		fmt.Println("Bad choice")
	} else {
		fmt.Println("root is ", float64(corr)*xPrev, "value of function: ", math.Pow(xPrev, float64(n)), ", count Of Iteration", i)
	}
	return float64(corr)*xPrev, i
}
