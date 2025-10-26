package functions

import (
	"testing"
)

func TestADD_MultipleNumbers(t *testing.T) {
	// GIVEN
	args := []any{5, 10, 15}
	rt := &MockRuntime{}

	// WHEN
	result, err := ADD(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(30) {
		t.Errorf("Expected 30, got %v", result)
	}
}

func TestADD_SingleNumber(t *testing.T) {
	// GIVEN
	args := []any{42}
	rt := &MockRuntime{}

	// WHEN
	result, err := ADD(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(42) {
		t.Errorf("Expected 42, got %v", result)
	}
}

func TestADD_NoArguments(t *testing.T) {
	// GIVEN
	args := []any{}
	rt := &MockRuntime{}

	// WHEN
	_, err := ADD(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
}

func TestADD_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{5, "not a number", 15}
	rt := &MockRuntime{}

	// WHEN
	_, err := ADD(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestSUBTRACT_Basic(t *testing.T) {
	// GIVEN
	args := []any{10, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := SUBTRACT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(5) {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestSUBTRACT_MultipleOperands(t *testing.T) {
	// GIVEN
	args := []any{20, 5, 3}
	rt := &MockRuntime{}

	// WHEN
	result, err := SUBTRACT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(12) {
		t.Errorf("Expected 12, got %v", result)
	}
}

func TestSUBTRACT_NoArguments(t *testing.T) {
	// GIVEN
	args := []any{}
	rt := &MockRuntime{}

	// WHEN
	_, err := SUBTRACT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
}

func TestSUBTRACT_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{10, "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := SUBTRACT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestMULTIPLY_Basic(t *testing.T) {
	// GIVEN
	args := []any{5, 4}
	rt := &MockRuntime{}

	// WHEN
	result, err := MULTIPLY(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(20) {
		t.Errorf("Expected 20, got %v", result)
	}
}

func TestMULTIPLY_MultipleOperands(t *testing.T) {
	// GIVEN
	args := []any{2, 3, 4}
	rt := &MockRuntime{}

	// WHEN
	result, err := MULTIPLY(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(24) {
		t.Errorf("Expected 24, got %v", result)
	}
}

func TestMULTIPLY_WithZero(t *testing.T) {
	// GIVEN
	args := []any{5, 0, 10}
	rt := &MockRuntime{}

	// WHEN
	result, err := MULTIPLY(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(0) {
		t.Errorf("Expected 0, got %v", result)
	}
}

func TestMULTIPLY_NoArguments(t *testing.T) {
	// GIVEN
	args := []any{}
	rt := &MockRuntime{}

	// WHEN
	_, err := MULTIPLY(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
}

func TestMULTIPLY_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{5, "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := MULTIPLY(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestDIVIDE_Basic(t *testing.T) {
	// GIVEN
	args := []any{20, 4}
	rt := &MockRuntime{}

	// WHEN
	result, err := DIVIDE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(5) {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestDIVIDE_MultipleOperands(t *testing.T) {
	// GIVEN
	args := []any{100, 5, 2}
	rt := &MockRuntime{}

	// WHEN
	result, err := DIVIDE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(10) {
		t.Errorf("Expected 10, got %v", result)
	}
}

func TestDIVIDE_ByZero(t *testing.T) {
	// GIVEN
	args := []any{10, 0}
	rt := &MockRuntime{}

	// WHEN
	_, err := DIVIDE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected division by zero error, got nil")
	}
	if err.Error() != "DIVIDE by zero is not allowed" {
		t.Errorf("Expected specific error message, got: %s", err.Error())
	}
}

func TestDIVIDE_NoArguments(t *testing.T) {
	// GIVEN
	args := []any{}
	rt := &MockRuntime{}

	// WHEN
	_, err := DIVIDE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
}

func TestDIVIDE_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{10, "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := DIVIDE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestMOD_Basic(t *testing.T) {
	// GIVEN
	args := []any{10, 3}
	rt := &MockRuntime{}

	// WHEN
	result, err := MOD(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(1) {
		t.Errorf("Expected 1, got %v", result)
	}
}

func TestMOD_NegativeNumbers(t *testing.T) {
	// GIVEN
	args := []any{-10, 3}
	rt := &MockRuntime{}

	// WHEN
	result, err := MOD(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// Go's math.Mod returns -1 for -10 % 3
	if result != float64(-1) {
		t.Errorf("Expected -1, got %v", result)
	}
}

func TestMOD_ByZero(t *testing.T) {
	// GIVEN
	args := []any{10, 0}
	rt := &MockRuntime{}

	// WHEN
	_, err := MOD(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected modulo by zero error, got nil")
	}
	if err.Error() != "MOD by zero is not allowed" {
		t.Errorf("Expected specific error message, got: %s", err.Error())
	}
}

func TestMOD_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{10, 5, 2}
	rt := &MockRuntime{}

	// WHEN
	_, err := MOD(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestMOD_InvalidFirstArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number", "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := MOD(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestMOD_InvalidSecondArgument(t *testing.T) {
	// GIVEN
	args := []any{10, "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := MOD(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestPOWER_Basic(t *testing.T) {
	// GIVEN
	args := []any{2, 3}
	rt := &MockRuntime{}

	// WHEN
	result, err := POWER(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(8) {
		t.Errorf("Expected 8, got %v", result)
	}
}

func TestPOWER_ZeroExponent(t *testing.T) {
	// GIVEN
	args := []any{5, 0}
	rt := &MockRuntime{}

	// WHEN
	result, err := POWER(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(1) {
		t.Errorf("Expected 1, got %v", result)
	}
}

func TestPOWER_NegativeExponent(t *testing.T) {
	// GIVEN
	args := []any{4, -2}
	rt := &MockRuntime{}

	// WHEN
	result, err := POWER(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(0.0625) {
		t.Errorf("Expected 0.0625, got %v", result)
	}
}

func TestPOWER_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{2, 3, 4}
	rt := &MockRuntime{}

	// WHEN
	_, err := POWER(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestPOWER_InvalidFirstArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number", "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := POWER(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestPOWER_InvalidSecondArgument(t *testing.T) {
	// GIVEN
	args := []any{2, "not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := POWER(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestNEGATE_PositiveToNegative(t *testing.T) {
	// GIVEN
	args := []any{5}
	rt := &MockRuntime{}

	// WHEN
	result, err := NEGATE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(-5) {
		t.Errorf("Expected -5, got %v", result)
	}
}

func TestNEGATE_NegativeToPositive(t *testing.T) {
	// GIVEN
	args := []any{-10}
	rt := &MockRuntime{}

	// WHEN
	result, err := NEGATE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(10) {
		t.Errorf("Expected 10, got %v", result)
	}
}

func TestNEGATE_Zero(t *testing.T) {
	// GIVEN
	args := []any{0}
	rt := &MockRuntime{}

	// WHEN
	result, err := NEGATE(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(0) {
		t.Errorf("Expected 0, got %v", result)
	}
}

func TestNEGATE_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5, 10}
	rt := &MockRuntime{}

	// WHEN
	_, err := NEGATE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestNEGATE_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := NEGATE(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}
