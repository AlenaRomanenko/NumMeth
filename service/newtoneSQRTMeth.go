package service

import "fmt"
import (
	"math"
	"NumMeth/model"
)


func CalcSQRT(n int, a int) (float64, int, float64)  {
	var corr = 1
	if n % 2 == 0 && a <0 {
		return 0.0,-1,0.0
	}
	if a < 0 {
		a = (-1) * a
		corr = -1
	}

	var xPrev = float64(a)

	var xNext = 1 / float64(n) * ((float64(n) - 1) + xPrev + float64(a)/math.Pow(xPrev, (float64)(n-1)))
	var i = 0

	for math.Abs(xNext-xPrev) > model.Eps && i < model.IterationCount {
		xPrev = xNext
		xNext = 1 / float64(n) * ((float64(n)-1)*xPrev + float64(a)/math.Pow(xPrev, (float64)(n-1)))
		fmt.Println("1", xPrev, "2", xNext)
		i++
	}

	if i == model.IterationCount {
		fmt.Println("Оберіть інші початкові значення")
	} else {
		fmt.Println("Корінь: ", float64(corr)*xPrev, "value of function: ", math.Pow(xPrev, float64(n)), ", Кількість ітерацій", i)
	}
	return float64(corr)*xPrev, i, math.Pow(xPrev, float64(n))-float64(a)
}
