package interval

import "golang.org/x/exp/constraints"

/*
	Public Function that returns the smallest value from input.

	Parameters:
		elements []T{} slice of type T (Ordered)

	Returns:
		minima T

*/
func Min[T constraints.Ordered](elements ...T) (minima T) {
	if len(elements) < 1 {
		panic("Not enough arguments based for evaluation")
	} else {
		minima = elements[0]
		for _, element := range elements {
			if minima > element {
				minima = element
			}
		}
		return
	}
}

/*
	Public Function that returns the biggest value from input.

	Parameters:
		elements []T{} slice of type T (Ordered)

	Returns:
		maxima T

*/
func Max[T constraints.Ordered](elements ...T) (maxima T) {
	if len(elements) < 1 {
		panic("Not enough arguments based for evaluation")
	} else {
		maxima = elements[0]
		for _, element := range elements {
			if maxima < element {
				maxima = element
			}
		}
		return
	}
}
