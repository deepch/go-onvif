package onvif

import (
	"encoding/json"
)

var deviceXMLNs = []string{
	`xmlns:tds="http://www.onvif.org/ver10/device/wsdl"`,
	`xmlns:tt="http://www.onvif.org/ver10/schema"`,
}

// GetDeviceInformation fetch information of ONVIF camera
func (device Device) GetDeviceInformation() (DeviceInformation, error) {
	// Create initial result
	result := DeviceInformation{}

	// Create SOAP
	soap := SOAP{
		Body:  "<tds:GetDeviceInformation/>",
		XMLNs: deviceXMLNs,
	}

	// Send SOAP request
	response, err := soap.SendRequest(device.XAddr)
	if err != nil {
		return result, err
	}

	// Parse response to interface
	deviceInfo, err := response.ValueForPath("Envelope.Body.GetDeviceInformationResponse")
	if err != nil {
		return result, err
	}

	// Parse interface to struct
	err = device.interfaceToStruct(&deviceInfo, &result)
	if err != nil {
		return result, err
	}

	return result, nil
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

func (device Device) interfaceToStruct(src, dst interface{}) error {
	bt, err := json.Marshal(&src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bt, &dst)
	if err != nil {
		return err
	}

	return nil
}
