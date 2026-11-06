package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Fatalf("expected 5, got %d", result)
	}
}

func TestDivide(t *testing.T) {
	val, err := divide(10, 2)
	if err != nil || val != 5 {
		t.Fatalf("expected 5 without error, got %f, %v", val, err)
	}

	_, err = divide(1, 0)
	if err == nil {
		t.Fatal("expected error on division by zero")
	}
}
