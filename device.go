package onvif

var deviceXMLNs = []string{
	`xmlns:tds="http://www.onvif.org/ver10/device/wsdl"`,
	`xmlns:tt="http://www.onvif.org/ver10/schema"`,
}

// GetSystemDateAndTime fetch date and time from ONVIF camera
func (device Device) GetSystemDateAndTime() (string, error) {
	// Create SOAP
	soap := SOAP{
		Body:  "<tds:GetSystemDateAndTime/>",
		XMLNs: deviceXMLNs,
	}

	// Send SOAP request
	response, err := soap.SendRequest(device.XAddr)
	if err != nil {
		return "", err
	}

	// Parse response
	dateTime, _ := response.ValueForPathString("Envelope.Body.GetSystemDateAndTimeResponse.SystemDateAndTime")
	return dateTime, nil
}
