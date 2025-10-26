package functions

import (
	"testing"
)

func TestVAR_ExistingVariable(t *testing.T) {
	// GIVEN
	rt := NewVarMockRuntime()
	rt.SetVar("total", 100)
	args := []any{"total"}

	// WHEN
	result, err := VAR(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 100 {
		t.Errorf("Expected 100, got %v", result)
	}
}

func TestVAR_UndefinedVariable(t *testing.T) {
	// GIVEN
	rt := NewVarMockRuntime()
	args := []any{"nonexistent"}

	// WHEN
	_, err := VAR(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for undefined variable, got nil")
	}
}

func TestVAR_DifferentTypes(t *testing.T) {
	// GIVEN
	rt := NewVarMockRuntime()

	// Test with string
	rt.SetVar("name", "John")
	args1 := []any{"name"}

	// Test with boolean
	rt.SetVar("active", true)
	args2 := []any{"active"}

	// Test with float
	rt.SetVar("price", 19.99)
	args3 := []any{"price"}

	// WHEN & THEN - String
	result1, err := VAR(args1, rt)
	if err != nil {
		t.Errorf("Expected no error for string var, got %v", err)
	}
	if result1 != "John" {
		t.Errorf("Expected 'John', got %v", result1)
	}

	// WHEN & THEN - Boolean
	result2, err := VAR(args2, rt)
	if err != nil {
		t.Errorf("Expected no error for boolean var, got %v", err)
	}
	if result2 != true {
		t.Errorf("Expected true, got %v", result2)
	}

	// WHEN & THEN - Float
	result3, err := VAR(args3, rt)
	if err != nil {
		t.Errorf("Expected no error for float var, got %v", err)
	}
	if result3 != 19.99 {
		t.Errorf("Expected 19.99, got %v", result3)
	}
}

func TestVAR_NonStringArgument(t *testing.T) {
	// GIVEN
	rt := NewVarMockRuntime()
	rt.SetVar("5", "value")
	args := []any{5} // Numeric argument, should be converted to string "5"

	// WHEN
	result, err := VAR(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error due to string conversion, got %v", err)
	}
	if result != "value" {
		t.Errorf("Expected 'value', got %v", result)
	}
}

func TestVAR_NoArguments(t *testing.T) {
	// GIVEN
	rt := NewVarMockRuntime()
	args := []any{}

	// WHEN
	_, err := VAR(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
}

func TestVAR_TooManyArguments(t *testing.T) {
	// GIVEN
	rt := NewVarMockRuntime()
	rt.SetVar("total", 100)
	args := []any{"total", "extra"}

	// WHEN
	result, err := VAR(args, rt)

	// THEN
	// The VAR function doesn't currently check for too many arguments,
	// it just uses the first one. This test verifies current behavior.
	if err != nil {
		t.Errorf("Expected no error despite extra args, got %v", err)
	}
	if result != 100 {
		t.Errorf("Expected 100 (first arg used), got %v", result)
	}
}
