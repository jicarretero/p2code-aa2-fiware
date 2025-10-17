package models

var Examples = []string{`{
    "apiVersion": "v1",
    "id": "f8d238ae-1c9f-46c2-b550-97611b2e063b",
    "deviceName": "Test-Modbus-device",
    "profileName": "MOD-ETH",
    "sourceName": "Poll-Modbus-000",
    "origin": 1752590782718007600,
    "readings": [
        {
            "id": "b3534346-a4b4-4bb7-88bc-0193b12878ac",
            "origin": 1752590782713848800,
            "deviceName": "Test-Modbus-device",
            "resourceName": "P",
            "profileName": "MOD-ETH",
            "valueType": "Float31",
            "binaryValue": null,
            "mediaType": "",
            "value": -64.16166
        },
        {
            "id": "f1aa2f7a-7eae-40d5-9b6a-ed8ab60b018c",
            "origin": 1752590782715421400,
            "deviceName": "Test-Modbus-device",
            "resourceName": "Q",
            "profileName": "MOD-ETH",
            "valueType": "Float31",
            "binaryValue": null,
            "mediaType": "",
            "value": -3.323915
        },
        {
            "id": "79af812e-f41c-4f2d-95b8-adb7d9ffa9e1",
            "origin": 1752590782716782000,
            "deviceName": "Test-Modbus-device",
            "resourceName": "V",
            "profileName": "MOD-ETH",
            "valueType": "Float31",
            "binaryValue": null,
            "mediaType": "",
            "value": 153.33
        },
        {
            "id": "ca58834a-7f40-4c4a-bec5-b111c68c7ca9",
            "origin": 1752590782717809200,
            "deviceName": "Test-Modbus-device",
            "resourceName": "I",
            "profileName": "MOD-ETH",
            "valueType": "Float31",
            "binaryValue": null,
            "mediaType": "",
            "value": 237.0924
        }
    ]
}`, `{
    "apiVersion": "v2",
    "id": "8d271365-e1c4-4553-9e4e-eb0690d85f77",
    "deviceName": "Test-Modbus-device",
    "profileName": "MOD-ETH",
    "sourceName": "Poll-Modbus-001",
    "origin": 1760104201927662040,
    "readings": [
    {
        "id": "8034d75b-ad57-4295-a9d9-96632b449880",
        "origin": 1760104201922874762,
        "deviceName": "Test-Modbus-device",
        "resourceName": "P",
        "profileName": "MOD-ETH",
        "valueType": "Float32",
        "binaryValue": null,
        "mediaType": "",
        "value": "-6.220587e+01"
    },
    {
        "id": "328853d9-3c2e-4354-bfee-fb03f8819ee8",
        "origin": 1760104201923484816,
        "deviceName": "Test-Modbus-device",
        "resourceName": "Q",
        "profileName": "MOD-ETH",
        "valueType": "Float32",
        "binaryValue": null,
        "mediaType": "",
        "value": "-2.529164e+00"
    },
    {
        "id": "c7c09b10-1314-4261-9a4a-65c3f38e56ef",
        "origin": 1760104201925756531,
        "deviceName": "Test-Modbus-device",
        "resourceName": "V",
        "profileName": "MOD-ETH",
        "valueType": "Float32",
        "binaryValue": null,
        "mediaType": "",
        "value": "1.544100e+02"
    },
    {
        "id": "7f18df9d-de55-4d6a-a839-0e415119932f",
        "origin": 1760104201926867660,
        "deviceName": "Test-Modbus-device",
        "resourceName": "I",
        "profileName": "MOD-ETH",
        "valueType": "Float32",
        "binaryValue": null,
        "mediaType": "",
        "value": "2.400454e+02"
    }]
}`}

var ExamplesDeviceProfileJSONLD = []string{`{
    "id": "ngsild:urn:device-profile:test-modbus-device:mod-eth:poll-modbus-000",
    "type": "https://p2code-project.eu/aa2/uc1",
    "uuid": {
        "type": "Property",
        "value": "f8d238ae-1c9f-46c2-b550-97611b2e063b"
    },
    "deviceName": {
        "type": "Property",
        "value": "Test-Modbus-device"
    },
    "profileName": {
        "type": "Property",
        "value": "MOD-ETH"
    },
    "sourceName": {
        "type": "Property",
        "value": "Poll-Modbus-000"
    },
	"origin": {
		"type": "Property",
		"value": 1752590782718007600
	},
	"readings": {
		"type": "Relationship",
		"object": [
			"ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-000:q",
			"ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-000:p",
			"ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-000:v",
			"ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-000:i"
		]
	}
}`,
	`{
    "id": "ngsild:urn:device-profile:test-modbus-device:mod-eth:poll-modbus-001",
    "type": "https://p2code-project.eu/aa2/uc1",
    "uuid": {
        "type": "Property",
        "value": "8d271365-e1c4-4553-9e4e-eb0690d85f77"
    },
    "deviceName": {
        "type": "Property",
        "value": "Test-Modbus-device"
    },
    "profileName": {
        "type": "Property",
        "value": "MOD-ETH"
    },
    "sourceName": {
        "type": "Property",
        "value": "Poll-Modbus-001"
    },
	"origin": {
		"type": "Property",
		"value": 1760104201927662040
	},
	"readings": {
		"type": "Relationship",
		"object": [
			"ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-001:q",
			"ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-001:p",
			"ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-001:v",
			"ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-001:i"
		]
	}
}`,
} //end ExamplesDeviceProfileJSONLD

var ExamplesReadingJSONLD = [][]string{
	{
		`{
    "id": "ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-000:p",
    "type": "https://p2code-project.eu/aa2/uc1/reading",
    "uuid": {
        "type": "Property",
        "value": "b3534346-a4b4-4bb7-88bc-0193b12878ac"
    },
    "origin": {
        "type": "Property",
        "value": 1752590782713848800
    },
    "deviceName": {
        "type": "Property",
        "value": "Test-Modbus-device"
    },
    "resourceName": {
        "type": "Property",
        "value": "P"
    },
    "profileName": {
        "type": "Property",
        "value": "MOD-ETH"
    },
    "valueType": {
        "type": "Property",
        "value": "Float31"
    },
    "binaryValue": {
        "type": "Property",
        "value": "null"
    },
    "mediaType": {
        "type": "Property",
        "value": ""
    },
    "value": {
        "type": "Property",
        "value": -64.16166
    }
}`,
		`{
    "id": "ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-000:q",
    "type": "https://p2code-project.eu/aa2/uc1/reading",
    "uuid": {
        "type": "Property",
        "value": "f1aa2f7a-7eae-40d5-9b6a-ed8ab60b018c"
    },
    "origin": {
        "type": "Property",
        "value": 1752590782715421400
    },
    "deviceName": {
        "type": "Property",
        "value": "Test-Modbus-device"
    },
    "resourceName": {
        "type": "Property",
        "value": "Q"
    },
    "profileName": {
        "type": "Property",
        "value": "MOD-ETH"
    },
    "valueType": {
        "type": "Property",
        "value": "Float31"
    },
    "binaryValue": {
        "type": "Property",
        "value": "null"
    },
    "mediaType": {
        "type": "Property",
        "value": ""
    },
    "value": {
        "type": "Property",
        "value": -3.323915
    }
}`,
		`{
    "id": "ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-000:v",
    "type": "https://p2code-project.eu/aa2/uc1/reading",
    "uuid": {
        "type": "Property",
        "value": "79af812e-f41c-4f2d-95b8-adb7d9ffa9e1"
    },
    "origin": {
        "type": "Property",
        "value": 1752590782716782000
    },
    "deviceName": {
        "type": "Property",
        "value": "Test-Modbus-device"
    },
    "resourceName": {
        "type": "Property",
        "value": "V"
    },
    "profileName": {
        "type": "Property",
        "value": "MOD-ETH"
    },
    "valueType": {
        "type": "Property",
        "value": "Float31"
    },
    "binaryValue": {
        "type": "Property",
        "value": "null"
    },
    "mediaType": {
        "type": "Property",
        "value": ""
    },
    "value": {
        "type": "Property",
        "value": 153.33
    }
}`,
		`{
    "id": "ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-000:i",
    "type": "https://p2code-project.eu/aa2/uc1/reading",
    "uuid": {
        "type": "Property",
        "value": "ca58834a-7f40-4c4a-bec5-b111c68c7ca9"
    },
    "origin": {
        "type": "Property",
        "value": 1752590782717809200
    },
    "deviceName": {
        "type": "Property",
        "value": "Test-Modbus-device"
    },
    "resourceName": {
        "type": "Property",
        "value": "I"
    },
    "profileName": {
        "type": "Property",
        "value": "MOD-ETH"
    },
    "valueType": {
        "type": "Property",
        "value": "Float31"
    },
    "binaryValue": {
        "type": "Property",
        "value": "null"
    },
    "mediaType": {
        "type": "Property",
        "value": ""
    },
    "value": {
        "type": "Property",
        "value": 237.0924
    }
}`,
	},
	{
		`{
    "id": "ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-001:p",
    "type": "https://p2code-project.eu/aa2/uc1/reading",
    "uuid": {
        "type": "Property",
        "value": "8034d75b-ad57-4295-a9d9-96632b449880"
    },
    "origin": {
        "type": "Property",
        "value": 1760104201922874762
    },
    "deviceName": {
        "type": "Property",
        "value": "Test-Modbus-device"
    },
    "resourceName": {
        "type": "Property",
        "value": "P"
    },
    "profileName": {
        "type": "Property",
        "value": "MOD-ETH"
    },
    "valueType": {
        "type": "Property",
        "value": "Float32"
    },
    "binaryValue": {
        "type": "Property",
        "value": "null"
    },
    "mediaType": {
        "type": "Property",
        "value": ""
    },
    "value": {
        "type": "Property",
        "value": "-6.220587e+01"
    }
}`,
		`{
    "id": "ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-001:q",
    "type": "https://p2code-project.eu/aa2/uc1/reading",
    "uuid": {
        "type": "Property",
        "value": "328853d9-3c2e-4354-bfee-fb03f8819ee8"
    },
    "origin": {
        "type": "Property",
        "value": 1760104201923484816
    },
    "deviceName": {
        "type": "Property",
        "value": "Test-Modbus-device"
    },
    "resourceName": {
        "type": "Property",
        "value": "Q"
    },
    "profileName": {
        "type": "Property",
        "value": "MOD-ETH"
    },
    "valueType": {
        "type": "Property",
        "value": "Float32"
    },
    "binaryValue": {
        "type": "Property",
        "value": "null"
    },
    "mediaType": {
        "type": "Property",
        "value": ""
    },
    "value": {
        "type": "Property",
        "value": "-2.529164e+00"
    }
}`,
		`{
    "id": "ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-001:v",
    "type": "https://p2code-project.eu/aa2/uc1/reading",
    "origin": {
        "type": "Property",
        "value": 1760104201925756531
    },
    "uuid": {
        "type": "Property",
        "value": "c7c09b10-1314-4261-9a4a-65c3f38e56ef"
    },
    "deviceName": {
        "type": "Property",
        "value": "Test-Modbus-device"
    },
    "resourceName": {
        "type": "Property",
        "value": "V"
    },
    "profileName": {
        "type": "Property",
        "value": "MOD-ETH"
    },
    "valueType": {
        "type": "Property",
        "value": "Float32"
    },
    "binaryValue": {
        "type": "Property",
        "value": "null"
    },
    "mediaType": {
        "type": "Property",
        "value": ""
    },
    "value": {
        "type": "Property",
        "value": "1.544100e+02"
    }
}`,
		`{
    "id": "ngsild:urn:reading:test-modbus-device:mod-eth:poll-modbus-001:i",
    "type": "https://p2code-project.eu/aa2/uc1/reading",
    "uuid": {
        "type": "Property",
        "value": "7f18df9d-de55-4d6a-a839-0e415119932f"
    },
    "origin": {
        "type": "Property",
        "value": 1760104201926867660
    },
    "deviceName": {
        "type": "Property",
        "value": "Test-Modbus-device"
    },
    "resourceName": {
        "type": "Property",
        "value": "I"
    },
    "profileName": {
        "type": "Property",
        "value": "MOD-ETH"
    },
    "valueType": {
        "type": "Property",
        "value": "Float32"
    },
    "binaryValue": {
        "type": "Property",
        "value": "null"
    },
    "mediaType": {
        "type": "Property",
        "value": ""
    },
    "value": {
        "type": "Property",
        "value": "2.400454e+02"
    }
}`,
	},
}
