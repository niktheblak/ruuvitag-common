package sensor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateColumnMapping(t *testing.T) {
	err := ValidateColumnMapping(nil)
	assert.Error(t, err)
	err = ValidateColumnMapping(map[string]string{
		"time":        "ts",
		"name":        "name",
		"temperature": "temperature",
		"humidity":    "humidity",
		"pressure":    "pressure",
	})
	assert.NoError(t, err)
	err = ValidateColumnMapping(map[string]string{
		"time":        "ts",
		"mac":         "addr",
		"temperature": "temperature",
		"humidity":    "humidity",
		"pressure":    "pressure",
	})
	assert.NoError(t, err)
	err = ValidateColumnMapping(map[string]string{
		"time":        "ts",
		"temperature": "temperature",
		"humidity":    "humidity",
		"pressure":    "pressure",
	})
	assert.ErrorIs(t, err, ErrMissingColumn)
	err = ValidateColumnMapping(map[string]string{
		"name":        "name",
		"mac":         "mac",
		"temperature": "temperature",
		"humidity":    "humidity",
		"pressure":    "pressure",
	})
	assert.ErrorIs(t, err, ErrMissingColumn)
}

func TestValidateRequestedColumns(t *testing.T) {
	columns := map[string]string{
		"time":        "ts",
		"name":        "name",
		"temperature": "temperature",
		"humidity":    "humidity",
		"pressure":    "pressure",
	}
	err := ValidateRequestedColumns(columns, []string{})
	assert.ErrorIs(t, err, ErrMissingColumn)
	err = ValidateRequestedColumns(columns, []string{"ts", "temperature"})
	assert.ErrorIs(t, err, ErrMissingColumn)
	err = ValidateRequestedColumns(columns, []string{"ts", "name", "temperature", "oxidization"})
	assert.ErrorIs(t, err, ErrInvalidColumn)
	err = ValidateRequestedColumns(columns, []string{"ts", "name", "temperature"})
	assert.NoError(t, err)
}
