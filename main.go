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
	fmt.Println("ManufacturerData", a.ManufacturerData)
	fmt.Println("LocalName", a.LocalName)
	fmt.Println("ServiceData", a.ServiceData)
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
