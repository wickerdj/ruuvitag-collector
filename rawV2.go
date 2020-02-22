package main

// DataFormat5 - Data Format 5 Protocol Specification (RAWv2)
type DataFormat5 struct {
	ManufacturerID    uint16
	DataFormat        uint8
	Temperature       int16
	Humidity          uint16
	Pressure          uint16
	AccelerationX     int16
	AccelerationY     int16
	AccelerationZ     int16
	BatteryVoltage    uint16
	MovementCounter   uint8
	MeasurementNumber uint16
}

// ParseSensorFormat5 parses the data
func ParseSensorFormat5(data []byte) (sd Data, err error) {
	reader := bytes.NewReader(data)
	var result DataFormat5
	err = binary.Read(reader, binary.BigEndian, &result)
	if err != nil {
		return
	}
	sd.Temperature = float64(result.Temperature) * 0.005
	sd.Humidity = float64(result.Humidity) / 400.0
	sd.Pressure = float64(int(result.Pressure)+50000) / 100.0
	sd.AccelerationX = int(result.AccelerationX)
	sd.AccelerationY = int(result.AccelerationY)
	sd.AccelerationZ = int(result.AccelerationZ)
	sd.Battery = int(result.BatteryVoltage >> 5)
	sd.MovementCounter = int(result.MovementCounter)
	return
}
