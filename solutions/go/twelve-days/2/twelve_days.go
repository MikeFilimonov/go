package twelve

import (
	"fmt"
	"strings"
)

var firstNonChristmasDay = 13

var days = [12]string{0: "first", 1: "second", 2: "third", 3: "fourth", 4: "fifth",
	5: "sixth", 6: "seventh", 7: "eighth", 8: "ninth", 9: "tenth", 10: "eleventh", 11: "twelfth"}

var gifts = [12]string{
	0:  "a Partridge in a Pear Tree",
	1:  "two Turtle Doves",
	2:  "three French Hens",
	3:  "four Calling Birds",
	4:  "five Gold Rings",
	5:  "six Geese-a-Laying",
	6:  "seven Swans-a-Swimming",
	7:  "eight Maids-a-Milking",
	8:  "nine Ladies Dancing",
	9:  "ten Lords-a-Leaping",
	10: "eleven Pipers Piping",
	11: "twelve Drummers Drumming"}

func Verse(i int) string {

	i--
	var sb strings.Builder
	splitter := ", "
	if i == 0 {
		splitter = ""
	}

	for counter := i; counter > 0; counter-- {
		sb.WriteString(gifts[counter])
		sb.WriteString(splitter)
	}

	lastGift := gifts[0]
	if i > 0 {
		lastGift = fmt.Sprintf("and %s", lastGift)
	}

	sb.WriteString(lastGift)

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.",
		days[i], sb.String())
}

func Song() string {

	var sb strings.Builder

	for i := 1; i < firstNonChristmasDay; i++ {
		sb.WriteString(Verse(i))
		if i < firstNonChristmasDay-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
