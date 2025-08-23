package series

func All(n int, s string) []string {

	result := make([]string, 0)

	if n < 1 || n > len(s) {
		return result
	}

	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:n+i])
	}

	return result
}

func UnsafeFirst(n int, s string) string {

	if n < 1 || n > len(s) {
		return ""
	}

	return s[0:n]

}

func First(n int, s string) (first string, ok bool) {

	if n < 1 || n > len(s) {
		return "", false
	}

	return UnsafeFirst(n, s), true

}
