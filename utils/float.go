package utils

import (
	"fmt"
	"math"
)

func isClose(a, b float64) bool {
	rel_tol := 1e-09
	abs_tol := 0.0

	fmt.Println(a, b, math.Abs(a-b), math.Max(math.Abs(a), math.Abs(b)), rel_tol*math.Max(math.Abs(a), math.Abs(b)), abs_tol, math.Max(
		rel_tol*math.Max(math.Abs(a), math.Abs(b)),
		abs_tol,
	))

	return math.Abs(a-b) <= math.Max(
		rel_tol*math.Max(math.Abs(a), math.Abs(b)),
		abs_tol,
	)
}
