package influxdb

import (
	"context"
	"strings"

	influx "github.com/influxdata/influxdb1-client/v2"
	"github.com/wickerdj/ruuvitag-collector/pkg/sensor"
)

// InfluxdbExport - basic info
type InfluxdbExport struct {
	client      influx.Client
	database    string
	measurement string
}

// Config information for export
type Config struct {
	Addr        string
	Token       string
	Database    string
	Measurement string
	Username    string
	Password    string
}

// New - create new
func New(cfg Config) (InfluxdbExport, error) {
	client, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     cfg.Addr,
		Username: cfg.Username,
		Password: cfg.Password,
	})
	if err != nil {
		return nil, err
	}
	return InfluxdbExport{
		client:      client,
		database:    cfg.Database,
		measurement: cfg.Measurement,
	}, nil
}

// Name of this export
func Name() string {
	return "InfluxDB"
}

// Export the data
func Export(ctx context.Context, data sensor.Data) error {
	conf := influx.BatchPointsConfig{
		Database: e.database,
	}
	bp, err := influx.NewBatchPoints(conf)
	if err != nil {
		return err
	}
	point, err := influx.NewPoint(e.measurement, map[string]string{
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
	return e.client.Write(bp)
}

// Close the export
func Close() error {
	return e.client.Close()
}
