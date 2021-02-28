package influxdb

import (
	"strings"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/wickerdj/ruuvitag-collector/pkg/sensor"
)

const dbAddr = "http://192.168.1.30:8086"
const dbBucketName = "ruuvi"
const dbOrg = "wick"
const dbToken = "sITKb6pogww_buCm0yxa3oymE9RFb8zk_zuo_YrCcJrGYgcS8SKMYAOdXtrJaJw_M-oUPMrO9BnVZghux1qL7g=="
const dbPrecision = "s"
const dbMeasurement = "readings"

// Write the data to an Influx data store
func Write(data sensor.Data) {
	client := influxdb2.NewClient(dbAddr, dbToken)

	defer client.Close()

	point := influxdb2.NewPoint(dbMeasurement, map[string]string{
		"mac":  strings.ToUpper(data.Addr),
		"name": data.Name,
	}, map[string]interface{}{
		"temperature":    data.Temperature,
		"humidity":       data.Humidity,
		"pressure":       data.Pressure,
		"battery":        data.Battery,
		"acceleration_x": data.AccelerationX,
		"acceleration_y": data.AccelerationY,
		"acceleration_z": data.AccelerationZ,
	}, data.Timestamp)

	writeAPI := client.WriteAPI(dbOrg, dbBucketName)
	writeAPI.WritePoint(point)
}
