package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(v, initial)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {

	input := s.Reverse()
	return input.Foldl(fn, initial)

}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := make([]int, 0)
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func (s IntList) Length() int {
	result := 0
	for k := range s {
		result = k
	}

	return result
}

func (s IntList) Map(fn func(int) int) IntList {
	for k, v := range s {
		s[k] = fn(v)
	}
	return s
}

func (s IntList) Reverse() IntList {

	var result []int
	shift := s.Length() - 1
	for k, v := range s {
		result[shift-k] = v
	}

	return result

}

func (s IntList) Append(lst IntList) IntList {

	shift := len(s) - 1
	for k, v := range lst {
		s[shift+k] = v
	}

	return s

}

func (s IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		s = s.Append(v)
	}
	return s
}
