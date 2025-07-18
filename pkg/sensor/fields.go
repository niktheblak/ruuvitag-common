package sensor

import (
	"time"
)

// Fields contains all fields of a RuuviTag v5 measurement, but only the time field is mandatory.
// To be used e.g. with REST APIs that do not need to return empty fields.
type Fields struct {
	Timestamp         time.Time `json:"time"`
	Addr              *string   `json:"mac,omitempty"`
	Name              *string   `json:"name,omitempty"`
	Temperature       *float64  `json:"temperature,omitempty"`
	Humidity          *float64  `json:"humidity,omitempty"`
	DewPoint          *float64  `json:"dew_point,omitempty"`
	WetBulb           *float64  `json:"wet_bulb,omitempty"`
	Pressure          *float64  `json:"pressure,omitempty"`
	BatteryVoltage    *float64  `json:"battery_voltage,omitempty"`
	TxPower           *int      `json:"tx_power,omitempty"`
	AccelerationX     *int      `json:"acceleration_x,omitempty"`
	AccelerationY     *int      `json:"acceleration_y,omitempty"`
	AccelerationZ     *int      `json:"acceleration_z,omitempty"`
	MovementCounter   *int      `json:"movement_counter,omitempty"`
	MeasurementNumber *int      `json:"measurement_number,omitempty"`
}

// AllZeroFields creates a Fields struct with all pointer values set to a zero number.
func AllZeroFields() Fields {
	return Fields{
		Addr:              ZeroStringPointer(),
		Name:              ZeroStringPointer(),
		Temperature:       ZeroFloat64Pointer(),
		Humidity:          ZeroFloat64Pointer(),
		DewPoint:          ZeroFloat64Pointer(),
		WetBulb:           ZeroFloat64Pointer(),
		Pressure:          ZeroFloat64Pointer(),
		BatteryVoltage:    ZeroFloat64Pointer(),
		TxPower:           ZeroIntPointer(),
		AccelerationX:     ZeroIntPointer(),
		AccelerationY:     ZeroIntPointer(),
		AccelerationZ:     ZeroIntPointer(),
		MovementCounter:   ZeroIntPointer(),
		MeasurementNumber: ZeroIntPointer(),
	}
}

// FromData copies all values from the given Data struct regardless of whether they have nonzero value or not.
func FromData(d Data) Fields {
	return Fields{
		Timestamp:         d.Timestamp,
		Addr:              &d.Addr,
		Name:              &d.Name,
		Temperature:       &d.Temperature,
		Humidity:          &d.Humidity,
		DewPoint:          &d.DewPoint,
		WetBulb:           &d.WetBulb,
		Pressure:          &d.Pressure,
		BatteryVoltage:    &d.BatteryVoltage,
		TxPower:           &d.TxPower,
		AccelerationX:     &d.AccelerationX,
		AccelerationY:     &d.AccelerationY,
		AccelerationZ:     &d.AccelerationZ,
		MovementCounter:   &d.MovementCounter,
		MeasurementNumber: &d.MeasurementNumber,
	}
}

// NonZeroFields copies fields with non-zero value from the given Data struct.
func NonZeroFields(d Data) Fields {
	var f Fields
	f.Timestamp = d.Timestamp
	if d.Addr != "" {
		f.Addr = &d.Addr
	}
	if d.Name != "" {
		f.Name = &d.Name
	}
	if d.Temperature != 0 {
		f.Temperature = &d.Temperature
	}
	if d.Humidity != 0 {
		f.Humidity = &d.Humidity
	}
	if d.DewPoint != 0 {
		f.DewPoint = &d.DewPoint
	}
	if d.WetBulb != 0 {
		f.WetBulb = &d.WetBulb
	}
	if d.Pressure != 0 {
		f.Pressure = &d.Pressure
	}
	if d.BatteryVoltage != 0 {
		f.BatteryVoltage = &d.BatteryVoltage
	}
	if d.TxPower != 0 {
		f.TxPower = &d.TxPower
	}
	if d.AccelerationX != 0 {
		f.AccelerationX = &d.AccelerationX
	}
	if d.AccelerationY != 0 {
		f.AccelerationY = &d.AccelerationY
	}
	if d.AccelerationZ != 0 {
		f.AccelerationZ = &d.AccelerationZ
	}
	if d.MeasurementNumber != 0 {
		f.MeasurementNumber = &d.MeasurementNumber
	}
	if d.MovementCounter != 0 {
		f.MovementCounter = &d.MovementCounter
	}
	return f
}

func FieldsFromColumns(d Data, columns []string) Fields {
	var f Fields
	for _, c := range columns {
		switch c {
		case "time":
			f.Timestamp = d.Timestamp
		case "mac":
			f.Addr = &d.Addr
		case "name":
			f.Name = &d.Name
		case "temperature":
			f.Temperature = &d.Temperature
		case "humidity":
			f.Humidity = &d.Humidity
		case "pressure":
			f.Pressure = &d.Pressure
		case "acceleration_x":
			f.AccelerationX = &d.AccelerationX
		case "acceleration_y":
			f.AccelerationY = &d.AccelerationY
		case "acceleration_z":
			f.AccelerationZ = &d.AccelerationZ
		case "movement_counter":
			f.MovementCounter = &d.MovementCounter
		case "measurement_number":
			f.MeasurementNumber = &d.MeasurementNumber
		case "dew_point":
			f.DewPoint = &d.DewPoint
		case "wet_bulb":
			f.WetBulb = &d.WetBulb
		case "battery_voltage":
			f.BatteryVoltage = &d.BatteryVoltage
		case "tx_power":
			f.TxPower = &d.TxPower
		}
	}
	return f
}
