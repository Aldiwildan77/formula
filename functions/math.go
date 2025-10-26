package functions

import (
	"fmt"
	"formula/core"
	"math"
	"strconv"
)

func SQRT(args []any, _ core.Runtime) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("SQRT requires exactly one parameter")
	}

	num, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("SQRT requires a number")
	}

	if num < 0 {
		return nil, fmt.Errorf("SQRT requires a non-negative number")
	}

	return math.Sqrt(num), nil
}

func ABS(args []any, _ core.Runtime) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("ABS requires exactly one parameter")
	}

	num, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("ABS requires a number")
	}

	return math.Abs(num), nil
}

func FLOOR(args []any, _ core.Runtime) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("FLOOR requires exactly one parameter")
	}

	num, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("FLOOR requires a number")
	}

	return math.Floor(num), nil
}

func CEIL(args []any, _ core.Runtime) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("CEIL requires exactly one parameter")
	}

	num, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("CEIL requires a number")
	}

	return math.Ceil(num), nil
}

func ROUND(args []any, _ core.Runtime) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("ROUND requires exactly one parameter")
	}

	num, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("ROUND requires a number")
	}

	return math.Round(num), nil
}

func MIN(args []any, _ core.Runtime) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("MIN requires at least one parameter")
	}

	min := math.Inf(1)
	for _, arg := range args {
		num, err := strconv.ParseFloat(fmt.Sprint(arg), 64)
		if err != nil {
			return nil, fmt.Errorf("MIN requires numbers")
		}
		if num < min {
			min = num
		}
	}

	return min, nil
}

func MAX(args []any, _ core.Runtime) (any, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("MAX requires at least one parameter")
	}

	max := math.Inf(-1)
	for _, arg := range args {
		num, err := strconv.ParseFloat(fmt.Sprint(arg), 64)
		if err != nil {
			return nil, fmt.Errorf("MAX requires numbers")
		}
		if num > max {
			max = num
		}
	}

	return max, nil
}
