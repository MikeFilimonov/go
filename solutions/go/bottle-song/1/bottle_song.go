package bottlesong

import "fmt"

func Recite(startBottles, takeDown int) []string {

	result := make([]string, 0)
	constantPart := "And if one green bottle should accidentally fall,"

	addDelimiter := takeDown > 1

	for i := startBottles; i > startBottles-takeDown; i-- {

		hanging, willBe := parts(i)
		result = append(result, fmt.Sprintf("%s hanging on the wall,", hanging))
		result = append(result, fmt.Sprintf("%s hanging on the wall,", hanging))
		result = append(result, constantPart)
		result = append(result, fmt.Sprintf("There'll be %s hanging on the wall.", willBe))

		if addDelimiter && i > startBottles-takeDown+1 {
			result = append(result, "")
		}

	}

	return result

}

func parts(count int) (string, string) {

	switch count {
	case 10:
		return "Ten green bottles", "nine green bottles"
	case 9:
		return "Nine green bottles", "eight green bottles"
	case 8:
		return "Eight green bottles", "seven green bottles"
	case 7:
		return "Seven green bottles", "six green bottles"
	case 6:
		return "Six green bottles", "five green bottles"
	case 5:
		return "Five green bottles", "four green bottles"
	case 4:
		return "Four green bottles", "three green bottles"
	case 3:
		return "Three green bottles", "two green bottles"
	case 2:
		return "Two green bottles", "one green bottle"
	case 1:
		return "One green bottle", "no green bottles"
	default:
		return "tbd", "tbd"
	}

}
