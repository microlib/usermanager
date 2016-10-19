package main

import "testing"

func TestMock(t *testing.T) {
	stringtest := "hey"
	if stringtest != "hey" {
		t.Errorf("Expected hey, got", stringtest)
	}
}
