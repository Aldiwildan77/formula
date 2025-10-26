package core

type Expr any

type Exprs []Expr

type Func struct {
	Name string
	Args []Expr
}

type Literal struct {
	Value any
}

type Variable struct {
	Name string
}
