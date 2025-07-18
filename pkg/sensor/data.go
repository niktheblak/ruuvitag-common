package sensor

import (
	"time"
)

// Data contains all fields of a RuuviTag v5 measurement.
type Data struct {
	Timestamp         time.Time `json:"time"`
	Addr              string    `json:"mac"`
	Name              string    `json:"name"`
	Temperature       float64   `json:"temperature"`
	Humidity          float64   `json:"humidity"`
	DewPoint          float64   `json:"dew_point"`
	WetBulb           float64   `json:"wet_bulb"`
	Pressure          float64   `json:"pressure"`
	BatteryVoltage    float64   `json:"battery_voltage"`
	TxPower           int       `json:"tx_power"`
	AccelerationX     int       `json:"acceleration_x"`
	AccelerationY     int       `json:"acceleration_y"`
	AccelerationZ     int       `json:"acceleration_z"`
	MovementCounter   int       `json:"movement_counter"`
	MeasurementNumber int       `json:"measurement_number"`
}

// FromFields creates a new Data struct by copying all non-zero pointer values from the given Fields struct.
func FromFields(f Fields) Data {
	var d Data
	d.Timestamp = f.Timestamp
	if f.Addr != nil {
		d.Addr = *f.Addr
	}
	if f.Name != nil {
		d.Name = *f.Name
	}
	if f.Temperature != nil {
		d.Temperature = *f.Temperature
	}
	if f.Humidity != nil {
		d.Humidity = *f.Humidity
	}
	if f.DewPoint != nil {
		d.DewPoint = *f.DewPoint
	}
	if f.WetBulb != nil {
		d.WetBulb = *f.WetBulb
	}
	if f.Pressure != nil {
		d.Pressure = *f.Pressure
	}
	if f.BatteryVoltage != nil {
		d.BatteryVoltage = *f.BatteryVoltage
	}
	if f.TxPower != nil {
		d.TxPower = *f.TxPower
	}
	if f.AccelerationX != nil {
		d.AccelerationX = *f.AccelerationX
	}
	if f.AccelerationY != nil {
		d.AccelerationY = *f.AccelerationY
	}
	if f.AccelerationZ != nil {
		d.AccelerationZ = *f.AccelerationZ
	}
	if f.MeasurementNumber != nil {
		d.MeasurementNumber = *f.MeasurementNumber
	}
	if f.MovementCounter != nil {
		d.MovementCounter = *f.MovementCounter
	}
	return d
}
