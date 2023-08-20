package cputemp

import "testing"

func TestBasic(t *testing.T) {
	temp, err := Get()
	if err != nil {
		t.Fatal(err)
	}
	if temp < 10.0 {
		t.Fatal("to low")
	}
	if temp > 100.0 {
		t.Fatal("to high")
	}
	t.Logf("temp is %f°C", temp)
}
