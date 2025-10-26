package functions

import (
	"testing"
)

func TestIF_ConditionTrue(t *testing.T) {
	// GIVEN
	args := []any{true, "then-value", "else-value"}
	rt := &MockRuntime{}

	// WHEN
	result, err := IF(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != "then-value" {
		t.Errorf("Expected 'then-value', got %v", result)
	}
}

func TestIF_ConditionFalse(t *testing.T) {
	// GIVEN
	args := []any{false, "then-value", "else-value"}
	rt := &MockRuntime{}

	// WHEN
	result, err := IF(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != "else-value" {
		t.Errorf("Expected 'else-value', got %v", result)
	}
}

func TestIF_NonBooleanCondition(t *testing.T) {
	// GIVEN
	args := []any{"not a boolean", "then-value", "else-value"}
	rt := &MockRuntime{}

	// WHEN
	_, err := IF(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for non-boolean condition, got nil")
	}
}

func TestIF_WrongArgumentCount_TooFew(t *testing.T) {
	// GIVEN
	args := []any{true, "then-value"}
	rt := &MockRuntime{}

	// WHEN
	_, err := IF(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for too few arguments, got nil")
	}
}

func TestIF_WrongArgumentCount_TooMany(t *testing.T) {
	// GIVEN
	args := []any{true, "then-value", "else-value", "extra"}
	rt := &MockRuntime{}

	// WHEN
	_, err := IF(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for too many arguments, got nil")
	}
}

func TestAND_AllTrue(t *testing.T) {
	// GIVEN
	args := []any{true, true, true}
	rt := &MockRuntime{}

	// WHEN
	result, err := AND(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestAND_OneFalse(t *testing.T) {
	// GIVEN
	args := []any{true, false, true}
	rt := &MockRuntime{}

	// WHEN
	result, err := AND(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestAND_AllFalse(t *testing.T) {
	// GIVEN
	args := []any{false, false, false}
	rt := &MockRuntime{}

	// WHEN
	result, err := AND(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestAND_ShortCircuit(t *testing.T) {
	// GIVEN
	args := []any{false, true, true}
	rt := &MockRuntime{}

	// WHEN
	result, err := AND(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false (short circuit), got %v", result)
	}
}

func TestAND_NonBooleanArgument(t *testing.T) {
	// GIVEN
	args := []any{true, "not a boolean", true}
	rt := &MockRuntime{}

	// WHEN
	_, err := AND(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for non-boolean argument, got nil")
	}
}

func TestAND_TooFewArguments(t *testing.T) {
	// GIVEN
	args := []any{true}
	rt := &MockRuntime{}

	// WHEN
	_, err := AND(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for too few arguments, got nil")
	}
}

func TestOR_OneTrue(t *testing.T) {
	// GIVEN
	args := []any{false, true, false}
	rt := &MockRuntime{}

	// WHEN
	result, err := OR(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestOR_AllTrue(t *testing.T) {
	// GIVEN
	args := []any{true, true, true}
	rt := &MockRuntime{}

	// WHEN
	result, err := OR(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestOR_AllFalse(t *testing.T) {
	// GIVEN
	args := []any{false, false, false}
	rt := &MockRuntime{}

	// WHEN
	result, err := OR(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestOR_ShortCircuit(t *testing.T) {
	// GIVEN
	args := []any{true, false, false}
	rt := &MockRuntime{}

	// WHEN
	result, err := OR(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true (short circuit), got %v", result)
	}
}

func TestOR_NonBooleanArgument(t *testing.T) {
	// GIVEN
	args := []any{false, "not a boolean", false}
	rt := &MockRuntime{}

	// WHEN
	_, err := OR(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for non-boolean argument, got nil")
	}
}

func TestOR_TooFewArguments(t *testing.T) {
	// GIVEN
	args := []any{false}
	rt := &MockRuntime{}

	// WHEN
	_, err := OR(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for too few arguments, got nil")
	}
}

func TestNOT_TrueToFalse(t *testing.T) {
	// GIVEN
	args := []any{true}
	rt := &MockRuntime{}

	// WHEN
	result, err := NOT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestNOT_FalseToTrue(t *testing.T) {
	// GIVEN
	args := []any{false}
	rt := &MockRuntime{}

	// WHEN
	result, err := NOT(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestNOT_NonBooleanArgument(t *testing.T) {
	// GIVEN
	args := []any{"not a boolean"}
	rt := &MockRuntime{}

	// WHEN
	_, err := NOT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for non-boolean argument, got nil")
	}
}

func TestNOT_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{true, false}
	rt := &MockRuntime{}

	// WHEN
	_, err := NOT(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}
