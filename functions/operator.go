package functions

import (
	"fmt"
	"formula/core"
	"math"
	"strconv"
)

const NumberBitSize = 64

func ADD(args []any, _ core.Runtime) (any, error) {
	result := math.NaN()
	for _, arg := range args {
		num, err := strconv.ParseFloat(fmt.Sprint(arg), NumberBitSize)
		if err != nil {
			return nil, fmt.Errorf("ADD requires a number")
		}

		if math.IsNaN(result) {
			result = num
			continue
		}

		result += num
	}

	if math.IsNaN(result) {
		return nil, fmt.Errorf("ADD requires an arguments")
	}

	return result, nil
}

func SUBTRACT(args []any, _ core.Runtime) (any, error) {
	result := math.NaN()
	for _, arg := range args {
		num, err := strconv.ParseFloat(fmt.Sprint(arg), NumberBitSize)
		if err != nil {
			return nil, fmt.Errorf("SUBTRACT requires a number")
		}

		if math.IsNaN(result) {
			result = num
			continue
		}

		result -= num
	}

	if math.IsNaN(result) {
		return nil, fmt.Errorf("SUBTRACT requires an arguments")
	}

	return result, nil
}

func MULTIPLY(args []any, _ core.Runtime) (any, error) {
	result := math.NaN()
	for _, arg := range args {
		num, err := strconv.ParseFloat(fmt.Sprint(arg), NumberBitSize)
		if err != nil {
			return nil, fmt.Errorf("MULTIPLY requires a number")
		}

		if math.IsNaN(result) {
			result = num
			continue
		}

		result *= num
	}

	if math.IsNaN(result) {
		return nil, fmt.Errorf("MULTIPLY requires an arguments")
	}

	return result, nil
}

func DIVIDE(args []any, _ core.Runtime) (any, error) {
	result := math.NaN()
	for _, arg := range args {
		num, err := strconv.ParseFloat(fmt.Sprint(arg), NumberBitSize)
		if err != nil {
			return nil, fmt.Errorf("DIVIDE requires a number")
		}

		if math.IsNaN(result) {
			result = num
			continue
		}

		if num == 0 {
			return nil, fmt.Errorf("DIVIDE by zero is not allowed")
		}

		result /= num
	}

	if math.IsNaN(result) {
		return nil, fmt.Errorf("DIVIDE requires an arguments")
	}

	return result, nil
}

func MOD(args []any, _ core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("MOD requires exactly two parameters")
	}

	num1, err := strconv.ParseFloat(fmt.Sprint(args[0]), NumberBitSize)
	if err != nil {
		return nil, fmt.Errorf("MOD requires a number")
	}

	num2, err := strconv.ParseFloat(fmt.Sprint(args[1]), NumberBitSize)
	if err != nil {
		return nil, fmt.Errorf("MOD requires a number")
	}

	if num2 == 0 {
		return nil, fmt.Errorf("MOD by zero is not allowed")
	}

	return math.Mod(num1, num2), nil
}

func POWER(args []any, _ core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("POWER requires exactly two parameters: base and exponent")
	}

	base, err := strconv.ParseFloat(fmt.Sprint(args[0]), NumberBitSize)
	if err != nil {
		return nil, fmt.Errorf("POWER requires a number")
	}

	exponent, err := strconv.ParseFloat(fmt.Sprint(args[1]), NumberBitSize)
	if err != nil {
		return nil, fmt.Errorf("POWER requires a number")
	}

	return math.Pow(base, exponent), nil
}

func NEGATE(args []any, _ core.Runtime) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("NEGATE requires exactly one parameter")
	}

	num, err := strconv.ParseFloat(fmt.Sprint(args[0]), NumberBitSize)
	if err != nil {
		return nil, fmt.Errorf("NEGATE requires a number")
	}

	return -num, nil
}
