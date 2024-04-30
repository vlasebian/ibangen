package iban

import (
	"testing"
)

func TestValidate(t *testing.T) {
	iban := IBAN("GB82WEST12345698765432")
	if !iban.IsValid() {
		t.Errorf("expected valid, got invalid")
	}

	iban = IBAN("NL91ABNA0417164301")
	if iban.IsValid() {
		t.Errorf("expected invalid, got valid")
	}
}
