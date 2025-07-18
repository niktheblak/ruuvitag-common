package sensor

import (
	"errors"
	"fmt"
	"maps"
	"slices"
)

// DefaultColumns names of fields contained in a RuuviTag v5 protocol transmission.
var DefaultColumns = []string{
	"time",
	"mac",
	"name",
	"temperature",
	"humidity",
	"pressure",
	"acceleration_x",
	"acceleration_y",
	"acceleration_z",
	"movement_counter",
	"measurement_number",
	"dew_point",
	"battery_voltage",
	"tx_power",
	"wet_bulb",
}

var (
	ErrInvalidColumn = errors.New("invalid column")
	ErrMissingColumn = errors.New("missing mandatory column")
)

// DefaultColumnMap column mapping where each column has the default name.
var DefaultColumnMap map[string]string

func init() {
	DefaultColumnMap = make(map[string]string)
	for _, column := range DefaultColumns {
		DefaultColumnMap[column] = column
	}
}

// ValidateColumnMapping validates that a given custom column mapping contains necessary fields to look up RuuviTag data.
func ValidateColumnMapping(columns map[string]string) error {
	if len(columns) == 0 {
		return fmt.Errorf("columns cannot be empty")
	}
	_, ok := columns["time"]
	if !ok {
		return fmt.Errorf("%w: time", ErrMissingColumn)
	}
	_, nameOK := columns["name"]
	_, macOK := columns["mac"]
	if !nameOK && !macOK {
		return fmt.Errorf("%w: name or mac", ErrMissingColumn)
	}
	for cn := range columns {
		_, ok := DefaultColumnMap[cn]
		if !ok {
			return fmt.Errorf("%w: %s", ErrInvalidColumn, cn)
		}
	}
	return nil
}

// ValidateRequestedColumns validates that given requested columns conform to given custom column mapping.
//
// The custom column mapping is expected to be validated with ValidateColumnMapping prior to calling this function.
func ValidateRequestedColumns(columns map[string]string, requested []string) error {
	timeOK := false
	nameOK := false
	macOK := false
	columnNames := slices.Collect(maps.Values(columns))
	for _, c := range requested {
		if !slices.Contains(columnNames, c) {
			return fmt.Errorf("%w: %s", ErrInvalidColumn, c)
		}
		if c == columns["time"] {
			timeOK = true
		}
		if c == columns["name"] {
			nameOK = true
		}
		if c == columns["mac"] {
			macOK = true
		}
	}
	if !timeOK {
		return fmt.Errorf("%w: %s", ErrMissingColumn, columns["time"])
	}
	if !nameOK && !macOK {
		return fmt.Errorf("%w: %s or %s", ErrMissingColumn, columns["name"], columns["mac"])
	}
	return nil
}
