package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFirstone(t *testing.T) {
	fmt.Println("This is the first test")
	for _, Example := range Examples {
		d, err := DeserializeData([]byte(Example))

		fmt.Printf("%p\n", &d)

		if err != nil {
			t.Errorf("Not achieved expected result")
		}

		_, err = SerializeData(d)
		if err != nil {
			t.Errorf("Not propert serialized in []byte")
		}
	}
}

func TestSecondone(t *testing.T) {
	fmt.Println("This is the second test - Should not be able to deserialize anything.")
	_, err := DeserializeData([]byte("hola mundo"))
	if err == nil {
		fmt.Println(reflect.TypeOf(err))
		t.Errorf("Expected error %s", err)
	}
}
