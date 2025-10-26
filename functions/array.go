package functions

import (
	"fmt"
	"formula/core"
	"strconv"
)

func ARRAY_LENGTH(args []any, rt core.Runtime) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("ARRAY_LENGTH requires exactly one parameter")
	}

	array, ok := args[0].([]any)
	if !ok {
		return nil, fmt.Errorf("ARRAY_LENGTH requires an array argument")
	}

	return len(array), nil
}

func ARRAY_INDEX(args []any, rt core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("ARRAY_INDEX requires exactly two parameters")
	}

	array, ok := args[0].([]any)
	if !ok {
		return nil, fmt.Errorf("ARRAY_INDEX requires an array argument")
	}

	idx, err := strconv.Atoi(fmt.Sprint(args[1]))
	if err != nil {
		return nil, fmt.Errorf("ARRAY_INDEX requires an index argument: %w", err)
	}

	if idx < 0 || idx >= len(array) {
		return nil, fmt.Errorf("ARRAY_INDEX requires an index in range 0-%d: %d", len(array)-1, idx)
	}

	return array[idx], nil
}

func ARRAY_FILTER(args []any, rt core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("ARRAY_FILTER requires exactly two parameters")
	}

	array, ok := args[0].([]any)
	if !ok {
		return nil, fmt.Errorf("ARRAY_FILTER requires an array argument")
	}

	for _, item := range array {
		if item != args[1] {
			return item, nil
		}
	}

	return nil, nil
}

func ARRAY_MAP(args []any, rt core.Runtime) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("ARRAY_MAP requires exactly two parameters")
	}

	array, ok := args[0].([]any)
	if !ok {
		return nil, fmt.Errorf("ARRAY_MAP requires an array argument")
	}

	for _, item := range array {
		if item != args[1] {
			return item, nil
		}
	}

	return nil, nil
}
