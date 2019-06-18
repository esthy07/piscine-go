package piscine

import "math"

func RecursiveFactorial(x int) int {
	if 0 > x {
		return 0
	} else if 0 == x || 1 == x{
		return 1
	} else {
		res := x*RecursiveFactorial(x-1) > math.MaxInt32
		if res {
			return 0
		} else {
			return x * RecursviveFactorial(x-1)
		}	
	}
}

