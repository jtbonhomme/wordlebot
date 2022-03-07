package wordle

import (
	"math"
)

// IntToPowerOf3 decompose an integer into power of 3
// The integer shall be comprise between 0 and 3^5
// Ex.:
// i = a x 3^4 + b x 3^3 + c x 3^2 + d x 3^1 + e x 3^0
// IntToPowerOf3(i) will return [a, b, c, d, e]
func IntToPowerOf3(i int) []int {
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
