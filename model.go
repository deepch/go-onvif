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

// NetworkCapabilities contains networking capabilities of ONVIF camera
type NetworkCapabilities struct {
	DynDNS     bool
	IPFilter   bool
	IPVersion6 bool
	ZeroConfig bool
}

// DeviceCapabilities contains capabilities of an ONVIF camera
type DeviceCapabilities struct {
	Network   NetworkCapabilities
	Events    map[string]bool
	Streaming map[string]bool
}

// HostnameInformation contains hostname info of an ONVIF camera
type HostnameInformation struct {
	Name      string
	FromDHCP  bool
	Extension string
}
