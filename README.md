# ruuvitag-collector
Listens for data coming from a [RuuviTag](https://ruuvi.com/ruuvitag-specs/) sensor and pushes the data if InfluxDB.

# Additional Libraries
`go get github.com/paypal/gatt`

# Disable of Bluetooth Scanning on the Raspberry Pi
`sudo hciconfig hci0 down`

`sudo service bluetooth stop`

The above commands are good for testing. If the Raspberry Pi is rebooted the start up script will restart Bluetooth Scanning

# Build for the Rasberry Pi
`env GOOS=linux GOARCH=arm GOARM=5 go build`

# Inspired by
This project was inspired by the work of
- [https://github.com/peterhellberg/ruuvitag](https://github.com/peterhellberg/ruuvitag)
- [https://github.com/srados/blistener](https://github.com/srados/blistener)
- [https://github.com/niktheblak/ruuvitag-gollector](https://github.com/niktheblak/ruuvitag-gollector)
- [https://github.com/Turee/goruuvitag](https://github.com/Turee/goruuvitag)
- [Scan For BLE iBeacon Devices with Golang on A Raspberry Pi Zero W](https://www.thepolyglotdeveloper.com/2018/02/scan-ble-ibeacon-devices-golang-raspberry-pi-zero-w/)
