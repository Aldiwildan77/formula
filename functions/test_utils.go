package functions

import "formula/core"

// MockRuntime provides a mock implementation of core.Runtime for testing
type MockRuntime struct{}

func (m *MockRuntime) Eval(e core.Expr) (any, error) {
	return nil, nil
}

func (m *MockRuntime) Register(name string, fn core.FuncRuntime) {
}

func (m *MockRuntime) GetVar(name string) (any, bool) {
	return nil, false
}

func (m *MockRuntime) HasVar(name string) bool {
	return false
}

func (m *MockRuntime) SetVar(name string, val any) {
}

// VarMockRuntime is a specialized mock runtime for variable tests
type VarMockRuntime struct {
	vars map[string]any
}

func NewVarMockRuntime() *VarMockRuntime {
	return &VarMockRuntime{
		vars: make(map[string]any),
	}
}

func (m *VarMockRuntime) Eval(e core.Expr) (any, error) {
	return nil, nil
}

func (m *VarMockRuntime) Register(name string, fn core.FuncRuntime) {
}

func (m *VarMockRuntime) GetVar(name string) (any, bool) {
	val, ok := m.vars[name]
	return val, ok
}

func (m *VarMockRuntime) SetVar(name string, val any) {
	m.vars[name] = val
}

func (m *VarMockRuntime) HasVar(name string) bool {
	return false
}
