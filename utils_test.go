package utils

import (
	"testing"
)

func TestContains(t *testing.T) {
	t.Run("int scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []int{1, 2, 3, 4}
			element := 2
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []int{1, 2, 3, 4}
			element := 5
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []int{}
			element := 5
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []int(nil)
			element := 5
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("int8 scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []int8{1, 2, 3, 4}
			element := int8(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []int8{1, 2, 3, 4}
			element := int8(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []int8{}
			element := int8(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []int8(nil)
			element := int8(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("int16 scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []int16{1, 2, 3, 4}
			element := int16(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []int16{1, 2, 3, 4}
			element := int16(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []int16{}
			element := int16(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []int16(nil)
			element := int16(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("int32 scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []int32{1, 2, 3, 4}
			element := int32(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []int32{1, 2, 3, 4}
			element := int32(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []int32{}
			element := int32(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []int32(nil)
			element := int32(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("int64 scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []int64{1, 2, 3, 4}
			element := int64(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []int64{1, 2, 3, 4}
			element := int64(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []int64{}
			element := int64(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []int64(nil)
			element := int64(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("float32 scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []float32{1, 2, 3, 4}
			element := float32(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []float32{1, 2, 3, 4}
			element := float32(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []float32{}
			element := float32(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []float32(nil)
			element := float32(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("float64 scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []float64{1, 2, 3, 4}
			element := float64(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []float64{1, 2, 3, 4}
			element := float64(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []float64{}
			element := float64(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []float64(nil)
			element := float64(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("uint scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []uint{1, 2, 3, 4}
			element := uint(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []uint{1, 2, 3, 4}
			element := uint(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []uint{}
			element := uint(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []uint(nil)
			element := uint(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("uint8 scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []uint8{1, 2, 3, 4}
			element := uint8(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []uint8{1, 2, 3, 4}
			element := uint8(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []uint8{}
			element := uint8(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []uint8(nil)
			element := uint8(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("uint16 scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []uint16{1, 2, 3, 4}
			element := uint16(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []uint16{1, 2, 3, 4}
			element := uint16(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []uint16{}
			element := uint16(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []uint16(nil)
			element := uint16(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("uint32 scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []uint32{1, 2, 3, 4}
			element := uint32(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []uint32{1, 2, 3, 4}
			element := uint32(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []uint32{}
			element := uint32(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []uint32(nil)
			element := uint32(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("uint64 scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []uint64{1, 2, 3, 4}
			element := uint64(2)
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []uint64{1, 2, 3, 4}
			element := uint64(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []uint64{}
			element := uint64(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []uint64(nil)
			element := uint64(5)
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})

	t.Run("string scenarios", func(t *testing.T) {
		t.Run("returns true if element is in slice", func(t *testing.T) {
			input := []string{"1", "2", "3", "4"}
			element := "2"
			result := Contains(input, element)

			if !result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if element is not in slice", func(t *testing.T) {
			input := []string{"1", "2", "3", "4"}
			element := "5"
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is empty", func(t *testing.T) {
			input := []string{}
			element := "5"
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})

		t.Run("returns false if slice is nil", func(t *testing.T) {
			input := []string(nil)
			element := "5"
			result := Contains(input, element)

			if result {
				t.Errorf("test failed")
			}
		})
	})
}

func TestSum(t *testing.T) {
	result := 10
	unexpectedResult := 11

	t.Run("int scenarions", func(t *testing.T) {
		t.Run("sums slice of int", func(t *testing.T) {
			input := []int{1, 2, 3, 4}
			expectedOutput := result
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of int", func(t *testing.T) {
			input := []int{1, 2, 3, 4}
			unexpectedOutput := unexpectedResult
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint scenarios", func(t *testing.T) {
		t.Run("sums slice of int8", func(t *testing.T) {
			input := []int8{1, 2, 3, 4}
			expectedOutput := int8(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of int8", func(t *testing.T) {
			input := []int8{1, 2, 3, 4}
			unexpectedOutput := int8(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint16 scenarios", func(t *testing.T) {
		t.Run("sums slice of int16", func(t *testing.T) {
			input := []int16{1, 2, 3, 4}
			expectedOutput := int16(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of int16", func(t *testing.T) {
			input := []int16{1, 2, 3, 4}
			unexpectedOutput := int16(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("int32 scenarios", func(t *testing.T) {
		t.Run("sums slice of int32", func(t *testing.T) {
			input := []int32{1, 2, 3, 4}
			expectedOutput := int32(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of int32", func(t *testing.T) {
			input := []int32{1, 2, 3, 4}
			unexpectedOutput := int32(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("int64 scenarios", func(t *testing.T) {
		t.Run("sums slice of int64", func(t *testing.T) {
			input := []int64{1, 2, 3, 4}
			expectedOutput := int64(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of int64", func(t *testing.T) {
			input := []int64{1, 2, 3, 4}
			unexpectedOutput := int64(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("float32 scenarios", func(t *testing.T) {
		t.Run("sums slice of float32", func(t *testing.T) {
			input := []float32{1, 2, 3, 4}
			expectedOutput := float32(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of float32", func(t *testing.T) {
			input := []float32{1, 2, 3, 4}
			unexpectedOutput := float32(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %f, got %f", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("float64 scenarios", func(t *testing.T) {
		t.Run("sums slice of float64", func(t *testing.T) {
			input := []float64{1, 2, 3, 4}
			expectedOutput := float64(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of float64", func(t *testing.T) {
			input := []float64{1, 2, 3, 4}
			unexpectedOutput := float64(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %f, got %f", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint scenarios", func(t *testing.T) {
		t.Run("sums slice of uint", func(t *testing.T) {
			input := []uint{1, 2, 3, 4}
			expectedOutput := uint(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of uint", func(t *testing.T) {
			input := []uint{1, 2, 3, 4}
			unexpectedOutput := uint(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint8 scenarios", func(t *testing.T) {
		t.Run("sums slice of uint8", func(t *testing.T) {
			input := []uint8{1, 2, 3, 4}
			expectedOutput := uint8(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of uint8", func(t *testing.T) {
			input := []uint8{1, 2, 3, 4}
			unexpectedOutput := uint8(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint16 scenarios", func(t *testing.T) {
		t.Run("sums slice of uint16", func(t *testing.T) {
			input := []uint16{1, 2, 3, 4}
			expectedOutput := uint16(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of uint16", func(t *testing.T) {
			input := []uint16{1, 2, 3, 4}
			unexpectedOutput := uint16(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint32 scenarios", func(t *testing.T) {
		t.Run("sums slice of uint32", func(t *testing.T) {
			input := []uint32{1, 2, 3, 4}
			expectedOutput := uint32(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of uint32", func(t *testing.T) {
			input := []uint32{1, 2, 3, 4}
			unexpectedOutput := uint32(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint64 scenarios", func(t *testing.T) {
		t.Run("sums slice of uint64", func(t *testing.T) {
			input := []uint64{1, 2, 3, 4}
			expectedOutput := uint64(result)
			actualOutput := Sum(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("sums slice of uint64", func(t *testing.T) {
			input := []uint64{1, 2, 3, 4}
			unexpectedOutput := uint64(unexpectedResult)
			actualOutput := Sum(input)

			if actualOutput == unexpectedOutput {
				t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
			}
		})
	})
}

func TestSort(t *testing.T) {
	t.Run("sorts slice of int", func(t *testing.T) {
		input := []int{4, 3, 2, 1}
		expectedOutput := []int{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of int8", func(t *testing.T) {
		input := []int8{4, 3, 2, 1}
		expectedOutput := []int8{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of int16", func(t *testing.T) {
		input := []int16{4, 3, 2, 1}
		expectedOutput := []int16{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of int32", func(t *testing.T) {
		input := []int32{4, 3, 2, 1}
		expectedOutput := []int32{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of int64", func(t *testing.T) {
		input := []int64{4, 3, 2, 1}
		expectedOutput := []int64{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of uint", func(t *testing.T) {
		input := []uint{4, 3, 2, 1}
		expectedOutput := []uint{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of uint", func(t *testing.T) {
		input := []uint8{4, 3, 2, 1}
		expectedOutput := []uint8{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of uint16", func(t *testing.T) {
		input := []uint16{4, 3, 2, 1}
		expectedOutput := []uint16{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of uint32", func(t *testing.T) {
		input := []uint32{4, 3, 2, 1}
		expectedOutput := []uint32{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of uint64", func(t *testing.T) {
		input := []uint64{4, 3, 2, 1}
		expectedOutput := []uint64{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of float32", func(t *testing.T) {
		input := []float32{4, 3, 2, 1}
		expectedOutput := []float32{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %f, but got %f", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of float64", func(t *testing.T) {
		input := []float64{4, 3, 2, 1}
		expectedOutput := []float64{1, 2, 3, 4}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %f, but got %f", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of string", func(t *testing.T) {
		input := []string{"Zeynep", "Aksel"}
		expectedOutput := []string{"Aksel", "Zeynep"}
		Sort(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %s, but got %s", expectedOutput, input)
			}
		}
	})
}

func TestSortDesc(t *testing.T) {
	t.Run("sorts slice of int descending", func(t *testing.T) {
		input := []int{4, 2, 1, 3}
		expectedOutput := []int{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of int8 descending", func(t *testing.T) {
		input := []int8{4, 2, 1, 3}
		expectedOutput := []int8{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of int16 descending", func(t *testing.T) {
		input := []int16{4, 2, 1, 3}
		expectedOutput := []int16{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of int32 descending", func(t *testing.T) {
		input := []int32{4, 2, 1, 3}
		expectedOutput := []int32{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of int64 descending", func(t *testing.T) {
		input := []int64{4, 2, 1, 3}
		expectedOutput := []int64{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of uint descending", func(t *testing.T) {
		input := []uint{4, 2, 1, 3}
		expectedOutput := []uint{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of uint descending", func(t *testing.T) {
		input := []uint8{4, 2, 1, 3}
		expectedOutput := []uint8{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of uint16 descending", func(t *testing.T) {
		input := []uint16{4, 2, 1, 3}
		expectedOutput := []uint16{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of uint32 descending", func(t *testing.T) {
		input := []uint32{4, 2, 1, 3}
		expectedOutput := []uint32{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of uint64 descending", func(t *testing.T) {
		input := []uint64{4, 2, 1, 3}
		expectedOutput := []uint64{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %d, but got %d", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of float32 descending", func(t *testing.T) {
		input := []float32{4, 2, 1, 3}
		expectedOutput := []float32{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %f, but got %f", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of float64 descending", func(t *testing.T) {
		input := []float64{4, 2, 1, 3}
		expectedOutput := []float64{4, 3, 2, 1}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %f, but got %f", expectedOutput, input)
			}
		}
	})

	t.Run("sorts slice of string descending", func(t *testing.T) {
		input := []string{"Aksel", "Zeynep"}
		expectedOutput := []string{"Zeynep", "Aksel"}
		SortDesc(input)

		for i := 0; i < len(input); i++ {
			if input[i] != expectedOutput[i] {
				t.Errorf("Expected %s, but got %s", expectedOutput, input)
			}
		}
	})
}

func TestMin(t *testing.T) {
	result := 1

	t.Run("int scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []int{}
			expectedOutput := 0
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []int(nil)
			expectedOutput := 0
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []int{1, 2, 3, 4}
			expectedOutput := result
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("int8 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []int8{}
			expectedOutput := int8(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []int8(nil)
			expectedOutput := int8(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []int8{1, 2, 3, 4}
			expectedOutput := int8(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("int16 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []int16{}
			expectedOutput := int16(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []int16(nil)
			expectedOutput := int16(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []int16{1, 2, 3, 4}
			expectedOutput := int16(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("int32 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []int32{}
			expectedOutput := int32(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []int32(nil)
			expectedOutput := int32(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []int32{1, 2, 3, 4}
			expectedOutput := int32(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("int64 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []int64{}
			expectedOutput := int64(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []int64(nil)
			expectedOutput := int64(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []int64{1, 2, 3, 4}
			expectedOutput := int64(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("float32 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []float32{}
			expectedOutput := float32(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []float32(nil)
			expectedOutput := float32(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []float32{1, 2, 3, 4}
			expectedOutput := float32(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("float64 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []float64{}
			expectedOutput := float64(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []float64(nil)
			expectedOutput := float64(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []float64{1, 2, 3, 4}
			expectedOutput := float64(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []uint{}
			expectedOutput := uint(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []uint(nil)
			expectedOutput := uint(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []uint{1, 2, 3, 4}
			expectedOutput := uint(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint8 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []uint8{}
			expectedOutput := uint8(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []uint8(nil)
			expectedOutput := uint8(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []uint8{1, 2, 3, 4}
			expectedOutput := uint8(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint16 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []uint16{}
			expectedOutput := uint16(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []uint16(nil)
			expectedOutput := uint16(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []uint16{1, 2, 3, 4}
			expectedOutput := uint16(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint32 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []uint32{}
			expectedOutput := uint32(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []uint32(nil)
			expectedOutput := uint32(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []uint32{1, 2, 3, 4}
			expectedOutput := uint32(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint64 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []uint64{}
			expectedOutput := uint64(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []uint64(nil)
			expectedOutput := uint64(0)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []uint64{1, 2, 3, 4}
			expectedOutput := uint64(result)
			actualOutput := Min(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})
}

func TestMax(t *testing.T) {
	result := 4

	t.Run("int scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []int{}
			expectedOutput := 0
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []int(nil)
			expectedOutput := 0
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []int{1, 2, 3, 4}
			expectedOutput := result
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("int8 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []int8{}
			expectedOutput := int8(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []int8(nil)
			expectedOutput := int8(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []int8{1, 2, 3, 4}
			expectedOutput := int8(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("int16 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []int16{}
			expectedOutput := int16(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []int16(nil)
			expectedOutput := int16(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []int16{1, 2, 3, 4}
			expectedOutput := int16(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("int32 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []int32{}
			expectedOutput := int32(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []int32(nil)
			expectedOutput := int32(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []int32{1, 2, 3, 4}
			expectedOutput := int32(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("int64 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []int64{}
			expectedOutput := int64(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []int64(nil)
			expectedOutput := int64(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []int64{1, 2, 3, 4}
			expectedOutput := int64(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("float32 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []float32{}
			expectedOutput := float32(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []float32(nil)
			expectedOutput := float32(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []float32{1, 2, 3, 4}
			expectedOutput := float32(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("float64 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []float64{}
			expectedOutput := float64(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []float64(nil)
			expectedOutput := float64(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []float64{1, 2, 3, 4}
			expectedOutput := float64(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []uint{}
			expectedOutput := uint(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []uint(nil)
			expectedOutput := uint(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []uint{1, 2, 3, 4}
			expectedOutput := uint(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint8 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []uint8{}
			expectedOutput := uint8(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []uint8(nil)
			expectedOutput := uint8(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []uint8{1, 2, 3, 4}
			expectedOutput := uint8(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint16 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []uint16{}
			expectedOutput := uint16(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []uint16(nil)
			expectedOutput := uint16(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []uint16{1, 2, 3, 4}
			expectedOutput := uint16(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint32 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []uint32{}
			expectedOutput := uint32(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []uint32(nil)
			expectedOutput := uint32(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []uint32{1, 2, 3, 4}
			expectedOutput := uint32(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})

	t.Run("uint64 scenarios", func(t *testing.T) {
		t.Run("when slice is empty", func(t *testing.T) {
			input := []uint64{}
			expectedOutput := uint64(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is nil", func(t *testing.T) {
			input := []uint64(nil)
			expectedOutput := uint64(0)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			input := []uint64{1, 2, 3, 4}
			expectedOutput := uint64(result)
			actualOutput := Max(input)

			if actualOutput != expectedOutput {
				t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
			}
		})
	})
}

func TestFilter(t *testing.T) {
	t.Run("struct scenarios", func(t *testing.T) {
		type testCase struct {
			name    string
			surname string
			age     int
		}

		t.Run("when empty slice", func(t *testing.T) {
			filtered := Filter([]testCase{}, func(tc testCase) bool {
				return tc.age > 18
			})

			if len(filtered) != 0 {
				t.Errorf("Expected empty slice, but got %v", filtered)
			}
		})

		t.Run("when nil slice", func(t *testing.T) {
			filtered := Filter([]testCase(nil), func(tc testCase) bool {
				return tc.age > 18
			})

			if len(filtered) != 0 {
				t.Errorf("Expected empty slice, but got %v", filtered)
			}
		})

		t.Run("when nil slice", func(t *testing.T) {
			filtered := Filter(nil, func(tc testCase) bool {
				return tc.age > 18
			})

			if len(filtered) != 0 {
				t.Errorf("Expected empty slice, but got %v", filtered)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			testCases := []testCase{
				{
					name:    "Aksel",
					surname: "Arzuman",
					age:     30,
				},
				{
					name:    "Zeynep",
					surname: "Arzuman",
					age:     25,
				},
			}

			filtered := Filter(testCases, func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if len(filtered) != 2 {
				t.Errorf("Expected 2, but got %d", len(filtered))
			}

			if filtered[0].name != "Aksel" {
				t.Errorf("Expected Aksel, but got %s", filtered[0].name)
			}

			if filtered[1].name != "Zeynep" {
				t.Errorf("Expected Zeynep, but got %s", filtered[1].name)
			}
		})
	})
}

func TestDiff(t *testing.T) {
	t.Run("struct scenarios", func(t *testing.T) {
		type testCase struct {
			name    string
			surname string
			age     int
		}

		t.Run("when empty slice", func(t *testing.T) {
			diff := Diff([]testCase{}, []testCase{})

			if len(diff) != 0 {
				t.Errorf("Expected empty slice, but got %v", diff)
			}
		})

		t.Run("when nil slice", func(t *testing.T) {
			diff := Diff([]testCase(nil), []testCase{})

			if len(diff) != 0 {
				t.Errorf("Expected empty slice, but got %v", diff)
			}
		})

		t.Run("when nil slice", func(t *testing.T) {
			diff := Diff(nil, []testCase{})

			if len(diff) != 0 {
				t.Errorf("Expected empty slice, but got %v", diff)
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			testCases := []testCase{
				{
					name:    "Aksel",
					surname: "Arzuman",
					age:     30,
				},
				{
					name:    "Zeynep",
					surname: "Arzuman",
					age:     25,
				},
			}

			diff := Diff(testCases, []testCase{})

			if len(diff) != 2 {
				t.Errorf("Expected 2, but got %d", len(diff))
			}

			if diff[0] != testCases[0] {
				t.Errorf("Expected %v, but got %v", testCases[0], diff[0])
			}

			if diff[1] != testCases[1] {
				t.Errorf("Expected %v, but got %v", testCases[1], diff[1])
			}

			diff = Diff([]testCase{}, testCases)

			if len(diff) != 0 {
				t.Errorf("Expected empty, but got %d", len(diff))
			}
		})

		t.Run("when slices are not empty", func(t *testing.T) {
			s1 := []testCase{
				{
					name:    "Aksel",
					surname: "Arzuman",
					age:     30,
				},
				{
					name:    "Zeynep",
					surname: "Arzuman",
					age:     25,
				},
				{
					name:    "John",
					surname: "Doe",
					age:     25,
				},
				{
					name:    "Jane",
					surname: "Doe",
					age:     25,
				},
			}

			s2 := []testCase{
				{
					name:    "Aksel",
					surname: "Arzuman",
					age:     30,
				},
				{
					name:    "Zeynep",
					surname: "Arzuman",
					age:     25,
				},
			}

			diff := Diff(s1, s2)

			if len(diff) != 2 {
				t.Errorf("Expected 2, but got %d", len(diff))
			}

			if diff[0] != s1[2] {
				t.Errorf("Expected %v, but got %v", s1[2], diff[0])
			}

			if diff[1] != s1[3] {
				t.Errorf("Expected %v, but got %v", s1[3], diff[1])
			}

			diff = Diff(s2, s1)

			if len(diff) != 0 {
				t.Errorf("Expected empty, but got %d", len(diff))
			}
		})
	})
}

func TestAny(t *testing.T) {
	t.Run("struct scenarios", func(t *testing.T) {
		type testCase struct {
			name    string
			surname string
			age     int
		}

		t.Run("when empty slice", func(t *testing.T) {
			any := Any([]testCase{}, func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if any {
				t.Errorf("Expected false, but got true")
			}
		})

		t.Run("when nil slice", func(t *testing.T) {
			any := Any([]testCase(nil), func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if any {
				t.Errorf("Expected false, but got true")
			}
		})

		t.Run("when nil slice", func(t *testing.T) {
			any := Any(nil, func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if any {
				t.Errorf("Expected false, but got true")
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			testCases := []testCase{
				{
					name:    "Aksel",
					surname: "Arzuman",
					age:     30,
				},
				{
					name:    "Zeynep",
					surname: "Arzuman",
					age:     25,
				},
			}

			any := Any(testCases, func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if !any {
				t.Errorf("Expected true, but got false")
			}
		})
	})
}

func TestAll(t *testing.T) {
	t.Run("struct scenarios", func(t *testing.T) {
		type testCase struct {
			name    string
			surname string
			age     int
		}

		t.Run("when empty slice", func(t *testing.T) {
			all := All([]testCase{}, func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if !all {
				t.Errorf("Expected false, but got true")
			}
		})

		t.Run("when nil slice", func(t *testing.T) {
			all := All([]testCase(nil), func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if !all {
				t.Errorf("Expected false, but got true")
			}
		})

		t.Run("when nil slice", func(t *testing.T) {
			all := All(nil, func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if !all {
				t.Errorf("Expected false, but got true")
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			testCases := []testCase{
				{
					name:    "Aksel",
					surname: "Arzuman",
					age:     30,
				},
				{
					name:    "Zeynep",
					surname: "Arzuman",
					age:     25,
				},
			}

			all := All(testCases, func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if !all {
				t.Errorf("Expected true, but got false")
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			testCases := []testCase{
				{
					name:    "Aksel",
					surname: "Arzuman",
					age:     30,
				},
				{
					name:    "Zeynep",
					surname: "Arzuman",
					age:     25,
				},
			}

			all := All(testCases, func(t testCase) bool {
				return t.name == "Aksel"
			})

			if all {
				t.Errorf("Expected false, but got true")
			}
		})
	})
}

func TestSelect(t *testing.T) {
	t.Run("struct scenarios", func(t *testing.T) {
		type testCase struct {
			name    string
			surname string
			age     int
		}

		t.Run("when empty slice", func(t *testing.T) {
			selected := Select([]testCase{}, func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if len(selected) != 0 {
				t.Errorf("Expected empty, but got %d", len(selected))
			}
		})

		t.Run("when nil slice", func(t *testing.T) {
			selected := Select([]testCase(nil), func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if len(selected) != 0 {
				t.Errorf("Expected empty, but got %d", len(selected))
			}
		})

		t.Run("when nil slice", func(t *testing.T) {
			selected := Select(nil, func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if len(selected) != 0 {
				t.Errorf("Expected empty, but got %d", len(selected))
			}
		})

		t.Run("when slice is not empty", func(t *testing.T) {
			testCases := []testCase{
				{
					name:    "Aksel",
					surname: "Arzuman",
					age:     30,
				},
				{
					name:    "Zeynep",
					surname: "Arzuman",
					age:     25,
				},
			}

			selected := Select(testCases, func(t testCase) bool {
				return t.surname == "Arzuman"
			})

			if len(selected) != 2 {
				t.Errorf("Expected 2, but got %d", len(selected))
			}
		})
	})
}
