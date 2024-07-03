package sensor

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
}

var DefaultColumnMap map[string]string

func init() {
	DefaultColumnMap = make(map[string]string)
	for _, column := range DefaultColumns {
		DefaultColumnMap[column] = DefaultColumnMap[column]
	}
}
