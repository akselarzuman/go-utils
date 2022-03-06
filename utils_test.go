package utils

import "testing"

func TestSum_Success_Cases(t *testing.T) {
	result := 10

	t.Run("sums slice of int", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		expectedOutput := result
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of int8", func(t *testing.T) {
		input := []int8{1, 2, 3, 4}
		expectedOutput := int8(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of int16", func(t *testing.T) {
		input := []int16{1, 2, 3, 4}
		expectedOutput := int16(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of int32", func(t *testing.T) {
		input := []int32{1, 2, 3, 4}
		expectedOutput := int32(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of int64", func(t *testing.T) {
		input := []int64{1, 2, 3, 4}
		expectedOutput := int64(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of float32", func(t *testing.T) {
		input := []float32{1, 2, 3, 4}
		expectedOutput := float32(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of float64", func(t *testing.T) {
		input := []float64{1, 2, 3, 4}
		expectedOutput := float64(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %f, but got %f", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of uint", func(t *testing.T) {
		input := []uint{1, 2, 3, 4}
		expectedOutput := uint(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of uint8", func(t *testing.T) {
		input := []uint8{1, 2, 3, 4}
		expectedOutput := uint8(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of uint16", func(t *testing.T) {
		input := []uint16{1, 2, 3, 4}
		expectedOutput := uint16(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of uint32", func(t *testing.T) {
		input := []uint32{1, 2, 3, 4}
		expectedOutput := uint32(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})

	t.Run("sums slice of uint64", func(t *testing.T) {
		input := []uint64{1, 2, 3, 4}
		expectedOutput := uint64(result)
		actualOutput := Sum(input)

		if actualOutput != expectedOutput {
			t.Errorf("Expected %d, but got %d", expectedOutput, actualOutput)
		}
	})
}

func TestSum_Fail_Cases(t *testing.T) {
	unexpectedResult := 11

	t.Run("sums slice of int", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		unexpectedOutput := unexpectedResult
		actualOutput := Sum(input)

		if actualOutput == unexpectedOutput {
			t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
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

	t.Run("sums slice of int16", func(t *testing.T) {
		input := []int16{1, 2, 3, 4}
		unexpectedOutput := int16(unexpectedResult)
		actualOutput := Sum(input)

		if actualOutput == unexpectedOutput {
			t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
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

	t.Run("sums slice of int64", func(t *testing.T) {
		input := []int64{1, 2, 3, 4}
		unexpectedOutput := int64(unexpectedResult)
		actualOutput := Sum(input)

		if actualOutput == unexpectedOutput {
			t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
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

	t.Run("sums slice of float64", func(t *testing.T) {
		input := []float64{1, 2, 3, 4}
		unexpectedOutput := float64(unexpectedResult)
		actualOutput := Sum(input)

		if actualOutput == unexpectedOutput {
			t.Errorf("Unexpected %f, got %f", unexpectedOutput, actualOutput)
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

	t.Run("sums slice of uint8", func(t *testing.T) {
		input := []uint8{1, 2, 3, 4}
		unexpectedOutput := uint8(unexpectedResult)
		actualOutput := Sum(input)

		if actualOutput == unexpectedOutput {
			t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
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

	t.Run("sums slice of uint32", func(t *testing.T) {
		input := []uint32{1, 2, 3, 4}
		unexpectedOutput := uint32(unexpectedResult)
		actualOutput := Sum(input)

		if actualOutput == unexpectedOutput {
			t.Errorf("Unexpected %d, got %d", unexpectedOutput, actualOutput)
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
}
