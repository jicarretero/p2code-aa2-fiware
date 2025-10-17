package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const DeviceProfileJSONLDType = "https://p2code-project.eu/aa2/uc1"
const ReadingJSONLDType = "https://p2code-project.eu/aa2/uc1/reading"

const NGSILDHeadURNReading = "ngsild:urn:reading"
const NGSILDHeadURNDeviceProfile = "ngsild:urn:device-profile"

type DeviceProfileJSONLD struct {
	ID          string        `json:"id"`
	Type        string        `json:"type"`
	UUID        Property      `json:"uuid"`
	DeviceName  Property      `json:"deviceName"`
	ProfileName Property      `json:"profileName"`
	SourceName  Property      `json:"sourceName"`
	Origin      PropertyInt64 `json:"origin"`
	Readings    Relationship  `json:"readings"`
}

type ReadingJSONLD struct {
	ID           string        `json:"id"`
	Type         string        `json:"type"`
	UUID         Property      `json:"uuid"`
	Origin       PropertyInt64 `json:"origin"`
	DeviceName   Property      `json:"deviceName"`
	ResourceName Property      `json:"resourceName"`
	ProfileName  Property      `json:"profileName"`
	ValueType    Property      `json:"valueType"`
	BinaryValue  Property      `json:"binaryValue"`
	MediaType    Property      `json:"mediaType"`
	Value        Property      `json:"value"`
}

type JSONLDObject interface {
	GetId() string
	GetType() string
	Equal(other JSONLDObject) bool
}

func (d DeviceProfileJSONLD) GetId() string {
	return d.ID
}
func (d DeviceProfileJSONLD) GetType() string {
	return d.Type
}
func (r ReadingJSONLD) GetId() string {
	return r.ID
}
func (r ReadingJSONLD) GetType() string {
	return r.Type
}

type PropertyInt64 struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

type PropertyFloat64 struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type Property struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type Relationship struct {
	Type   string   `json:"type"`
	Object []string `json:"object"`
}

func (r1 *Relationship) Equal(r2 *Relationship) bool {
	if len(r1.Object) != len(r2.Object) {
		return false
	}

	setR1 := make(map[string]bool)
	for _, item := range r1.Object {
		setR1[item] = true
	}
	for _, item := range r2.Object {
		if _, exists := setR1[item]; !exists {
			return false
		}
	}

	return true
}

// Test equality for a Property entity of type int64
func (p1 *PropertyInt64) Equal(p2 *PropertyInt64) bool {
	return p1.Type == p2.Type && reflect.DeepEqual(p1.Value, p2.Value)
}

// Test equality for a Property entity of type float64
func (p1 *PropertyFloat64) Equal(p2 *PropertyFloat64) bool {
	return p1.Type == p2.Type && reflect.DeepEqual(p1.Value, p2.Value)
}

// Test equality for 2 Properties.
func (p1 *Property) BinaryEqual(p2 *Property) bool {
	return p1.Type == p2.Type && reflect.DeepEqual(p1.Value, p2.Value)
}

// Test equality for 2 Properties.
func (p1 *Property) Equal(p2 *Property) bool {
	return p1.Type == p2.Type && reflect.DeepEqual(p1.Value, p2.Value)
}

// DeserializeDeviceProfileJSONLD deserializes a JSON byte slice into a DeviceProfileJSONLD struct.
// It returns the struct and an error if the deserialization fails.
func DeserializeDeviceProfileJSONLD(data []byte) (DeviceProfileJSONLD, error) {
	var profile DeviceProfileJSONLD
	err := json.Unmarshal(data, &profile)
	return profile, err
}

// DeserializeReadingJSONLD deserializes a JSON byte slice into a ReadingJSONLD struct.
// It returns the struct and an error if the deserialization fails.
func DeserializeReadingJSONLD(data []byte) (ReadingJSONLD, error) {
	var reading ReadingJSONLD
	err := json.Unmarshal(data, &reading)
	return reading, err
}

// NewDeviceProfileJSONLD creates a new DeviceProfileJSONLD struct from a DeviceData struct.
// This means from what it is supposed to come from the MQTT reading
func NewDeviceProfileJSONLD(dp DeviceData) []JSONLDObject {
	profile := DeviceProfileJSONLD{}

	profile.ID = NGSILDHeadURNDeviceProfile + ":" + strings.ToLower(dp.DeviceName) + ":" +
		strings.ToLower(dp.ProfileName) + ":" + strings.ToLower(dp.SourceName)
	profile.Type = DeviceProfileJSONLDType

	profile.UUID = Property{Type: "Property", Value: dp.ID}
	profile.DeviceName = Property{Type: "Property", Value: dp.DeviceName}
	profile.SourceName = Property{Type: "Property", Value: dp.SourceName}
	profile.ProfileName = Property{Type: "Property", Value: dp.ProfileName}
	profile.Origin = PropertyInt64{Type: "Property", Value: dp.Origin}
	fmt.Printf("ORIGIN: [%v] %v\n", reflect.TypeOf(profile.Origin.Value), profile.Origin.Value)
	len_readings := len(dp.Readings)
	profile.Readings = Relationship{Type: "Relationship", Object: make([]string, len_readings)}

	res := make([]JSONLDObject, len(dp.Readings)+1)

	for j := range len_readings {
		readingId := NGSILDHeadURNReading + ":" + strings.ToLower(dp.Readings[j].DeviceName) + ":" +
			strings.ToLower(dp.Readings[j].ProfileName) + ":" + strings.ToLower(dp.SourceName) + ":" +
			strings.ToLower(dp.Readings[j].ResourceName)
		profile.Readings.Object[j] = readingId
		reading := ReadingJSONLD{}
		reading.ID = readingId
		reading.Type = ReadingJSONLDType
		reading.UUID = Property{Type: "Property", Value: dp.Readings[j].ID}
		reading.Origin = PropertyInt64{Type: "Property", Value: dp.Readings[j].Origin}
		reading.DeviceName = Property{Type: "Property", Value: dp.Readings[j].DeviceName}
		reading.ResourceName = Property{Type: "Property", Value: dp.Readings[j].ResourceName}
		reading.ProfileName = Property{Type: "Property", Value: dp.Readings[j].ProfileName}
		reading.ValueType = Property{Type: "Property", Value: dp.Readings[j].ValueType}
		reading.MediaType = Property{Type: "Property", Value: dp.Readings[j].MediaType}
		reading.Value = Property{Type: "Property", Value: dp.Readings[j].Value}
		if dp.Readings[j].BinaryValue != nil {
			reading.BinaryValue = Property{Type: "Property", Value: base64.StdEncoding.EncodeToString(dp.Readings[j].BinaryValue.([]byte))}
		} else {
			reading.BinaryValue = Property{Type: "Property", Value: "null"}
		}
		res[j] = reading
	}

	res[len_readings] = profile
	return res
}

func (r1 ReadingJSONLD) Equal(r2 JSONLDObject) bool {
	v, ok := r2.(ReadingJSONLD)
	if !ok {
		return false
	}

	fmt.Printf("Profile names: [%v] [%v] %v\n", r1.Value.Value, v.Value.Value, r1.Value.Value == v.Value.Value)
	fmt.Printf("Profile names: [%v] [%v]\n", reflect.TypeOf(r1.Value.Value), reflect.TypeOf(v.Value.Value))

	return r1.ID == v.ID &&
		r1.Type == v.Type &&
		r1.Origin.Equal(&v.Origin) &&
		r1.UUID.Equal(&v.UUID) &&
		r1.DeviceName.Equal(&v.DeviceName) &&
		r1.ResourceName.Equal(&v.ResourceName) &&
		r1.ProfileName.Equal(&v.ProfileName) &&
		r1.ValueType.Equal(&v.ValueType) &&
		r1.BinaryValue.BinaryEqual(&v.BinaryValue) &&
		r1.MediaType.Equal(&v.MediaType) &&
		r1.Value.Equal(&v.Value)
}

func (d1 DeviceProfileJSONLD) Equal(d2 JSONLDObject) bool {
	v, ok := d2.(DeviceProfileJSONLD)
	if !ok {
		return false
	}
	return d1.ID == v.ID && d1.Type == v.Type &&
		d1.UUID.Equal(&v.UUID) &&
		d1.Origin.Equal(&v.Origin) &&
		d1.ProfileName.Equal(&v.ProfileName) &&
		d1.SourceName.Equal(&v.SourceName) &&
		d1.DeviceName.Equal(&v.DeviceName) &&
		d1.Readings.Equal(&v.Readings)
}

// Create a Struct of type ReadingJSONLD
func NewReadingJSONLD(rd Reading) *ReadingJSONLD {
	reading := ReadingJSONLD{}
	reading.ID = NGSILDHeadURNReading + ":" + strings.ToLower(rd.DeviceName) + ":" +
		strings.ToLower(rd.ProfileName) + ":" + strings.ToLower(rd.ResourceName)

	reading.Type = ReadingJSONLDType

	reading.UUID = Property{Type: "Property", Value: rd.ID}
	reading.DeviceName = Property{Type: "Property", Value: rd.DeviceName}
	reading.ResourceName = Property{Type: "Property", Value: rd.ResourceName}
	reading.ProfileName = Property{Type: "Property", Value: rd.ProfileName}
	reading.ValueType = Property{Type: "Property", Value: rd.ValueType}
	reading.MediaType = Property{Type: "Property", Value: rd.MediaType}
	reading.Value = Property{Type: "Property", Value: rd.Value}

	var value string

	if rd.BinaryValue != nil {
		value = "null"
	} else {
		value = base64.StdEncoding.EncodeToString(rd.BinaryValue.([]byte))
	}
	reading.BinaryValue = Property{Type: "Property", Value: value}

	return &reading
}
