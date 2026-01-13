package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {

	if s.Length() == 0 {
		return initial
	}

	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {

	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}

	return initial

}

func (s IntList) Filter(fn func(int) bool) IntList {

	result := make([]int, 0)
	if s.Length() == 0 {
		return result
	}
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
	if result > 0 {
		result++
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

	result := make(IntList, len(s))
	if s.Length() < 2 {
		return s
	}
	shift := s.Length() - 1
	for k, v := range s {
		result[shift-k] = v
	}

	return result

}

func (s IntList) Append(lst IntList) IntList {

	result := make([]int, 0, len(s)+len(lst))

	result = append(result, s...)
	result = append(result, lst...)

	return result

}

func (s IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		s = s.Append(v)
	}
	return s
}
