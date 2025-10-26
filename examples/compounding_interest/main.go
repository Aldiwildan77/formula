package main

import (
	"fmt"
	"formula/engine"
)

func main() {
	input := `
		MULTIPLY(
			VAR('P'), 
			POW(
				ADD(
					1, 
					DIVIDE(
						VAR('r'), 
						VAR('n')
					)
				), 
				MULTIPLY(
					VAR('n'), 
					VAR('t')
				)
			)
		)
	`

	debugger := &engine.Debugger{IsEnabled: true}

	parser := engine.NewParser(input, debugger)
	ast := parser.ParseExpr()

	debugger.DumpAST(ast)

	rt := engine.NewRuntime()
	rt.SetVar("P", 1000)
	rt.SetVar("r", 0.05)
	rt.SetVar("n", 12)
	rt.SetVar("t", 10)

	val, err := rt.Eval(ast)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result = %#v\n", val)
}
