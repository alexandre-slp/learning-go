package mymath

import "sort"

// CenteredAvg computes the average of a list of numbers
// after removing the largest and smallest values.
func CenteredAvg(xi []int) float64 {
	if len(xi) == 0 {
		return 0
	}

	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f
}
