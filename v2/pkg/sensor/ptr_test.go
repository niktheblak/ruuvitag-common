package sensor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFloat64Pointer(t *testing.T) {
	ptr := Float64Pointer(21.5)
	require.NotNil(t, ptr)
	assert.Equal(t, 21.5, *ptr)
}

func TestZeroFloat64Pointer(t *testing.T) {
	ptr := ZeroFloat64Pointer()
	require.NotNil(t, ptr)
	assert.Equal(t, 0.0, *ptr)
}

func TestIntPointer(t *testing.T) {
	ptr := IntPointer(21)
	require.NotNil(t, ptr)
	assert.Equal(t, 21, *ptr)
}

func TestZeroIntPointer(t *testing.T) {
	ptr := ZeroIntPointer()
	require.NotNil(t, ptr)
	assert.Equal(t, 0, *ptr)
}

func TestStringPointer(t *testing.T) {
	ptr := StringPointer("hello")
	require.NotNil(t, ptr)
	assert.Equal(t, "hello", *ptr)
}

func TestZeroStringPointer(t *testing.T) {
	ptr := ZeroStringPointer()
	require.NotNil(t, ptr)
	assert.Equal(t, "", *ptr)
}
