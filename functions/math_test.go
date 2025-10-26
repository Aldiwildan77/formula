package functions

import (
	"math"
	"testing"
)

func TestSQRT_PerfectSquare(t *testing.T) {
	// GIVEN
	args := []any{16}
	rt := &MockRuntime{}

	// WHEN
	result, err := SQRT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(4) {
		t.Errorf("Expected 4, got %v", result)
	}
}

func TestSQRT_Decimal(t *testing.T) {
	// GIVEN
	args := []any{2}
	rt := &MockRuntime{}

	// WHEN
	result, err := SQRT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := math.Sqrt(2)
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSQRT_Zero(t *testing.T) {
	// GIVEN
	args := []any{0}
	rt := &MockRuntime{}

	// WHEN
	result, err := SQRT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(0) {
		t.Errorf("Expected 0, got %v", result)
	}
}

func TestSQRT_NegativeNumber(t *testing.T) {
	// GIVEN
	args := []any{-4}
	rt := &MockRuntime{}

	// WHEN
	_, err := SQRT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for negative number, got nil")
	}
}

func TestSQRT_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := SQRT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestSQRT_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{4, 9}
	rt := &MockRuntime{}

	// WHEN
	_, err := SQRT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestABS_PositiveNumber(t *testing.T) {
	// GIVEN
	args := []any{5}
	rt := &MockRuntime{}

	// WHEN
	result, err := ABS(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(5) {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestABS_NegativeNumber(t *testing.T) {
	// GIVEN
	args := []any{-5}
	rt := &MockRuntime{}

	// WHEN
	result, err := ABS(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(5) {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestABS_Zero(t *testing.T) {
	// GIVEN
	args := []any{0}
	rt := &MockRuntime{}

	// WHEN
	result, err := ABS(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(0) {
		t.Errorf("Expected 0, got %v", result)
	}
}

func TestABS_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := ABS(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestABS_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5, 10}
	rt := &MockRuntime{}

	// WHEN
	_, err := ABS(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestFLOOR_PositiveDecimal(t *testing.T) {
	// GIVEN
	args := []any{5.7}
	rt := &MockRuntime{}

	// WHEN
	result, err := FLOOR(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(5) {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestFLOOR_NegativeDecimal(t *testing.T) {
	// GIVEN
	args := []any{-5.7}
	rt := &MockRuntime{}

	// WHEN
	result, err := FLOOR(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(-6) {
		t.Errorf("Expected -6, got %v", result)
	}
}

func TestFLOOR_Integer(t *testing.T) {
	// GIVEN
	args := []any{5}
	rt := &MockRuntime{}

	// WHEN
	result, err := FLOOR(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(5) {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestFLOOR_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := FLOOR(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestFLOOR_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5.7, 10.2}
	rt := &MockRuntime{}

	// WHEN
	_, err := FLOOR(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestCEIL_PositiveDecimal(t *testing.T) {
	// GIVEN
	args := []any{5.2}
	rt := &MockRuntime{}

	// WHEN
	result, err := CEIL(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(6) {
		t.Errorf("Expected 6, got %v", result)
	}
}

func TestCEIL_NegativeDecimal(t *testing.T) {
	// GIVEN
	args := []any{-5.2}
	rt := &MockRuntime{}

	// WHEN
	result, err := CEIL(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(-5) {
		t.Errorf("Expected -5, got %v", result)
	}
}

func TestCEIL_Integer(t *testing.T) {
	// GIVEN
	args := []any{5}
	rt := &MockRuntime{}

	// WHEN
	result, err := CEIL(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(5) {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestCEIL_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := CEIL(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestCEIL_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5.2, 10.7}
	rt := &MockRuntime{}

	// WHEN
	_, err := CEIL(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestROUND_RoundUp(t *testing.T) {
	// GIVEN
	args := []any{5.7}
	rt := &MockRuntime{}

	// WHEN
	result, err := ROUND(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(6) {
		t.Errorf("Expected 6, got %v", result)
	}
}

func TestROUND_RoundDown(t *testing.T) {
	// GIVEN
	args := []any{5.2}
	rt := &MockRuntime{}

	// WHEN
	result, err := ROUND(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(5) {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestROUND_HalfwayCase(t *testing.T) {
	// GIVEN
	args := []any{5.5}
	rt := &MockRuntime{}

	// WHEN
	result, err := ROUND(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(6) {
		t.Errorf("Expected 6 (round half up), got %v", result)
	}
}

func TestROUND_NegativeNumber(t *testing.T) {
	// GIVEN
	args := []any{-5.7}
	rt := &MockRuntime{}

	// WHEN
	result, err := ROUND(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(-6) {
		t.Errorf("Expected -6, got %v", result)
	}
}

func TestROUND_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a number"}
	rt := &MockRuntime{}

	// WHEN
	_, err := ROUND(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestROUND_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{5.5, 10.5}
	rt := &MockRuntime{}

	// WHEN
	_, err := ROUND(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestMIN_TwoNumbers(t *testing.T) {
	// GIVEN
	args := []any{10, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := MIN(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(5) {
		t.Errorf("Expected 5, got %v", result)
	}
}

func TestMIN_MultipleNumbers(t *testing.T) {
	// GIVEN
	args := []any{10, 5, 8, 3, 7}
	rt := &MockRuntime{}

	// WHEN
	result, err := MIN(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(3) {
		t.Errorf("Expected 3, got %v", result)
	}
}

func TestMIN_SingleNumber(t *testing.T) {
	// GIVEN
	args := []any{42}
	rt := &MockRuntime{}

	// WHEN
	result, err := MIN(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(42) {
		t.Errorf("Expected 42, got %v", result)
	}
}

func TestMIN_NegativeNumbers(t *testing.T) {
	// GIVEN
	args := []any{-5, -10, -3}
	rt := &MockRuntime{}

	// WHEN
	result, err := MIN(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(-10) {
		t.Errorf("Expected -10, got %v", result)
	}
}

func TestMIN_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{5, "not a number", 10}
	rt := &MockRuntime{}

	// WHEN
	_, err := MIN(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestMIN_NoArguments(t *testing.T) {
	// GIVEN
	args := []any{}
	rt := &MockRuntime{}

	// WHEN
	_, err := MIN(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
}

func TestMAX_TwoNumbers(t *testing.T) {
	// GIVEN
	args := []any{10, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := MAX(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(10) {
		t.Errorf("Expected 10, got %v", result)
	}
}

func TestMAX_MultipleNumbers(t *testing.T) {
	// GIVEN
	args := []any{10, 5, 8, 15, 7}
	rt := &MockRuntime{}

	// WHEN
	result, err := MAX(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(15) {
		t.Errorf("Expected 15, got %v", result)
	}
}

func TestMAX_SingleNumber(t *testing.T) {
	// GIVEN
	args := []any{42}
	rt := &MockRuntime{}

	// WHEN
	result, err := MAX(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(42) {
		t.Errorf("Expected 42, got %v", result)
	}
}

func TestMAX_NegativeNumbers(t *testing.T) {
	// GIVEN
	args := []any{-5, -10, -3}
	rt := &MockRuntime{}

	// WHEN
	result, err := MAX(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != float64(-3) {
		t.Errorf("Expected -3, got %v", result)
	}
}

func TestMAX_InvalidArgument(t *testing.T) {
	// GIVEN
	args := []any{5, "not a number", 10}
	rt := &MockRuntime{}

	// WHEN
	_, err := MAX(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid argument, got nil")
	}
}

func TestMAX_NoArguments(t *testing.T) {
	// GIVEN
	args := []any{}
	rt := &MockRuntime{}

	// WHEN
	_, err := MAX(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}
}
