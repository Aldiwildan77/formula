package functions

import (
	"fmt"
	"formula/core"
	"strconv"
)

func EQ(args []any, _ core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("EQ must have two parameters")
	}

	return fmt.Sprint(args[0]) == fmt.Sprint(args[1]), nil
}

func NEQ(args []any, _ core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("NEQ must have two parameters")
	}

	return fmt.Sprint(args[0]) != fmt.Sprint(args[1]), nil
}

func GT(args []any, _ core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("GT must have two parameters")
	}

	num1, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("GT requires numbers")
	}

	num2, err := strconv.ParseFloat(fmt.Sprint(args[1]), 64)
	if err != nil {
		return nil, fmt.Errorf("GT requires numbers")
	}

	return num1 > num2, nil
}

func GTE(args []any, _ core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("GTE must have two parameters")
	}

	num1, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("GTE requires numbers")
	}

	num2, err := strconv.ParseFloat(fmt.Sprint(args[1]), 64)
	if err != nil {
		return nil, fmt.Errorf("GTE requires numbers")
	}

	return num1 >= num2, nil
}

func LT(args []any, _ core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("LT must have two parameters")
	}

	num1, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("LT requires numbers")
	}

	num2, err := strconv.ParseFloat(fmt.Sprint(args[1]), 64)
	if err != nil {
		return nil, fmt.Errorf("LT requires numbers")
	}

	return num1 < num2, nil
}

func LTE(args []any, _ core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("LTE must have two parameters")
	}

	num1, err := strconv.ParseFloat(fmt.Sprint(args[0]), 64)
	if err != nil {
		return nil, fmt.Errorf("LTE requires numbers")
	}

	num2, err := strconv.ParseFloat(fmt.Sprint(args[1]), 64)
	if err != nil {
		return nil, fmt.Errorf("LTE requires numbers")
	}

	return num1 <= num2, nil
}
