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
		set[v] = true
	}

	return set
}

func (s Set) String() string {

	if s.IsEmpty() {
		return "{}"
	}

	values := make([]string, 0)

	for k := range s {
		values = append(values, fmt.Sprintf("%q", k))
	}

	description := strings.Join(values, ",")

	return fmt.Sprintf("{%s}", description)

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

	if s1.IsEmpty() || s2.IsEmpty() {
		return true
	}

	for k := range s1 {
		if s2[k] {
			return true
		}
	}
	return false
}

func Disjoint(s1, s2 Set) bool {

	if s1.IsEmpty() || s2.IsEmpty() {
		return true
	}

	return !Subset(s1, s2)
}

func Equal(s1, s2 Set) bool {

	set := s1

	if len(s1) != len(s2) {
		return false
	}
	// if len(s2.value) > len(s1.value) {
	// 	set = s2
	// }

	for k := range set {
		if !s2[k] {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {

	result := New()
	values := make(map[string]bool)
	for k := range s1 {
		if s2[k] {
			values[k] = true
		}
	}
	return result

}

func Difference(s1, s2 Set) Set {

	result := New()
	intersection := Intersection(s1, s2)
	for k := range intersection {
		if !s2[k] {
			result[k] = true
		}
	}

	return result
}

func Union(s1, s2 Set) Set {
	result := New()

	for k := range s1 {
		result[k] = true
	}

	for k := range s2 {
		result[k] = true
	}

	return result
}
