package main

import (
	"os"
	"reflect"
	"testing"
)

func TestInitialization(t *testing.T) {
	os.Args = []string{"alien-invasion.go", "test/input.txt", "4"}
	main()
	if len(aliens) != 4 {
		t.Error("Expected to create 4 aliens. Got", len(aliens))
	}
	if len(cities) != 2 {
		t.Error("Expected to create 2 cities. Got", len(cities))
	}
	expectedCitiesMap := map[string][]string{
		"Bar": []string{"Foo", "Bee"},
		"Foo": []string{"Bar", "Baz", "Qu-ux"},
	}
	if !reflect.DeepEqual(expectedCitiesMap, cities) {
		t.Error("Expected cities map to be", expectedCitiesMap, "Found", cities)
	}
}
