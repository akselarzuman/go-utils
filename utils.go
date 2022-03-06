package utils

// Contains checks whether the value is in slice or not.
// It supports int, int8, int16, int32, int64, float32, float64, uint, uint8, uint16, uint32, uint64 and string
// as types for slice values.
func Contains[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64 | string](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

// Sum sums the values of slice.
// It supports int, int8, int16, int32, int64, float32, float64, uint, uint8, uint16, uint32 and uint64
// as types for slice values.
func Sum[T int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}