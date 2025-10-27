package main

import (
	"fmt"
	"formula/core"
	"formula/engine"
	"strconv"
)

func main() {
	input := `
		/* 
			This FACTORIAL is a recursive function built with built in function.
		*/
		FACTORIAL(5)
	`

	debugger := &engine.Debugger{IsEnabled: true}

	parser := engine.NewParser(input, debugger)
	ast := parser.ParseExpr()

	rt := engine.NewRuntime(engine.WithDebugger(debugger))
	rt.Register("FACTORIAL", func(args []any, r core.Runtime) (any, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("FACTORIAL requires 1 parameter")
		}

		n, _ := strconv.Atoi(fmt.Sprint(args[0]))

		if n <= 1 {
			return 1, nil
		}

		// this code below works, but inefficient
		// ps := engine.NewParser(fmt.Sprintf("MULTIPLY(%d, FACTORIAL(%d))", n, n-1), nil)
		// nAst := ps.ParseExpr()
		// return r.Eval(nAst)

		// this code below is the efficient approach
		nAst := &core.Func{
			Name: "MULTIPLY",
			Args: core.Exprs{
				&core.Literal{Value: float64(n)},
				&core.Func{
					Name: "FACTORIAL",
					Args: core.Exprs{
						&core.Literal{Value: float64(n - 1)},
					},
				},
			},
		}

		return r.Eval(nAst)
	})

	val, err := rt.Eval(ast)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result = %#v\n", val)
}
