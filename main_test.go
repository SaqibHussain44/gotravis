package main

import (
	"testing"
)

func TestGetName(t *testing.T) {
	x := GetName()
	if x != "bcdapps1" {
		t.Error("Expected bcdapps got", x)
	}
}