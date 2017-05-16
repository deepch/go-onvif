package onvif

// Device contains data of ONVIF camera
type Device struct {
	ID       string
	Name     string
	XAddr    string
	User     string
	Password string
}

// DeviceInformation contains information of ONVIF camera
type DeviceInformation struct {
	FirmwareVersion string
	HardwareID      string
	Manufacturer    string
	Model           string
	SerialNumber    string
}
