package summultiples

func SumMultiples(limit int, divisors ...int) int {

	scores := make(map[int]bool)

	for _, v := range divisors {

		if v == 0 {
			continue
		}

		for i := v; i < limit; i += v {
			scores[i] = true
		}

	}

	result := 0
	for k := range scores {
		result += k
	}

	return result

}
