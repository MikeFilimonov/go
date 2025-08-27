// This package helps to convert a phrase to its acronym.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	regex := regexp.MustCompile("[^A-Za-z0-9- ]+")
	s = regex.ReplaceAllString(s, "")

	s = strings.ReplaceAll(s, "-", " ")

	var sb strings.Builder

	for _, v := range strings.Split(s, " ") {

		if len(v) > 0 {
			sb.WriteByte(v[0])
		}

	}

	return strings.ToUpper(sb.String())

}
