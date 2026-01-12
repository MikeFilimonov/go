package cryptosquare

import (
	"regexp"
	"strings"
)

func Encode(pt string) string {

	pt = normalize(pt)
	rows, columns := factors(len(pt))

	var sb strings.Builder
	var buffer []string
	for i := range columns {

		for j := range rows {

			position := i + j*columns

			if position < len(pt) {
				if len(pt[position:]) > 1 {
					sb.WriteString(pt[position : position+1])
				} else {
					sb.WriteString(pt[position:])
				}
			}

		}

		if len(sb.String()) < rows {
			sb.WriteString(" ")
		}

		buffer = append(buffer, sb.String())
		sb.Reset()

	}

	result := strings.Join(buffer, " ")
	return result

}

func normalize(input string) string {

	regex := regexp.MustCompile(`[^a-zA-Z0-9]`)
	noOdds := regex.ReplaceAllString(input, "")
	return strings.ToLower(noOdds)

}

func factors(product int) (int, int) {

	if product < 1 {
		return 0, 0
	}

	value := 1
	for value*(value+1) < product {
		value++
	}

	if value*value >= product {
		return value, value
	}

	return value, value + 1
}
