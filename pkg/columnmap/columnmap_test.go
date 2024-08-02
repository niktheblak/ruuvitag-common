package columnmap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/niktheblak/ruuvitag-common/pkg/sensor"
)

var testData = sensor.Data{
	Addr:        "ec-40-93-94-35-2a",
	Name:        "Test",
	Temperature: 22.5,
	Humidity:    46,
	DewPoint:    12.1,
	Pressure:    1002,
	Timestamp:   time.Date(2024, time.July, 3, 10, 24, 23, 213, time.UTC),
}

func TestCollect(t *testing.T) {
	columns := map[string]string{
		"time":        "ts",
		"name":        "name",
		"mac":         "addr",
		"temperature": "temperature",
		"humidity":    "humidity",
		"pressure":    "pressure",
		"dew_point":   "dew_point",
	}
	values := make(map[string]any)
	Collect(columns, testData, func(column string, v any) {
		values[column] = v
	})
	assert.Equal(t, testData.Timestamp, values["ts"])
	assert.Equal(t, testData.Name, values["name"])
	assert.Equal(t, testData.Addr, values["addr"])
	assert.Equal(t, testData.Temperature, values["temperature"])
	assert.Equal(t, testData.Humidity, values["humidity"])
	assert.Equal(t, testData.Pressure, values["pressure"])
	assert.Equal(t, testData.DewPoint, values["dew_point"])
}

func TestCollectFields(t *testing.T) {
	columns := map[string]string{
		"time":        "ts",
		"name":        "name",
		"mac":         "addr",
		"temperature": "temperature",
		"humidity":    "humidity",
		"pressure":    "pressure",
		"dew_point":   "dew_point",
	}
	fields := sensor.FromData(testData)
	values := make(map[string]any)
	CollectFields(columns, fields, func(column string, v any) {
		values[column] = v
	})
	assert.Equal(t, testData.Timestamp, values["ts"])
	assert.Equal(t, testData.Name, values["name"])
	assert.Equal(t, testData.Addr, values["addr"])
	assert.Equal(t, testData.Temperature, values["temperature"])
	assert.Equal(t, testData.Humidity, values["humidity"])
	assert.Equal(t, testData.Pressure, values["pressure"])
	assert.Equal(t, testData.DewPoint, values["dew_point"])
}
