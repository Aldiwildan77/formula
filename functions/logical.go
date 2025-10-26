package functions

import (
	"fmt"
	"formula/core"
)

func IF(args []any, _ core.Runtime) (any, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("IF requires three parameters: condition, then-value, else-value")
	}

	cond, ok := args[0].(bool)
	if !ok {
		return nil, fmt.Errorf("IF requires first argument to be a boolean condition")
	}

	if cond {
		return args[1], nil
	}
	return args[2], nil
}

func AND(args []any, _ core.Runtime) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("AND requires at least two parameters")
	}

	for _, arg := range args {
		val, ok := arg.(bool)
		if !ok {
			return nil, fmt.Errorf("AND requires boolean arguments")
		}
		if !val {
			return false, nil
		}
	}

	return true, nil
}

func OR(args []any, _ core.Runtime) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("OR requires at least two parameters")
	}

	for _, arg := range args {
		val, ok := arg.(bool)
		if !ok {
			return nil, fmt.Errorf("OR requires boolean arguments")
		}
		if val {
			return true, nil
		}
	}

	return false, nil
}

func NOT(args []any, _ core.Runtime) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("NOT requires exactly one parameter")
	}

	val, ok := args[0].(bool)
	if !ok {
		return nil, fmt.Errorf("NOT requires a boolean argument")
	}

	return !val, nil
}
