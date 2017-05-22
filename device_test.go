package onvif

import (
	"fmt"
	"log"
	"testing"
)

var defaultDevice = Device{
	XAddr: "http://192.168.1.75:5000/onvif/device_service",
}

func TestGetInformation(t *testing.T) {
	log.Println("Test GetInformation")

	res, err := defaultDevice.GetInformation()
	if err != nil {
		t.Error(err)
	}

	js := prettyJSON(&res)
	fmt.Println(js)
}

func TestGetCapabilities(t *testing.T) {
	log.Println("Test GetCapabilities")

	res, err := defaultDevice.GetCapabilities()
	if err != nil {
		t.Error(err)
	}

	js := prettyJSON(&res)
	fmt.Println(js)
}

func TestGetDiscoveryMode(t *testing.T) {
	log.Println("Test GetDiscoveryMode")

	res, err := defaultDevice.GetDiscoveryMode()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestGetScopes(t *testing.T) {
	log.Println("Test GetScopes")

	res, err := defaultDevice.GetScopes()
	if err != nil {
		t.Error(err)
	}

	js := prettyJSON(&res)
	fmt.Println(js)
}

func TestGetHostname(t *testing.T) {
	log.Println("Test GetHostname")

	res, err := defaultDevice.GetHostname()
	if err != nil {
		t.Error(err)
	}

	js := prettyJSON(&res)
	fmt.Println(js)
}
