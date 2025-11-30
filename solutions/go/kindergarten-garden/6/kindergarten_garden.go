package kindergarten

import (
	"fmt"
	"slices"
	"strings"
)

type Garden map[string][]string

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//	diagram := `
//	VVCCGG
//	VVCCGG`
func NewGarden(diagram string, children []string) (*Garden, error) {

	flowers := map[string]string{
		"G": "grass",
		"C": "clover",
		"R": "radishes",
		"V": "violets",
	}

	garden := make(Garden)

	if len(diagram) == 0 || len(children) == 0 {
		return nil, fmt.Errorf("insufficient data")
	}

	plantRows := strings.Split(diagram, "\n")[1:]
	if len(plantRows) != 2 {
		return nil, fmt.Errorf("wrong diagram format")
	}

	if len(plantRows[0]) != len(plantRows[1]) ||
		len(plantRows[0])%2 != 0 {
		return nil, fmt.Errorf("rows could be of the same length containing even number of cups")
	}

	sortedKids := make([]string, len(children))
	copy(sortedKids, children)
	slices.Sort(sortedKids)

	for k, kidName := range sortedKids {

		if _, ok := garden[kidName]; ok {
			return nil, fmt.Errorf("this kid is already added")
		}

		for _, row := range plantRows {

			row = strings.TrimSpace(row)

			for i := range 2 {

				plant, exists := flowers[string(row[k*2+i])]
				if exists {
					garden[kidName] = append(garden[kidName], plant)
				} else {
					return nil, fmt.Errorf("invalid cup code")
				}

			}

		}

	}

	return &garden, nil

}

func (g *Garden) Plants(child string) ([]string, bool) {

	value, exists := (*g)[child]
	if !exists {
		return make([]string, 0), exists
	}

	return value, exists

}
