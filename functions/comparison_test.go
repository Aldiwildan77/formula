package functions

import (
	"testing"
)

func TestEQ_EqualValues(t *testing.T) {
	// GIVEN
	args := []any{5, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := EQ(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestEQ_DifferentValues(t *testing.T) {
	// GIVEN
	args := []any{5, 10}
	rt := &MockRuntime{}

	// WHEN
	result, err := EQ(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestEQ_TypeCoercion(t *testing.T) {
	// GIVEN
	args := []any{5, "5"}
	rt := &MockRuntime{}

	// WHEN
	result, err := EQ(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true due to string conversion, got %v", result)
	}
}

func TestEQ_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5, 5, 5}
	rt := &MockRuntime{}

	// WHEN
	_, err := EQ(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestNEQ_NotEqualValues(t *testing.T) {
	// GIVEN
	args := []any{5, 10}
	rt := &MockRuntime{}

	// WHEN
	result, err := NEQ(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestNEQ_EqualValues(t *testing.T) {
	// GIVEN
	args := []any{5, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := NEQ(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestNEQ_TypeCoercion(t *testing.T) {
	// GIVEN
	args := []any{5, "5"}
	rt := &MockRuntime{}

	// WHEN
	result, err := NEQ(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false due to string conversion, got %v", result)
	}
}

func TestNEQ_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5, 5, 5}
	rt := &MockRuntime{}

	// WHEN
	_, err := NEQ(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestGT_GreaterThan(t *testing.T) {
	// GIVEN
	args := []any{10, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := GT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestGT_NotGreaterThan(t *testing.T) {
	// GIVEN
	args := []any{5, 10}
	rt := &MockRuntime{}

	// WHEN
	result, err := GT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestGT_EqualValues(t *testing.T) {
	// GIVEN
	args := []any{5, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := GT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false for equal values, got %v", result)
	}
}

func TestGT_InvalidSecondArgument(t *testing.T) {
	// GIVEN
	args := []any{5, "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := GT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestGT_InvalidFirstArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number", "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := GT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestGT_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5, 10, 15}
	rt := &MockRuntime{}

	// WHEN
	_, err := GT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestGTE_GreaterThan(t *testing.T) {
	// GIVEN
	args := []any{10, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := GTE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestGTE_EqualValues(t *testing.T) {
	// GIVEN
	args := []any{5, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := GTE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true for equal values, got %v", result)
	}
}

func TestGTE_LessThan(t *testing.T) {
	// GIVEN
	args := []any{5, 10}
	rt := &MockRuntime{}

	// WHEN
	result, err := GTE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestGTE_InvalidFirstArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number", "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := GTE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestGTE_InvalidSecondArgument(t *testing.T) {
	// GIVEN
	args := []any{5, "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := GTE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestGTE_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5, 10, 15}
	rt := &MockRuntime{}

	// WHEN
	_, err := GTE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestLT_LessThan(t *testing.T) {
	// GIVEN
	args := []any{5, 10}
	rt := &MockRuntime{}

	// WHEN
	result, err := LT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestLT_NotLessThan(t *testing.T) {
	// GIVEN
	args := []any{10, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := LT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestLT_EqualValues(t *testing.T) {
	// GIVEN
	args := []any{5, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := LT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false for equal values, got %v", result)
	}
}

func TestLT_InvalidFirstArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number", "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := LT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestLT_InvalidSecondArgument(t *testing.T) {
	// GIVEN
	args := []any{5, "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := LT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestLT_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5, 10, 15}
	rt := &MockRuntime{}

	// WHEN
	_, err := LT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestLTE_LessThan(t *testing.T) {
	// GIVEN
	args := []any{5, 10}
	rt := &MockRuntime{}

	// WHEN
	result, err := LTE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestLTE_EqualValues(t *testing.T) {
	// GIVEN
	args := []any{5, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := LTE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true for equal values, got %v", result)
	}
}

func TestLTE_GreaterThan(t *testing.T) {
	// GIVEN
	args := []any{10, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := LTE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestLTE_InvalidFirstArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number", "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := LTE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestLTE_InvalidSecondArgument(t *testing.T) {
	// GIVEN
	args := []any{5, "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := LTE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestLTE_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5, 10, 15}
	rt := &MockRuntime{}

	// WHEN
	_, err := LTE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}
