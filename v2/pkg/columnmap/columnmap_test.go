package columnmap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/niktheblak/ruuvitag-common/v2/pkg/sensor"
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

var testColumns = map[string]string{
	"time":        "ts",
	"name":        "name",
	"mac":         "addr",
	"temperature": "temperature",
	"humidity":    "humidity",
	"pressure":    "pressure",
	"dew_point":   "dew_point",
}

func TestCollect(t *testing.T) {
	values := make(map[string]any)
	for column, v := range Collect(testColumns, testData) {
		values[column] = v
	}
	assert.Equal(t, testData.Timestamp, values["ts"])
	assert.Equal(t, testData.Name, values["name"])
	assert.Equal(t, testData.Addr, values["addr"])
	assert.Equal(t, testData.Temperature, values["temperature"])
	assert.Equal(t, testData.Humidity, values["humidity"])
	assert.Equal(t, testData.Pressure, values["pressure"])
	assert.Equal(t, testData.DewPoint, values["dew_point"])
}

func TestCollectFields(t *testing.T) {
	fields := sensor.FromData(testData)
	values := make(map[string]any)
	for column, v := range CollectFields(testColumns, fields) {
		values[column] = v
	}
	assert.Equal(t, testData.Timestamp, values["ts"])
	assert.Equal(t, testData.Name, values["name"])
	assert.Equal(t, testData.Addr, values["addr"])
	assert.Equal(t, testData.Temperature, values["temperature"])
	assert.Equal(t, testData.Humidity, values["humidity"])
	assert.Equal(t, testData.Pressure, values["pressure"])
	assert.Equal(t, testData.DewPoint, values["dew_point"])
}

func TestTransform(t *testing.T) {
	data := testData
	data.MeasurementNumber = 102
	values := Transform(testColumns, testData)
	assert.Equal(t, testData.Timestamp, values["ts"])
	assert.Equal(t, testData.Name, values["name"])
	assert.Equal(t, testData.Addr, values["addr"])
	assert.Equal(t, testData.Temperature, values["temperature"])
	assert.Equal(t, testData.Humidity, values["humidity"])
	assert.Equal(t, testData.Pressure, values["pressure"])
	assert.Equal(t, testData.DewPoint, values["dew_point"])
	assert.NotContains(t, values, "acceleration_x")
	assert.NotContains(t, values, "measurement_number")
}

func TestTransformFields(t *testing.T) {
	fields := sensor.FromData(testData)
	fields.MeasurementNumber = sensor.IntPointer(102)
	values := TransformFields(testColumns, fields)
	assert.Equal(t, testData.Timestamp, values["ts"])
	assert.Equal(t, testData.Name, values["name"])
	assert.Equal(t, testData.Addr, values["addr"])
	assert.Equal(t, testData.Temperature, values["temperature"])
	assert.Equal(t, testData.Humidity, values["humidity"])
	assert.Equal(t, testData.Pressure, values["pressure"])
	assert.Equal(t, testData.DewPoint, values["dew_point"])
	assert.NotContains(t, values, "acceleration_x")
	assert.NotContains(t, values, "measurement_number")
}
