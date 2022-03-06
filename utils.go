package utils

import (
	"reflect"
	"sort"
)

type number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

type sortable interface {
	number | string
}

// Contains checks whether the value is in slice or not.
// It supports int, int8, int16, int32, int64, float32, float64, uint, uint8, uint16, uint32, uint64 and string
// as types for slice values.
func Contains[T number | string](slice []T, value T) bool {
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
func Sum[T number](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Sort sorts ascending the values of slice.
// It supports int, int8, int16, int32, int64, float32, float64, uint, uint8, uint16, uint32, uint64 and string
// as types for slice values.
func Sort[T sortable](slice []T) {
	t := reflect.TypeOf(slice)

	if t.Elem().Kind() == reflect.String {
		var s interface{} = slice
		sort.Strings(s.([]string))
		return
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

// SortDesc sorts descending the values of slice.
// It supports int, int8, int16, int32, int64, float32, float64, uint, uint8, uint16, uint32, uint64 and string
// as types for slice values.
func SortDesc[T sortable](slice []T) {
	t := reflect.TypeOf(slice)

	if t.Elem().Kind() == reflect.String {
		var s interface{} = slice
		sort.Sort(sort.Reverse(sort.StringSlice(s.([]string))))
		return
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
}

// Min get the min value of slice.
// It supports int, int8, int16, int32, int64, float32, float64, uint, uint8, uint16, uint32 and uint64
// as types for slice values.
// Returns 0 if slice is nil or empty.
func Min[T number](slice []T) T {
	if len(slice) == 0 {
		return T(0)
	}

	var min T = slice[0]

	for _, v := range slice {
		if v < min {
			min = v
		}
	}

	return min
}

// Max get the max value of slice.
// It supports int, int8, int16, int32, int64, float32, float64, uint, uint8, uint16, uint32 and uint64
// as types for slice values.
// Returns 0 if slice is nil or empty.
func Max[T number](slice []T) T {
	if len(slice) == 0 {
		return T(0)
	}

	var max T = slice[0]

	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max
}
