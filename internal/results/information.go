package results

import (
	"math"
)

// Information decompose an integer into power of 3
// 0 <= i < 3^5
// i = a x 3^4 + b x 3^3 + c x 3^2 + d x 3^1 + e x 3^0
func Information(i int) []int {
	res := []int{}
	if i >= int(math.Pow(3, 5)) || i < 0 {
		return res
	}

	for n := 4; n >= 0; n-- {
		a := i / int(math.Pow(3, float64(n)))
		res = append(res, a)
		i -= a * int(math.Pow(3, float64(n)))
	}

	return res
}
