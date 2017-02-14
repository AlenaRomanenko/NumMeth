
package model

import "math"

const Eps = 1E-16
const IterationCount = 10000

func F(choice int, x float64) float64 {
	//return math.Pow(x, float64(n)) - float64(a)
	if choice == 1 {
		return math.Pow(x, 2) - 3.765
	} else if choice == 2 {
		return (1-x*x)*(1-x*x) - x //(1-x^2)^2-x
	} else if (choice==3){
		return 3 - 5*x + x*x*x
	} else{
		return math.Exp(x)-1-2*x
	}
}

func FDer(choice int, x float64) float64 {
	//return float64(n) * math.Pow(x, float64(n-1))
	if choice == 1 {
		return 2 * x
	} else if choice == 2 {
		return -4*x*(1-x*x) - 1
	} else if (choice==3) {
		return -5 + 3*x*x
	} else {
		return math.Exp(x)-2
	}
}