package onvif

import (
	"encoding/json"
	"strings"
)

var deviceXMLNs = []string{
	`xmlns:tds="http://www.onvif.org/ver10/device/wsdl"`,
	`xmlns:tt="http://www.onvif.org/ver10/schema"`,
}

// GetDeviceInformation fetch information of ONVIF camera
func (device Device) GetDeviceInformation() (DeviceInformation, error) {
	// Create SOAP
	soap := SOAP{
		Body:  "<tds:GetDeviceInformation/>",
		XMLNs: deviceXMLNs,
	}

	// Send SOAP request
	response, err := soap.SendRequest(device.XAddr)
	if err != nil {
		return DeviceInformation{}, err
	}

	// Parse response to interface
	deviceInfo, err := response.ValueForPath("Envelope.Body.GetDeviceInformationResponse")
	if err != nil {
		return DeviceInformation{}, err
	}

	// Parse interface to struct
	result := DeviceInformation{}
	err = interfaceToStruct(&deviceInfo, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetSystemDateAndTime fetch date and time from ONVIF camera
func (device Device) GetSystemDateAndTime() (string, error) {
	// Create SOAP
	soap := SOAP{
		Body:  "</tds:GetSystemDateAndTime>",
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

// GetCapabilities fetch info of ONVIF camera's capabilities
func (device Device) GetCapabilities() (DeviceCapabilities, error) {
	// Create SOAP
	soap := SOAP{
		XMLNs: deviceXMLNs,
		Body:  `<tds:GetCapabilities></tds:Category></tds:GetCapabilities>`,
	}

	// Send SOAP request
	response, err := soap.SendRequest(device.XAddr)
	if err != nil {
		return DeviceCapabilities{}, err
	}

	// Get network capabilities
	envelopeBodyPath := "Envelope.Body.GetCapabilitiesResponse.Capabilities"
	ifaceNetCap, err := response.ValueForPath(envelopeBodyPath + ".Device.Network")
	if err != nil {
		return DeviceCapabilities{}, err
	}

	netCap := NetworkCapabilities{}
	mapNetCap, ok := ifaceNetCap.(map[string]interface{})
	if ok {
		netCap.DynDNS = interfaceToBool(mapNetCap["DynDNS"])
		netCap.IPFilter = interfaceToBool(mapNetCap["IPFilter"])
		netCap.IPVersion6 = interfaceToBool(mapNetCap["IPVersion6"])
		netCap.ZeroConfig = interfaceToBool(mapNetCap["ZeroConfiguration"])
		netCap.Extension = make(map[string]bool)

		if mapNetExtension, ok := mapNetCap["Extension"].(map[string]interface{}); ok {
			for key, value := range mapNetExtension {
				netCap.Extension[key] = interfaceToBool(value)
			}
		}
	}

	// Get events capabilities
	ifaceEventsCap, err := response.ValueForPath(envelopeBodyPath + ".Events")
	if err != nil {
		return DeviceCapabilities{}, err
	}

	eventsCap := make(map[string]bool)
	if mapEventsCap, ok := ifaceEventsCap.(map[string]interface{}); ok {
		for key, value := range mapEventsCap {
			if strings.ToLower(key) == "xaddr" {
				continue
			}

			key = strings.Replace(key, "WS", "", 1)
			eventsCap[key] = interfaceToBool(value)
		}
	}

	// Get events capabilities
	ifaceStreamingCap, err := response.ValueForPath(envelopeBodyPath + ".Media.StreamingCapabilities")
	if err != nil {
		return DeviceCapabilities{}, err
	}

	streamingCap := make(map[string]bool)
	if mapStreamingCap, ok := ifaceStreamingCap.(map[string]interface{}); ok {
		for key, value := range mapStreamingCap {
			key = strings.Replace(key, "_", " ", -1)
			streamingCap[key] = interfaceToBool(value)
		}
	}

	// Get PTZ capabilities
	ptzCap := true
	if _, err = response.ValueForPath(envelopeBodyPath + ".PTZ"); err != nil {
		ptzCap = false
	}

	// Create final result
	deviceCapabilities := DeviceCapabilities{
		Network:   netCap,
		Events:    eventsCap,
		Streaming: streamingCap,
		PTZ:       ptzCap,
	}

	return deviceCapabilities, nil
}

func interfaceToStruct(src, dst interface{}) error {
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

func interfaceToString(src interface{}) string {
	str, _ := src.(string)
	return str
}

func interfaceToBool(src interface{}) bool {
	strBool := interfaceToString(src)
	return strings.ToLower(strBool) == "true"
}
