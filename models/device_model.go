package models

import (
	"encoding/json"
)

type DeviceProfile struct {
	APIVersion  string `json:"apiVersion"`
	ID          string `json:"id"`
	DeviceName  string `json:"deviceName"`
	ProfileName string `json:"profileName"`
	SourceName  string `json:"sourceName"`
	Origin      int64  `json:"origin"`
}

type Reading struct {
	ID           string      `json:"id"`
	Origin       int64       `json:"origin"`
	DeviceName   string      `json:"deviceName"`
	ResourceName string      `json:"resourceName"`
	ProfileName  string      `json:"profileName"`
	ValueType    string      `json:"valueType"`
	BinaryValue  interface{} `json:"binaryValue"` // store as an empty interface to preserve JSON data
	MediaType    string      `json:"mediaType"`
	Value        interface{} `json:"value"`
	// Value        float64     `json:"value"`
}

type DeviceData struct {
	DeviceProfile
	Readings []Reading `json:"readings"`
}

func DeserializeData(data []byte) (DeviceData, error) {
	var deviceData DeviceData
	err := json.Unmarshal(data, &deviceData)
	// handle the error...
	return deviceData, err
}

func SerializeData(deviceData DeviceData) ([]byte, error) {
	data, err := json.Marshal(deviceData)
	if err != nil {
		return nil, err
	}
	return data, nil
}
