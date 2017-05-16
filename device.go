package onvif

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/clbanning/mxj"
)

var httpClient = &http.Client{
	Timeout: time.Second * 5,
}

// Device contains data of ONVIF camera
type Device struct {
	ID       string
	Name     string
	XAddr    string
	User     string
	Password string
}

// GetSystemDateAndTime fetch date and time from ONVIF camera
func (device Device) GetSystemDateAndTime() (string, error) {
	// Create request
	request := `<?xml version="1.0" encoding="UTF-8"?>
		<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope">
			<s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
				<GetSystemDateAndTime xmlns="http://www.onvif.org/ver10/device/wsdl"/>
			</s:Body>
		</s:Envelope>`

	request = regexp.MustCompile(`\>\s+\<`).ReplaceAllString(request, "><")
	request = regexp.MustCompile(`\s+`).ReplaceAllString(request, " ")

	// Create request
	buffer := bytes.NewBuffer([]byte(request))
	req, err := http.NewRequest("POST", device.XAddr, buffer)
	req.Header.Set("Content-Type", "application/soap+xml")
	req.Header.Set("Charset", "utf-8")

	// Send request
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse XML to map
	mapXML, err := mxj.NewMapXml(responseBody)
	if err != nil {
		return "", err
	}

	dateTime, _ := mapXML.ValueForPathString("Envelope.Body.GetSystemDateAndTimeResponse.SystemDateAndTime")
	return dateTime, nil
}
