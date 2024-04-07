package test

import (
	"reflect"
	"testing"
)

func TestCreateEntity(t *testing.T) {
	c := CreateEntity()

	// if reflect.TypeOf(x).Kind() ==
	if reflect.TypeOf(c).Elem().Name() != "Collector" {
		t.Errorf("expected %s but got %s", "Collector", reflect.TypeOf(c).Elem().Name())
	}

}
