package sensor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = Data{
	Addr:        "ec-40-93-94-35-2a",
	Name:        "Test",
	Temperature: 22.5,
	Humidity:    46,
	DewPoint:    12.1,
	Pressure:    1002,
	Timestamp:   time.Date(2024, time.July, 3, 10, 24, 23, 213, time.UTC),
}

func TestAllZeroFields(t *testing.T) {
	fields := AllZeroFields()
	require.NotNil(t, fields.Addr)
	require.NotNil(t, fields.Name)
	require.NotNil(t, fields.Temperature)
	require.NotNil(t, fields.Humidity)
	require.NotNil(t, fields.DewPoint)
	require.NotNil(t, fields.Pressure)
	require.NotNil(t, fields.BatteryVoltage)
	require.NotNil(t, fields.TxPower)
	require.NotNil(t, fields.MeasurementNumber)
	require.NotNil(t, fields.MovementCounter)
	require.NotNil(t, fields.AccelerationX)
	require.NotNil(t, fields.AccelerationY)
	require.NotNil(t, fields.AccelerationZ)
	assert.True(t, fields.Timestamp.IsZero())
	assert.Equal(t, "", *fields.Addr)
	assert.Equal(t, "", *fields.Name)
	assert.Equal(t, 0.0, *fields.Temperature)
	assert.Equal(t, 0.0, *fields.Humidity)
	assert.Equal(t, 0.0, *fields.Pressure)
	assert.Equal(t, 0.0, *fields.DewPoint)
	assert.Equal(t, 0.0, *fields.BatteryVoltage)
	assert.Equal(t, 0, *fields.TxPower)
	assert.Equal(t, 0, *fields.MeasurementNumber)
	assert.Equal(t, 0, *fields.MovementCounter)
	assert.Equal(t, 0, *fields.AccelerationX)
	assert.Equal(t, 0, *fields.AccelerationY)
	assert.Equal(t, 0, *fields.AccelerationZ)
}

func TestFromData(t *testing.T) {
	fields := FromData(testData)
	require.NotNil(t, fields.Addr)
	require.NotNil(t, fields.Name)
	require.NotNil(t, fields.Temperature)
	require.NotNil(t, fields.Humidity)
	require.NotNil(t, fields.DewPoint)
	require.NotNil(t, fields.Pressure)
	require.NotNil(t, fields.BatteryVoltage)
	require.NotNil(t, fields.TxPower)
	require.NotNil(t, fields.MeasurementNumber)
	require.NotNil(t, fields.MovementCounter)
	require.NotNil(t, fields.AccelerationX)
	require.NotNil(t, fields.AccelerationY)
	require.NotNil(t, fields.AccelerationZ)
	assert.Equal(t, testData.Timestamp, fields.Timestamp)
	assert.Equal(t, testData.Addr, *fields.Addr)
	assert.Equal(t, testData.Name, *fields.Name)
	assert.Equal(t, testData.Temperature, *fields.Temperature)
	assert.Equal(t, testData.Humidity, *fields.Humidity)
	assert.Equal(t, testData.Pressure, *fields.Pressure)
	assert.Equal(t, testData.DewPoint, *fields.DewPoint)
	assert.Equal(t, 0.0, *fields.BatteryVoltage)
	assert.Equal(t, 0, *fields.TxPower)
	assert.Equal(t, 0, *fields.MeasurementNumber)
	assert.Equal(t, 0, *fields.MovementCounter)
	assert.Equal(t, 0, *fields.AccelerationX)
	assert.Equal(t, 0, *fields.AccelerationY)
	assert.Equal(t, 0, *fields.AccelerationZ)
}

func TestNonZeroFields(t *testing.T) {
	fields := NonZeroFields(testData)
	require.NotNil(t, fields.Addr)
	require.NotNil(t, fields.Name)
	require.NotNil(t, fields.Temperature)
	require.NotNil(t, fields.Humidity)
	require.NotNil(t, fields.DewPoint)
	require.NotNil(t, fields.Pressure)
	assert.Equal(t, testData.Timestamp, fields.Timestamp)
	assert.Equal(t, testData.Addr, *fields.Addr)
	assert.Equal(t, testData.Name, *fields.Name)
	assert.Equal(t, testData.Temperature, *fields.Temperature)
	assert.Equal(t, testData.Humidity, *fields.Humidity)
	assert.Equal(t, testData.Pressure, *fields.Pressure)
	assert.Equal(t, testData.DewPoint, *fields.DewPoint)
	assert.Nil(t, fields.BatteryVoltage)
	assert.Nil(t, fields.TxPower)
	assert.Nil(t, fields.MeasurementNumber)
	assert.Nil(t, fields.MovementCounter)
	assert.Nil(t, fields.AccelerationX)
	assert.Nil(t, fields.AccelerationY)
	assert.Nil(t, fields.AccelerationZ)
}
