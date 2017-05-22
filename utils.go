package onvif

import (
	"encoding/json"
	"strings"
)

var testDevice = Device{
	XAddr: "http://192.168.1.75:5000/onvif/device_service",
}

func interfaceToString(src interface{}) string {
	str, _ := src.(string)
	return str
}

func interfaceToBool(src interface{}) bool {
	strBool := interfaceToString(src)
	return strings.ToLower(strBool) == "true"
}

func prettyJSON(src interface{}) string {
	result, _ := json.MarshalIndent(&src, "", "    ")
	return string(result)
}
