package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

type Problem struct {
	Operands  []int
	Operators []string
}

func Answer(question string) (int, bool) {

	problem, err := parseProblem(question)

	if err != nil || !isValidProblem(problem) {
		return 0, false
	}

	result := Compute(problem)
	return result, true

}

func parseProblem(description string) (*Problem, error) {

	question := "what is"
	description = strings.ToLower(description)

	if !strings.Contains(description, question) {
		return nil, fmt.Errorf("problem should contain \"what is\" literal")
	}

	for _, v := range []string{"by", question, "?"} {
		description = strings.ReplaceAll(description, v, "")
	}
	description = strings.TrimSpace(description)
	fmt.Printf("%v\n", description)

	problem := Problem{}

	elements := strings.Fields(description)

	if len(elements) == 0 {
		return nil, fmt.Errorf("the first element should be a number")
	}

	_, err := strconv.Atoi(elements[0])
	if err != nil {
		return nil, fmt.Errorf("the first element should be a number")
	}

	awaitsForOperand := true

	for _, v := range elements {

		result, err := strconv.Atoi(v)

		if err == nil {

			if !awaitsForOperand {
				return nil, fmt.Errorf("reject two numbers in a row")
			}

			problem.Operands = append(problem.Operands, result)
			awaitsForOperand = false

		} else {

			if awaitsForOperand {
				return nil, fmt.Errorf("expected for an operator, got an operand")
			}
			problem.Operators = append(problem.Operators, v)
			awaitsForOperand = true
		}
	}

	return &problem, nil

}

func isValidProblem(problem *Problem) bool {
	return len(problem.Operands) > 0 && len(problem.Operands)-1 == len(problem.Operators)
}

func Compute(input *Problem) int {

	result := input.Operands[0]

	if len(input.Operands) == 1 {
		return result
	}

	for k, operation := range input.Operators {

		subresult, err := commandToOperation(result, operation, input.Operands[k+1])
		if err != nil {
			return 0
		}
		result = subresult

	}

	return result
}

func commandToOperation(a int, command string, b int) (int, error) {

	switch command {
	case "plus":
		return a + b, nil
	case "minus":
		return a - b, nil
	case "multiplied":
		return a * b, nil
	case "divided":

		if b == 0 {
			return 0, fmt.Errorf("divided by zero")
		}
		return a / b, nil

	default:
		return 0, fmt.Errorf("failed to map the operation %s", command)

	}

}
