package influxdb

import (
	"fmt"
	"log"
	"strings"

	influx "github.com/influxdata/influxdb1-client/v2"
	"github.com/wickerdj/ruuvitag-collector/pkg/sensor"
)

const dbAddr = "http://192.168.1.204:8086"
const dbUserName = "username"
const dbPassword = "password"
const dbName = "ruuvi"
const dbPrecision = "s"
const dbMeasurement = "readings"

// Write the data to an Influx data store
func Write(data sensor.Data) {
	c, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr: dbAddr,
		// UserName: dbUserName,
		// Password: dbPassword,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  dbName,
		Precision: dbPrecision,
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	point, err := influx.NewPoint(dbMeasurement, map[string]string{
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

	bp.AddPoint(point)

	c.Write(bp)
}
