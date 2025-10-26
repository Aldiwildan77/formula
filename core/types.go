package core

type VarRuntime map[string]any

type FuncRuntime func(args []any, r Runtime) (any, error)

type FuncsRuntime map[string]FuncRuntime

type Runtime interface {
	GetVar(name string) (any, bool)
	SetVar(name string, val any)
	HasVar(name string) bool
	Register(name string, fn FuncRuntime)
	Eval(e Expr) (any, error)
}
