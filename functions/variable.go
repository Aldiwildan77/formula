package functions

import (
	"fmt"
	"formula/core"
	"strconv"
	"strings"
)

func VAR(args []any, rt core.Runtime) (any, error) {
	const singleLevel = 1

	if len(args) < 1 {
		return nil, fmt.Errorf("VAR requires at least one parameter")
	}

	path := fmt.Sprint(args[0])
	paths := strings.Split(path, ".")

	root := paths[0]

	v, ok := rt.GetVar(root)
	if !ok {
		return nil, fmt.Errorf("variable %s not found", root)
	}

	if len(paths) == singleLevel {
		return v, nil
	}

	return traverse(v, paths[singleLevel:]), nil
}

func traverse(data any, paths []string) any {
	current := data
	for _, key := range paths {
		switch val := current.(type) {
		case map[string]any:
			current = val[key]
		case []any:
			idx, err := strconv.Atoi(key)
			if err != nil || idx < 0 || idx >= len(val) {
				return nil
			}
			current = val[idx]
		default:
			return nil
		}
	}
	return current
}
