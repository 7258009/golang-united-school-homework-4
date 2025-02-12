package string_sum

import (
	"errors"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.NewReplacer(" ", "", "\t", "", "\n", "", "\r", "", "\v", "", "\f", "").Replace(input)

	if input == "" {
		return "", fmt.Errorf("error: %w", errorEmptyInput)
	}

	values := strings.Split(input, "+")

	if len(values) == 1 {
		lastIndex := strings.LastIndex(input, "-")

		switch {
		case lastIndex <= 0:
			return "", fmt.Errorf("wrong operands: %w", errorNotTwoOperands)
		case len(input)-1 == lastIndex:
			return "", fmt.Errorf("wrong operands: %w", errorNotTwoOperands)
		default:
			return calculate(input[:lastIndex], input[lastIndex:])
		}
	}

	if len(values) != 2 {
		return "", fmt.Errorf("wrong operands: %w", errorNotTwoOperands)
	}

	return calculate(values[0], values[1])
}

func calculate(value1, value2 string) (string, error) {
	number1, err1 := strconv.Atoi(value1)
	if err1 != nil {
		return "", fmt.Errorf("error in calculation: %w", err1)
	}

	number2, err2 := strconv.Atoi(value2)
	if err2 != nil {
		return "", fmt.Errorf("error in calculation: %w", err2)
	}

	return strconv.Itoa(number1 + number2), nil
}

