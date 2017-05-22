package onvif

import (
	"encoding/json"
	"strings"
)

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
