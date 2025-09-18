package bottlesong

import (
	"fmt"
	"strings"
)

var wordedNums = map[int]string{0: "No", 1: "One",
	2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 10: "Ten"}

func ending(i int) string {
	if i != 1 {
		return "s"
	}
	return ""
}

func Recite(startBottles, takeDown int) []string {

	result := make([]string, 0)
	constantPart := "And if one green bottle should accidentally fall,"

	stopBottles := startBottles - takeDown

	for i := startBottles; i > stopBottles; i-- {

		result = append(result, fmt.Sprintf("%s green bottle%s hanging on the wall,", wordedNums[i], ending(i)))
		result = append(result, fmt.Sprintf("%s green bottle%s hanging on the wall,", wordedNums[i], ending(i)))
		result = append(result, constantPart)
		result = append(result, fmt.Sprintf("There'll be %s green bottle%s hanging on the wall.",
			strings.ToLower(wordedNums[i-1]), ending(i-1)))

		if takeDown > 1 && i > stopBottles+1 {
			result = append(result, "")
		}

	}

	return result

}
