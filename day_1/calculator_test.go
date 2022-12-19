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

	t.Run("Should multiply two numbers", func(t *testing.T) {
		result := mul(1, 2)
		if result != 2 {
			t.Error("result should be", result)
		}
	})

	t.Run("Should divide two numbers", func(t *testing.T) {
		result := div(4, 2)
		if result != 2 {
			t.Error("result should be", result)
		}
	})
	t.Run("Should get value of sin", func(t *testing.T) {
		result := sin(0.25)
		if result != 0.004363309284746571 {
			t.Error("result should be", result)
		}
	})

	t.Run("Should get value of cos", func(t *testing.T) {
		result := cos(0.25)
		if result != 0.9999904807207345 {
			t.Error("result should be", result)
		}
	})

	t.Run("Should get value of tan", func(t *testing.T) {
		result := tan(1)
		if result != 0.017455064928217585 {
			t.Error("result should be", result)
		}
	})

}
