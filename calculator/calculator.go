package calculator

import (
	"calculator/stack"
	"errors"
	"strconv"
)

func isOperator(operator string) bool {
	if operator == "*" || operator == "+" || operator == "/" || operator == "-" {
		return true
	}
	return false
}

func isNumber(num string) bool {
	if _, err := strconv.Atoi(string(num)); err == nil {
		return true
	} else if _, err := strconv.ParseFloat(num, 64); err == nil {
		return true
	}
	return false
}

func priorityOperation(operator string) int {
	if operator == "*" || operator == "/" {
		return 2
	}
	if operator == "+" || operator == "-" {
		return 1
	}
	return 0
}

func calculation(arg1 float64, arg2 float64, operand string) (float64, error) {
	var result float64
	switch operand {
	case "+":
		result = arg1 + arg2
	case "-":
		result = arg1 - arg2
	case "*":
		result = arg1 * arg2
	case "/":
		if arg2 == 0 {
			return 0, errors.New("error: division on zero")
		}
		result = arg1 / arg2
	}
	return result, nil
}

func parseString(line string) []string {
	operand := stack.New()
	result := []string{}

	numBuff := ""
	for _, val := range line {
		if isNumber(string(val)) || string(val) == "." {
			numBuff += string(val)
		} else {
			if numBuff != "" {
				result = append(result, numBuff)
				numBuff = ""
			}

			if isOperator(string(val)) {
				for (priorityOperation(string(val)) <= priorityOperation(operand.Peek())) && operand.Len() != 0 {
					result = append(result, operand.Pop())
				}
				operand.Push(string(val))
			}

			if string(val) == "(" {
				operand.Push(string(val))
			}

			if string(val) == ")" {
				operator := operand.Pop()
				for operator != "(" {
					result = append(result, operator)
					operator = operand.Pop()
				}
			}
		}
	}

	if numBuff != "" {
		result = append(result, numBuff)
	}

	for operand.Len() != 0 {
		result = append(result, operand.Pop())
	}
	return result
}

func Calculate(line string) (float64, error) {
	expression := parseString(line)
	numberStack := stack.New()
	for _, val := range expression {
		if isNumber(val) {
			numberStack.Push(val)
		} else if isOperator(val) {
			if numberStack.Len() < 2 {
				return 0, errors.New("error: There are not enough numbers for operands")
			}

			num2, _ := strconv.ParseFloat(numberStack.Pop(), 64)
			num1, _ := strconv.ParseFloat(numberStack.Pop(), 64)

			var res float64
			var err error
			if res, err = calculation(num1, num2, val); err != nil {
				return 0, err
			}
			numberStack.Push(strconv.FormatFloat(res, 'f', 6, 64))
		}
	}

	if numberStack.Len() != 1 {
		return 0, errors.New("error: The value is not calculated correctly")
	}

	return strconv.ParseFloat(numberStack.Pop(), 64)
}
