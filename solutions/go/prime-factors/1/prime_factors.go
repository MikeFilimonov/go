package prime

func Factors(n int64) []int64 {

	var factors []int64
	startingInitialFactor := int64(2)
	factor := startingInitialFactor

	for n > 1 {

		if n%factor != 0 {
			factor++
			continue
		}

		n /= factor
		factors = append(factors, factor)
		factor = startingInitialFactor

	}

	return factors

}
