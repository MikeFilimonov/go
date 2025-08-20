package prime

func Factors(n int64) []int64 {

	var factors []int64
	factor := int64(2)

	for n > 1 {

		for n%factor == 0 {
			factors = append(factors, factor)
			n /= factor
		}

		factor++

	}

	return factors

}
