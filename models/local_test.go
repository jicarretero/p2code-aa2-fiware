package models

import (
	"fmt"
	"testing"
)

func TestMap01(testing *testing.T) {
	// Mapa/diccionario de stings
	d1 := make(map[string]string)

	// Acceso al mapa
	d1["jicg"] = "joseignacio.carretereo"
	d1["fla"] = "fernando.lopez"
	d1["abella"] = "alberto.abella"

	// Recorrido del mapa
	for k, v := range d1 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	fmt.Printf("len: %d\n", len(d1))
}

type Person struct {
	Name    string
	Surname string
	Age     int
}

func TestMapGeneric(testing *testing.T) {
	// Mapa/diccionario de stings
	d1 := make(map[interface{}]interface{})

	// Acceso al mapa
	d1["jicg"] = Person{"jose ignacio", "carretereo", 51}
	d1["fla"] = "fernando.lopez"
	d1["abella"] = "alberto.abella"

	// Recorrido del mapa
	for k, v := range d1 {
		fmt.Printf("Key: %s, Value: %+v\n", k, v)
	}

	fmt.Printf("len: %d\n", len(d1))
}
