package brokerld

import "testing"

func TestPath(t *testing.T) {
	base_path := "http://localhost:1026"
	full_path := GetFullPath(base_path)

	if full_path != "http://localhost:1026/ngsi-ld/v1/entities" {
		t.Errorf("Expected %s, got %s", "http://localhost:1026/ngsi-ld/v1/entities", full_path)
	}
}
