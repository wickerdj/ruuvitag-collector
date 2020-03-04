package main

import (
	"context"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
	"github.com/wickerdj/ruuvitag-collector/pkg/influxdb"
	"github.com/wickerdj/ruuvitag-collector/pkg/sensor"
)

func setup(ctx context.Context) context.Context {
	d, err := dev.DefaultDevice()
	if err != nil {
		panic(err)
	}
	ble.SetDefaultDevice(d)

	return ble.WithSigHandler(context.WithCancel(ctx))
}

func main() {
	ctx := setup(context.Background())

	ble.Scan(ctx, true, handler, filter)
}

func handler(a ble.Advertisement) {
	d, err := sensor.Parse(a.ManufacturerData(), a.Addr().String())
	if err == nil {
		// fmt.Printf("[%s] RSSI: %3d: %+v\n", a.Addr(), a.RSSI(), d)
		influxdb.Write(d)
	}
}

func filter(a ble.Advertisement) bool {
	return sensor.IsRuuviTag(a.ManufacturerData())
}
