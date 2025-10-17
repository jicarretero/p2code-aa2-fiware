package models

import (
	"fmt"
	"testing"
)

func TestCreateDeviceProfileJSONLD(t *testing.T) {
	var err error
	var dp_expected_result DeviceProfileJSONLD
	readings_expected_results := make(map[string]ReadingJSONLD)
	var dp DeviceData

	for k, Example := range Examples {
		dp_expected_result, err = DeserializeDeviceProfileJSONLD([]byte(ExamplesDeviceProfileJSONLD[k]))
		if err != nil {
			t.Errorf("Problem deserializing the Profile")
		}

		for _, example_reading := range ExamplesReadingJSONLD[k] {
			var reading ReadingJSONLD
			reading, err = DeserializeReadingJSONLD([]byte(example_reading))
			if err != nil {
				t.Errorf("Problem deserializing the Reading")
			}
			readings_expected_results[reading.GetId()] = reading
		}

		dp, err = DeserializeData([]byte(Example))
		if err != nil {
			t.Errorf("Problem deserializing MQTTT data Profile")
		}

		dp_results := NewDeviceProfileJSONLD(dp)

		for i, res := range dp_results {
			switch res.(type) {
			case ReadingJSONLD:
				expected_reading := readings_expected_results[res.GetId()]
				if !res.Equal(expected_reading) {
					t.Errorf("Reading data different to what it is expected: \n%v\n%v", res, expected_reading)
				}
			case DeviceProfileJSONLD:
				if !res.Equal(dp_expected_result) {
					t.Errorf("Data different to what it is expected: \n%v\n%v", res, dp_expected_result)
				}
			default:
				fmt.Printf("Unknown type of data: %T @%d\n", res, i)
			}
		}
	}
}
