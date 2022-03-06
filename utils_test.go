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
}

func TestMin(t *testing.T) {
	result := 1

	t.Run("int scenario", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		expectedOutput := result
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("int8 scenario", func(t *testing.T) {
		input := []int8{1, 2, 3, 4}
		expectedOutput := int8(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("int16 scenario", func(t *testing.T) {
		input := []int16{1, 2, 3, 4}
		expectedOutput := int16(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("int32 scenario", func(t *testing.T) {
		input := []int32{1, 2, 3, 4}
		expectedOutput := int32(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("int64 scenario", func(t *testing.T) {
		input := []int64{1, 2, 3, 4}
		expectedOutput := int64(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("float32 scenario", func(t *testing.T) {
		input := []float32{1, 2, 3, 4}
		expectedOutput := float32(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
		}
	})

	t.Run("float64 scenario", func(t *testing.T) {
		input := []float64{1, 2, 3, 4}
		expectedOutput := float64(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
		}
	})

	t.Run("uint scenario", func(t *testing.T) {
		input := []uint{1, 2, 3, 4}
		expectedOutput := uint(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("uint8 scenario", func(t *testing.T) {
		input := []uint8{1, 2, 3, 4}
		expectedOutput := uint8(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("uint16 scenario", func(t *testing.T) {
		input := []uint16{1, 2, 3, 4}
		expectedOutput := uint16(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("uint32 scenario", func(t *testing.T) {
		input := []uint32{1, 2, 3, 4}
		expectedOutput := uint32(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("uint64 scenario", func(t *testing.T) {
		input := []uint64{1, 2, 3, 4}
		expectedOutput := uint64(result)
		actualOutput := Min(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})
}
