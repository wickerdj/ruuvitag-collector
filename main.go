package main

import (
	"fmt"
	"log"
	"strings"

	influx "github.com/influxdata/influxdb1-client/v2"
	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	"github.com/wickerdj/ruuvitag-collector/pkg/sensor"
)

func onStateChanged(device gatt.Device, s gatt.State) {
	switch s {
	case gatt.StatePoweredOn:
		device.Scan([]gatt.UUID{}, true)
		return
	default:
		device.StopScanning()
	}
}

func onDiscovery(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {

	fmt.Printf("\nPeripheral ID:%s, NAME:(%s)\n", p.ID(), p.Name())
	fmt.Println("  TX Power Level    =", a.TxPowerLevel)

	d, err := sensor.Parse(a.ManufacturerData, p.ID())
	if err != nil {
		log.Printf("bad data id:%v\n", p.ID())
	} else {
		fmt.Printf("\tAddr: %v\n\tTemperature: %v\n\tHumidity: %v\n\tBattery: %v\n\tTimestamp: %v\n", d.Addr, d.Temperature, d.Humidity, d.Battery, d.Timestamp)
	}

	write(d)
}

func write(data sensor.Data) {
	c, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr: "http://192.168.1.204:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  "ruuvi",
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	point, err := influx.NewPoint("readings", map[string]string{
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

func main() {

	device, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Problem with device err: %s\n", err)
	}

	device.Handle(gatt.PeripheralDiscovered(onDiscovery))
	device.Init(onStateChanged)
	select {}
}
