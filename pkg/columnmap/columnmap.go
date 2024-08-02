package columnmap

import (
	"github.com/niktheblak/ruuvitag-common/pkg/sensor"
)

// Collect calls a function on each sensor data field, providing the preferred column name for the field.
func Collect(columns map[string]string, data sensor.Data, f func(column string, v any)) {
	for _, c := range sensor.DefaultColumns {
		cn, ok := columns[c]
		if !ok {
			continue
		}
		switch c {
		case "time":
			f(cn, data.Timestamp)
		case "mac":
			f(cn, data.Addr)
		case "name":
			f(cn, data.Name)
		case "temperature":
			f(cn, data.Temperature)
		case "humidity":
			f(cn, data.Humidity)
		case "pressure":
			f(cn, data.Pressure)
		case "acceleration_x":
			f(cn, data.AccelerationX)
		case "acceleration_y":
			f(cn, data.AccelerationY)
		case "acceleration_z":
			f(cn, data.AccelerationZ)
		case "movement_counter":
			f(cn, data.MovementCounter)
		case "measurement_number":
			f(cn, data.MeasurementNumber)
		case "dew_point":
			f(cn, data.DewPoint)
		case "battery_voltage":
			f(cn, data.BatteryVoltage)
		case "tx_power":
			f(cn, data.TxPower)
		}
	}
}

// CollectFields calls a function on each sensor data field that has a non-nil value.
func CollectFields(columns map[string]string, fields sensor.Fields, f func(column string, v any)) {
	for _, c := range sensor.DefaultColumns {
		cn, ok := columns[c]
		if !ok {
			continue
		}
		switch c {
		case "time":
			f(cn, fields.Timestamp)
		case "mac":
			if fields.Addr != nil {
				f(cn, *fields.Addr)
			}
		case "name":
			if fields.Name != nil {
				f(cn, *fields.Name)
			}
		case "temperature":
			if fields.Temperature != nil {
				f(cn, *fields.Temperature)
			}
		case "humidity":
			if fields.Humidity != nil {
				f(cn, *fields.Humidity)
			}
		case "pressure":
			if fields.Pressure != nil {
				f(cn, *fields.Pressure)
			}
		case "acceleration_x":
			if fields.AccelerationX != nil {
				f(cn, *fields.AccelerationX)
			}
		case "acceleration_y":
			if fields.AccelerationY != nil {
				f(cn, *fields.AccelerationY)
			}
		case "acceleration_z":
			if fields.AccelerationZ != nil {
				f(cn, *fields.AccelerationZ)
			}
		case "movement_counter":
			if fields.MovementCounter != nil {
				f(cn, *fields.MovementCounter)
			}
		case "measurement_number":
			if fields.MeasurementNumber != nil {
				f(cn, *fields.MeasurementNumber)
			}
		case "dew_point":
			if fields.DewPoint != nil {
				f(cn, *fields.DewPoint)
			}
		case "battery_voltage":
			if fields.BatteryVoltage != nil {
				f(cn, *fields.BatteryVoltage)
			}
		case "tx_power":
			if fields.TxPower != nil {
				f(cn, *fields.TxPower)
			}
		}
	}
}

// Transform creates a map with the values copied from the given sensor.Data, using the provided column names as map keys.
//
// Any values that are not included in the given column map are not included in the returned map.
func Transform(columns map[string]string, data sensor.Data) map[string]any {
	values := make(map[string]any)
	Collect(columns, data, func(column string, v any) {
		values[column] = v
	})
	return values
}

// TransformFields creates a map with the values copied from the given sensor.Fields, using the provided column names as map keys.
//
// Any fields that have a nil value or are not included in the given column map are not included in the returned map.
func TransformFields(columns map[string]string, fields sensor.Fields) map[string]any {
	values := make(map[string]any)
	CollectFields(columns, fields, func(column string, v any) {
		values[column] = v
	})
	return values
}
