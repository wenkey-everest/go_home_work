package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("should add two numbers", func(t *testing.T) {
		result := add(1, 2)
		if result != 3 {
			t.Error("result should be", result)
		}
	})

	t.Run("Should subtract two numbers", func(t *testing.T) {
		result := sub(1, 2)
		if result != -1 {
			t.Error("result should be", result)
		}
	})

	t.Run("Should multiple two numbers", func(t *testing.T) {
		result := mul(1, 2)
		if result != 2 {
			t.Error("result should be", result)
		}
	})

	t.Run("Should division two numbers", func(t *testing.T) {
		result := div(4, 2)
		if result != 2 {
			t.Error("result should be", result)
		}
	})

}
