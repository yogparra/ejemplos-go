package main

import "testing"

func TestHolaMundo(t *testing.T) {
	expected := "Hola mundo desde go."
	got := HolaMundo()

	if expected != got {
		t.Errorf("got '%s' expected '%s'", got, expected)
	}
}
