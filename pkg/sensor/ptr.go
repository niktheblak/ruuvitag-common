package sensor

// Float64Pointer returns a pointer to the given float64.
func Float64Pointer(v float64) *float64 {
	return &v
}

// ZeroFloat64Pointer returns a pointer to a float64 with zero value.
func ZeroFloat64Pointer() *float64 {
	var v float64
	return &v
}

// IntPointer returns a pointer to the given int.
func IntPointer(v int) *int {
	return &v
}

// ZeroIntPointer returns a pointer to an int with zero value.
func ZeroIntPointer() *int {
	var v int
	return &v
}

// StringPointer returns a pointer to the given string.
func StringPointer(v string) *string {
	return &v
}

// ZeroStringPointer returns a *string pointer to an empty string.
func ZeroStringPointer() *string {
	var v string
	return &v
}
