package wordy

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Problem struct {
	Operands   []int
	Operators []rune
}

func Answer(question string) (int, bool) {

	problem, err := parseProblem(question)
	if err != nil && !isValidProblem(problem) {
		return 0, false
	}

	return Compute(problem), true

}

func parseProblem(description string) (*Problem, error) {

	question := "what is"
	description = strings.ToLower(description)

	if !strings.ContainsAny(description, question) {
		return &Problem{}, fmt.Errorf("problem should contain \"what is\" literal")
	}

	for _, v := range []string{"by", question, "?"} {
		description = strings.ReplaceAll(description, v, "")
	}
	description = strings.TrimSpace(description)

	problem := Problem{}

	elements := strings.Split(description, " ")

	awaitsForOperand := true

	for _, v := range elements {

		result, err := strconv.Atoi(v)
		if err == nil {
			problem.Operands = append(problem.Operands, result)
			awaitsForOperand = false

		} else {

			operation, err := commandToOperation(v)
			if err != nil || awaitsForOperand {
				return nil, err
			}
			problem.Operators = append(problem.Operators, operation)
			awaitsForOperand = true
		}
	}

	return &problem, nil

}

func isValidProblem(problem *Problem) bool {
	return len(problem.Operands)-1 == len(problem.Operators)
}

func Compute(input *Problem) int {

	if len(input.Operands) == 1 {
		return input.Operands[0]
	}

	firstOperatorPosition := 0

	// if the first operator is '+' or '-' and the combo includes '*' or '/'
	if slices.Index(input.Operators, '*') > 0 {
		firstOperatorPosition = slices.Index(input.Operators, '*')
	}
	if firstOperatorPosition == 0 && slices.Index(input.Operators, '/') > 0{
		firstOperatorPosition = slices.Index(input.Operators, '/')		
	}

	if requiresReordering{
		firstOperatorPosition = 
	}

	for _, v := range input.Operators {

		if requiresReordering {

		}

	}

	return 0
}

func commandToOperation(command string) (rune, error) {

	switch command {
	case "plus":
		return '+', nil

	case "minus":
		return '-', nil
	case "multiplied":
		return '*', nil
	case "divided":
		return '/', nil
	default:
		return 0, fmt.Errorf("failed to map the operation %s", command)

	}

}
