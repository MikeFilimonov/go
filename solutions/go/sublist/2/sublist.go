package sublist

import (
	"strconv"
	"strings"
)

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {

	lengthDiff := len(l1) - len(l2)
	var s1, s2 string
	var sb strings.Builder

	for _, v := range l1 {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString(";")
	}
	s1 = sb.String()

	sb.Reset()

	for _, v := range l2 {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString(";")
	}

	s2 = sb.String()

	if lengthDiff == 0 {
		if s1 == s2 {
			return RelationEqual
		}
	}

	if len(l1) == 0 {
		return RelationSublist
	}
	if len(l2) == 0 {
		return RelationSuperlist
	}

	if lengthDiff < 0 {
		if strings.Contains(s2, s1) {
			return RelationSublist
		}
	}

	if lengthDiff > 0 {
		if strings.Contains(s1, s2) {
			return RelationSuperlist
		}
	}

	return RelationUnequal
}
