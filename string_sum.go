package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
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
	operands, err := parse(input)
	if err != nil {
		return "", fmt.Errorf("error parsing operands: %v", err)
	}
	return strconv.Itoa(operands[0] + operands[1]), nil
}

func parse(input string) (result []int, err error) {
	s := strings.TrimFunc(input, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	if len(s) == 0 {
		return nil, errorEmptyInput
	}
	var curr string
	for i, r := range s {
		if len(curr) == 0 {
			curr += string(r)
		} else {
			if unicode.IsDigit(r) {
				curr += string(r)
			}
			if string(r) == "-" || string(r) == "+" {
				currInt, err := strconv.Atoi(curr)
				if err != nil {
					return []int{}, err
				}
				result = append(result, currInt)
				curr = string(r)
			}
			if i == len(s)-1 {
				currInt, err := strconv.Atoi(curr)
				if err != nil {
					return []int{}, errorNotTwoOperands
				}
				result = append(result, currInt)
			}
		}
	}
	if len(result) != 2 {
		return []int{}, errorNotTwoOperands
	}
	return result, nil
}
