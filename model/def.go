package model

import "math"

const Eps = 1E-16
const IterationCount = 10000

func F(choice int, x float64) float64{
	if choice == 1 {
		return math.Pow(x, 2) - 3.765
	} else if choice == 2 {
		return (1-x*x)*(1-x*x) - x //(1-x^2)^2-x
	} else if choice==3{
		return 3 - 5*x + x*x*x
	} else if choice==4{
		return math.Exp(x)-1-2*x
	} else if choice==5 {
		return math.Log(x)
	} else {
		return finance(x,18)
	}
}

func FDer(choice int, x float64) float64 {
	if choice == 1 {
		return 2 * x
	} else if choice == 2 {
		return -4*x*(1-x*x) - 1
	} else if choice==3 {
		return -5 + 3*x*x
	} else if choice ==4 {
		return math.Exp(x)-2
	} else if choice==5 {
		return 1/x
	} else {
		return financeDer(x, 18)
	}
}

func finance(x float64, times int) (float64){
	res := float64(1-times)
	iterX := x
	for i:=1; i<=11; i++{
		res += iterX
		iterX *= x
	}
	return res
}

func financeDer(x float64, times int) (float64){
	res :=1.0
	iterX :=x
	for i:=1; i<=10; i++{
		res += (float64(i)+1)*iterX
		iterX *= x
	}
	return res}