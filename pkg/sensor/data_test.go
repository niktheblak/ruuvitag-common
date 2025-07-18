package sensor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testFields = Fields{
	Addr:        StringPointer("ec-40-93-94-35-2a"),
	Name:        StringPointer("Test"),
	Temperature: Float64Pointer(22.5),
	Humidity:    Float64Pointer(46),
	DewPoint:    Float64Pointer(12.1),
	WetBulb:     Float64Pointer(21.5),
	Pressure:    Float64Pointer(1002),
	Timestamp:   time.Date(2024, time.July, 3, 10, 24, 23, 213, time.UTC),
}

func TestFromFields(t *testing.T) {
	data := FromFields(testFields)
	assert.Equal(t, testData.Timestamp, data.Timestamp)
	assert.Equal(t, testData.Addr, data.Addr)
	assert.Equal(t, testData.Name, data.Name)
	assert.Equal(t, testData.Temperature, data.Temperature)
	assert.Equal(t, testData.Humidity, data.Humidity)
	assert.Equal(t, testData.Pressure, data.Pressure)
	assert.Equal(t, testData.DewPoint, data.DewPoint)
	assert.Equal(t, testData.WetBulb, data.WetBulb)
	assert.Equal(t, 0.0, data.BatteryVoltage)
	assert.Equal(t, 0, data.TxPower)
	assert.Equal(t, 0, data.MeasurementNumber)
	assert.Equal(t, 0, data.MovementCounter)
	assert.Equal(t, 0, data.AccelerationX)
	assert.Equal(t, 0, data.AccelerationY)
	assert.Equal(t, 0, data.AccelerationZ)
}
