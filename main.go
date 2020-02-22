package main

import (
	"fmt"
	"log"

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
