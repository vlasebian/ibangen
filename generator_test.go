package main

import "testing"

func TestSpainGenerator(t *testing.T) {
	g := SpainGenerator{}
	checksum := g.computeNationalChecksum("21000418", "4910097123")
	if len(checksum) != 2 {
		t.Errorf("Expected checksum to have length 2, got %d", len(checksum))
	}
}
