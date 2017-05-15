package onvif

// Device contains data of Onvif device
type Device struct {
	ID       string
	Name     string
	XAddrs   []string
	User     string
	Password string
}
