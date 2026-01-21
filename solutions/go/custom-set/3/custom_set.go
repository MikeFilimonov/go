package stringset

import (
	"fmt"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

type Set map[string]bool

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {

	set := New()

	for _, v := range l {
		set.Add(v)
	}

	return set
}

func (s Set) String() string {

	if s.IsEmpty() {
		return "{}"
	}

	var sb strings.Builder

	for k := range s {

		sb.WriteString(fmt.Sprintf("%+q, ", k))

	}

	description := sb.String()

	return fmt.Sprintf("{%s}", description[:len(description)-2])

}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	return s[elem]
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {

	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {

	for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {

	set := s1

	if len(s1) != len(s2) {
		return false
	}

	for k := range set {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {

	result := New()

	for k := range s1 {
		if s2.Has(k) {
			result.Add(k)
		}
	}
	return result

}

func Difference(s1, s2 Set) Set {

	result := New()
	for k := range s1 {
		if !s2.Has(k) {
			result.Add(k)
		}
	}

	return result
}

func Union(s1, s2 Set) Set {
	result := New()

	for k := range s1 {
		result.Add(k)
	}

	for k := range s2 {
		result.Add(k)
	}

	return result
}
