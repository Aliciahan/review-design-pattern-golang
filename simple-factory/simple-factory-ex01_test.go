package simple_factory

import "testing"

func TestTypes(t *testing.T) {
	api := NewApi(1)
	if api.Say("Tom") != "Hi, Tom" {
		t.Fatal("Problem Init Type 1")
	}
	api = NewApi(2)
	if api.Say("Tom") != "Hello, Tom" {
		t.Fatal("Problem Init Type 1")
	}
}
