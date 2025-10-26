package main

import (
	"fmt"
	"formula/engine"
)

func main() {
	input := `
		VAR('person.address.city', 'Malang')
	`

	debugger := &engine.Debugger{IsEnabled: true}

	parser := engine.NewParser(input, debugger)
	ast := parser.ParseExpr()

	debugger.DumpAST(ast)

	rt := engine.NewRuntime()
	rt.SetVar("person", `{
		"name": "Aldiwildan",
		"age": 26,
		"address": {"city": "Tokyo", "zip": "10002"},
		"hobbies": ["coding", "reading"],
		"contacts": [
			{"type": "email", "value": "Aldiwildan@gmail.com"},
			{"type": "phone", "value": "081234567890"}
		]
	}`)

	val, err := rt.Eval(ast)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result = %#v\n", val)
}
