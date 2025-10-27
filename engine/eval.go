package engine

import (
	"encoding/json"
	"fmt"
	"formula/core"
	"formula/functions"
)

type runtime struct {
	Vars     core.VarRuntime
	Funcs    core.FuncsRuntime
	debugger *Debugger
}

func WithDebugger(debugger *Debugger) func(*runtime) {
	return func(r *runtime) {
		r.debugger = debugger
	}
}

func WithVars(vars core.VarRuntime) func(*runtime) {
	return func(r *runtime) {
		r.Vars = vars
	}
}

func WithFuncs(funcs core.FuncsRuntime) func(*runtime) {
	return func(r *runtime) {
		r.Funcs = funcs
	}
}

func NewRuntime(opts ...func(*runtime)) core.Runtime {
	// Default
	r := &runtime{
		Vars:  make(core.VarRuntime),
		Funcs: make(core.FuncsRuntime),
	}

	for _, opt := range opts {
		opt(r)
	}

	r.registerBuiltins()
	return r
}

func (r *runtime) registerBuiltins() {
	// Arithmetic Operators
	r.Funcs["ADD"] = functions.ADD
	r.Funcs["SUBTRACT"] = functions.SUBTRACT
	r.Funcs["MULTIPLY"] = functions.MULTIPLY
	r.Funcs["DIVIDE"] = functions.DIVIDE
	r.Funcs["MOD"] = functions.MOD
	r.Funcs["POWER"] = functions.POWER
	r.Funcs["POW"] = functions.POWER
	r.Funcs["NEGATE"] = functions.NEGATE

	// Comparison Operators
	r.Funcs["EQ"] = functions.EQ
	r.Funcs["EQUAL"] = functions.EQ
	r.Funcs["NEQ"] = functions.NEQ
	r.Funcs["GT"] = functions.GT
	r.Funcs["GTE"] = functions.GTE
	r.Funcs["LT"] = functions.LT
	r.Funcs["LTE"] = functions.LTE

	// Logical Operators
	r.Funcs["IF"] = functions.IF
	r.Funcs["AND"] = functions.AND
	r.Funcs["OR"] = functions.OR
	r.Funcs["NOT"] = functions.NOT

	// Mathematical Functions
	r.Funcs["SQRT"] = functions.SQRT
	r.Funcs["ABS"] = functions.ABS
	r.Funcs["FLOOR"] = functions.FLOOR
	r.Funcs["CEIL"] = functions.CEIL
	r.Funcs["ROUND"] = functions.ROUND
	r.Funcs["MIN"] = functions.MIN
	r.Funcs["MAX"] = functions.MAX

	// Variable
	r.Funcs["VAR"] = functions.VAR

	// Array Functions
	r.Funcs["ARRAY_LENGTH"] = functions.ARRAY_LENGTH
	r.Funcs["ARRAY_INDEX"] = functions.ARRAY_INDEX
	r.Funcs["ARRAY_FILTER"] = functions.ARRAY_FILTER
	r.Funcs["ARRAY_MAP"] = functions.ARRAY_MAP
}

func (r *runtime) Eval(e core.Expr) (any, error) {
	switch n := e.(type) {
	case *core.Literal:
		return n.Value, nil
	case *core.Func:
		fnName := n.Name
		fn, ok := r.Funcs[fnName]
		if !ok {
			return nil, fmt.Errorf("function %s not found", fnName)
		}

		// Evaluate arguments first
		var args []any
		for _, a := range n.Args {
			v, err := r.Eval(a)
			if err != nil {
				return nil, err
			}
			args = append(args, v)
		}

		// Track function entry if debugger is enabled
		if r.debugger != nil {
			r.debugger.EnterFunc(fnName, args...)
		}

		// Execute the function
		result, err := fn(args, r)

		// Track function exit if debugger is enabled
		if r.debugger != nil {
			r.debugger.ExitFunc(fnName, result, err)
		}

		return result, err
	default:
		return nil, fmt.Errorf("invalid expr")
	}
}

func (r *runtime) Register(name string, fn core.FuncRuntime) {
	r.Funcs[name] = fn
}

func (r *runtime) GetVar(name string) (any, bool) {
	v, ok := r.Vars[name]
	return v, ok
}

func (r *runtime) HasVar(name string) bool {
	_, ok := r.Vars[name]
	return ok
}

func (r *runtime) SetVar(name string, val any) {
	switch data := val.(type) {
	case string:
		var jsonVal any
		if err := json.Unmarshal([]byte(data), &jsonVal); err == nil {
			r.Vars[name] = jsonVal
			return
		}
	default:
		r.Vars[name] = val
	}
}
