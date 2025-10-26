package main

import (
	"fmt"
	"formula/engine"
)

func main() {
	input := "IF(EQ(ADD(2,5),VAR('total')), TRUE, FALSE)"

	debugger := &engine.Debugger{IsEnabled: true}

	parser := engine.NewParser(input, debugger)
	ast := parser.ParseExpr()

	rt := engine.NewRuntime()
	rt.SetVar("total", "7")

	val, err := rt.Eval(ast)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result = %#v\n", val)

	// debugger.Print()
	debugger.DumpAST(ast)
}
