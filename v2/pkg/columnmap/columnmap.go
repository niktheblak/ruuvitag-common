package columnmap

import (
	"iter"

	"github.com/niktheblak/ruuvitag-common/v2/pkg/sensor"
)

// Collect return an iterator on each sensor data field with field name as key.
func Collect(columns map[string]string, data sensor.Data) iter.Seq2[string, any] {
	return func(yield func(string, any) bool) {
		for _, c := range sensor.DefaultColumns {
			cn, ok := columns[c]
			if !ok {
				continue
			}
			var v any = nil
			switch c {
			case "time":
				v = data.Timestamp
			case "mac":
				v = data.Addr
			case "name":
				v = data.Name
			case "temperature":
				v = data.Temperature
			case "humidity":
				v = data.Humidity
			case "pressure":
				v = data.Pressure
			case "acceleration_x":
				v = data.AccelerationX
			case "acceleration_y":
				v = data.AccelerationY
			case "acceleration_z":
				v = data.AccelerationZ
			case "movement_counter":
				v = data.MovementCounter
			case "measurement_number":
				v = data.MeasurementNumber
			case "dew_point":
				v = data.DewPoint
			case "battery_voltage":
				v = data.BatteryVoltage
			case "tx_power":
				v = data.TxPower
			}
			if v == nil {
				continue
			}
			if !yield(cn, v) {
				return
			}
		}
	}
}

// CollectFields returns an iterator over each sensor data field that has a non-nil value with field name as key.
func CollectFields(columns map[string]string, fields sensor.Fields) iter.Seq2[string, any] {
	return func(yield func(string, any) bool) {
		for _, c := range sensor.DefaultColumns {
			cn, ok := columns[c]
			if !ok {
				continue
			}
			var v any = nil
			switch c {
			case "time":
				v = fields.Timestamp
			case "mac":
				if fields.Addr != nil {
					v = *fields.Addr
				}
			case "name":
				if fields.Name != nil {
					v = *fields.Name
				}
			case "temperature":
				if fields.Temperature != nil {
					v = *fields.Temperature
				}
			case "humidity":
				if fields.Humidity != nil {
					v = *fields.Humidity
				}
			case "pressure":
				if fields.Pressure != nil {
					v = *fields.Pressure
				}
			case "acceleration_x":
				if fields.AccelerationX != nil {
					v = *fields.AccelerationX
				}
			case "acceleration_y":
				if fields.AccelerationY != nil {
					v = *fields.AccelerationY
				}
			case "acceleration_z":
				if fields.AccelerationZ != nil {
					v = *fields.AccelerationZ
				}
			case "movement_counter":
				if fields.MovementCounter != nil {
					v = *fields.MovementCounter
				}
			case "measurement_number":
				if fields.MeasurementNumber != nil {
					v = *fields.MeasurementNumber
				}
			case "dew_point":
				if fields.DewPoint != nil {
					v = *fields.DewPoint
				}
			case "battery_voltage":
				if fields.BatteryVoltage != nil {
					v = *fields.BatteryVoltage
				}
			case "tx_power":
				if fields.TxPower != nil {
					v = *fields.TxPower
				}
			}
			if v == nil {
				continue
			}
			if !yield(cn, v) {
				return
			}
		}
	}
}

// Transform creates a map with the values copied from the given sensor.Data, using the provided column names as map keys.
//
// Any values that are not included in the given column map are not included in the returned map.
func Transform(columns map[string]string, data sensor.Data) map[string]any {
	values := make(map[string]any)
	for column, v := range Collect(columns, data) {
		values[column] = v
	}
	return values
}

// TransformFields creates a map with the values copied from the given sensor.Fields, using the provided column names as map keys.
//
// Any fields that have a nil value or are not included in the given column map are not included in the returned map.
func TransformFields(columns map[string]string, fields sensor.Fields) map[string]any {
	values := make(map[string]any)
	for column, v := range CollectFields(columns, fields) {
		values[column] = v
	}
	return values
}
