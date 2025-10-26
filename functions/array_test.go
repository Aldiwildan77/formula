package functions

import (
	"testing"
)

// ARRAY_LENGTH tests

func TestARRAY_LENGTH_ValidArray(t *testing.T) {
	// GIVEN
	args := []any{[]any{1, 2, 3, 4, 5}}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_LENGTH(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 5 {
		t.Errorf("Expected length 5, got %v", result)
	}
}

func TestARRAY_LENGTH_EmptyArray(t *testing.T) {
	// GIVEN
	args := []any{[]any{}}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_LENGTH(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 0 {
		t.Errorf("Expected length 0, got %v", result)
	}
}

func TestARRAY_LENGTH_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{[]any{1, 2, 3}, "extra arg"}
	rt := &MockRuntime{}

	// WHEN
	_, err := ARRAY_LENGTH(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestARRAY_LENGTH_NonArrayArgument(t *testing.T) {
	// GIVEN
	args := []any{"not an array"}
	rt := &MockRuntime{}

	// WHEN
	_, err := ARRAY_LENGTH(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for non-array argument, got nil")
	}
}

// ARRAY_INDEX tests

func TestARRAY_INDEX_ValidIndex(t *testing.T) {
	// GIVEN
	args := []any{[]any{"a", "b", "c", "d"}, 2}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_INDEX(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != "c" {
		t.Errorf("Expected 'c', got %v", result)
	}
}

func TestARRAY_INDEX_FirstElement(t *testing.T) {
	// GIVEN
	args := []any{[]any{10, 20, 30}, 0}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_INDEX(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 10 {
		t.Errorf("Expected 10, got %v", result)
	}
}

func TestARRAY_INDEX_LastElement(t *testing.T) {
	// GIVEN
	args := []any{[]any{10, 20, 30}, 2}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_INDEX(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 30 {
		t.Errorf("Expected 30, got %v", result)
	}
}

func TestARRAY_INDEX_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{[]any{1, 2, 3}}
	rt := &MockRuntime{}

	// WHEN
	_, err := ARRAY_INDEX(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestARRAY_INDEX_NonArrayArgument(t *testing.T) {
	// GIVEN
	args := []any{"not an array", 0}
	rt := &MockRuntime{}

	// WHEN
	_, err := ARRAY_INDEX(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for non-array argument, got nil")
	}
}

func TestARRAY_INDEX_IndexOutOfBounds(t *testing.T) {
	// GIVEN
	args := []any{[]any{1, 2, 3}, 5}
	rt := &MockRuntime{}

	// WHEN
	_, err := ARRAY_INDEX(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for index out of bounds, got nil")
	}
}

func TestARRAY_INDEX_NegativeIndex(t *testing.T) {
	// GIVEN
	args := []any{[]any{1, 2, 3}, -1}
	rt := &MockRuntime{}

	// WHEN
	_, err := ARRAY_INDEX(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for negative index, got nil")
	}
}

// ARRAY_FILTER tests

func TestARRAY_FILTER_FindElement(t *testing.T) {
	// GIVEN
	args := []any{[]any{1, 2, 3, 4, 5}, 3}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_FILTER(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 1 {
		t.Errorf("Expected first non-matching element (1), got %v", result)
	}
}

func TestARRAY_FILTER_AllMatch(t *testing.T) {
	// GIVEN
	args := []any{[]any{5, 5, 5}, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_FILTER(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil when all elements match, got %v", result)
	}
}

func TestARRAY_FILTER_EmptyArray(t *testing.T) {
	// GIVEN
	args := []any{[]any{}, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_FILTER(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil for empty array, got %v", result)
	}
}

func TestARRAY_FILTER_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{[]any{1, 2, 3}}
	rt := &MockRuntime{}

	// WHEN
	_, err := ARRAY_FILTER(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestARRAY_FILTER_NonArrayArgument(t *testing.T) {
	// GIVEN
	args := []any{"not an array", 5}
	rt := &MockRuntime{}

	// WHEN
	_, err := ARRAY_FILTER(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for non-array argument, got nil")
	}
}

// ARRAY_MAP tests

func TestARRAY_MAP_FindElement(t *testing.T) {
	// GIVEN
	args := []any{[]any{1, 2, 3, 4, 5}, 3}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_MAP(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != 1 {
		t.Errorf("Expected first non-matching element (1), got %v", result)
	}
}

func TestARRAY_MAP_AllMatch(t *testing.T) {
	// GIVEN
	args := []any{[]any{5, 5, 5}, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_MAP(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil when all elements match, got %v", result)
	}
}

func TestARRAY_MAP_EmptyArray(t *testing.T) {
	// GIVEN
	args := []any{[]any{}, 5}
	rt := &MockRuntime{}

	// WHEN
	result, err := ARRAY_MAP(args, rt)

	// THEN
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil for empty array, got %v", result)
	}
}

func TestARRAY_MAP_WrongArgumentCount(t *testing.T) {
	// GIVEN
	args := []any{[]any{1, 2, 3}}
	rt := &MockRuntime{}

	// WHEN
	_, err := ARRAY_MAP(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for wrong argument count, got nil")
	}
}

func TestARRAY_MAP_NonArrayArgument(t *testing.T) {
	// GIVEN
	args := []any{"not an array", 5}
	rt := &MockRuntime{}

	// WHEN
	_, err := ARRAY_MAP(args, rt)

	// THEN
	if err == nil {
		t.Error("Expected error for non-array argument, got nil")
	}
}
