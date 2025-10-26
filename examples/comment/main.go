package main

import (
	"fmt"
	"formula/engine"
)

func main() {
	input := `
		/* 
			This is a multi-line comment 
			VAR('item_price') is the price of the item
			VAR('quantity') is the quantity of the item
		*/
		MULTIPLY(
			VAR('item_price'), 
			# This is a comment VAR('quantity')
			VAR('quantity') -- This is a comment
		)
	`

	debugger := &engine.Debugger{IsEnabled: true}

	parser := engine.NewParser(input, debugger)
	ast := parser.ParseExpr()

	debugger.DumpAST(ast)

	rt := engine.NewRuntime()
	rt.SetVar("item_price", 1000)
	rt.SetVar("quantity", 3)

	val, err := rt.Eval(ast)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result = %#v\n", val)
}
