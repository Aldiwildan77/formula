package engine

import (
	"fmt"
	"formula/core"
	"time"
)

type PhaseType string

const (
	PhaseLexer  PhaseType = "lexer"
	PhaseParser PhaseType = "parser"
	PhaseEval   PhaseType = "eval"
)

type Debug struct {
	Phase     PhaseType
	Depth     int
	Message   string
	Timestamp time.Time
}

type Debugger struct {
	IsEnabled bool
	Histories []Debug
	evalDepth int
}

func (d *Debugger) Log(phase PhaseType, depth int, message string) {
	if !d.IsEnabled {
		return
	}

	if d.Histories == nil {
		d.Histories = make([]Debug, 0)
	}

	d.Histories = append(d.Histories, Debug{
		Phase:     phase,
		Depth:     depth,
		Message:   message,
		Timestamp: time.Now(),
	})
}

func (d *Debugger) Print() {
	for _, e := range d.Histories {
		fmt.Printf("[%s]\t[%d]\t%s\n", e.Phase, e.Depth, e.Message)
	}
}

func (d *Debugger) EnterFunc(name string, args ...any) {
	if !d.IsEnabled {
		return
	}

	// Build the prefix - continue vertical lines from parent contexts
	prefix := ""
	for i := 0; i < d.evalDepth; i++ {
		prefix += "│   "
	}

	// Build argument string
	argStr := ""
	if len(args) > 0 {
		argStrs := make([]string, len(args))
		for i, arg := range args {
			argStrs[i] = fmt.Sprintf("%v", arg)
		}
		argStr = fmt.Sprintf("(%s)", fmt.Sprintf("%v", args[0]))
		if len(args) > 1 {
			for i := 1; i < len(args); i++ {
				argStr = fmt.Sprintf("%s, %v", argStr, args[i])
			}
		}
	} else {
		argStr = "()"
	}

	// Use tree branch for function entry
	branch := "├── "
	if prefix == "" {
		branch = ""
	}

	fmt.Printf("%s%s%s%s\n", prefix, branch, name, argStr)
	d.Log(PhaseEval, d.evalDepth, fmt.Sprintf("→ %s%s", name, argStr))
	d.evalDepth++
}

func (d *Debugger) ExitFunc(name string, result any, err error) {
	if !d.IsEnabled {
		return
	}

	// Build the prefix at current depth (before decrementing)
	// This shows the result is inside the function context
	prefix := ""
	for i := 0; i < d.evalDepth; i++ {
		prefix += "│   "
	}

	if err != nil {
		fmt.Printf("%s%s = ERROR: %v\n", prefix, name, err)
		d.Log(PhaseEval, d.evalDepth, fmt.Sprintf("← %s error: %v", name, err))
	} else {
		fmt.Printf("%s%s = %v\n", prefix, name, result)
		d.Log(PhaseEval, d.evalDepth, fmt.Sprintf("← %s = %v", name, result))
	}

	d.evalDepth--
}

func (d *Debugger) DumpAST(expr core.Expr) {
	if !d.IsEnabled {
		return
	}

	fmt.Println("\n=== AST Visualization ===")
	PrintAST(expr, "", true)
	fmt.Println("=========================")
}

func PrintAST(expr core.Expr, prefix string, isLast bool) {
	var branch string
	if prefix == "" {
		branch = "" // Root node (no connector)
	} else if isLast {
		branch = "└── "
	} else {
		branch = "├── "
	}

	switch e := expr.(type) {
	case *core.Func, core.Func:
		var fn core.Func
		switch v := e.(type) {
		case *core.Func:
			fn = *v
		case core.Func:
			fn = v
		}

		fmt.Printf("%s%sFuncCall(%s)\n", prefix, branch, fn.Name)

		for i, arg := range fn.Args {
			isLastArg := i == len(fn.Args)-1

			// Update prefix for next level
			nextPrefix := prefix
			if prefix != "" {
				if isLast {
					nextPrefix += "    "
				} else {
					nextPrefix += "│   "
				}
			} else {
				if isLast {
					nextPrefix = "    "
				} else {
					nextPrefix = "│   "
				}
			}

			PrintAST(arg, nextPrefix, isLastArg)
		}

	case *core.Literal, core.Literal:
		var lit core.Literal
		switch v := e.(type) {
		case *core.Literal:
			lit = *v
		case core.Literal:
			lit = v
		}
		fmt.Printf("%s%sLiteral(%v)\n", prefix, branch, lit.Value)

	case *core.Variable, core.Variable:
		var v core.Variable
		switch x := e.(type) {
		case *core.Variable:
			v = *x
		case core.Variable:
			v = x
		}
		fmt.Printf("%s%sVariable(%s)\n", prefix, branch, v.Name)

	default:
		fmt.Printf("%s%s<unknown: %T>\n", prefix, branch, e)
	}
}
