package main

import (
	"fmt"
	"log"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
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

	if p.ID() == "F8:1D:CC:84:AE:D4" {
		d, err := Parse(a.ManufacturerData)
		if err != nil {
			log.Fatalln("bad data")
		}
		fmt.Printf("\tAddr: %s\n\tTemperature: %s\n", d.Addr, d.Temperature)
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
