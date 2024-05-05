package main

import (
	"fmt"
	"regexp"

	"github.com/biter777/countries"
)

const IBANMaxLength = 34

// International Bank Account Number.
type IBAN []byte

func NewIBAN(countryCode, bban string) IBAN {
	checkDigits := computeCheckDigits(countryCode, bban)

	var iban IBAN
	iban = append(iban, []byte(countryCode)...)
	iban = append(iban, fmt.Sprint(checkDigits)...)
	iban = append(iban, bban...)

	return iban
}

func FromString(s string) (IBAN, error) {
	iban := IBAN(s)
	if !iban.IsValid() {
		return nil, fmt.Errorf("invalid IBAN: %s", s)
	}

	return iban, nil
}

func (i IBAN) String() string {
	return string(i)
}

func (i IBAN) CountryCode() string {
	return string(i[:2])
}

func (i IBAN) CheckDigits() string {
	return string(i[2:4])
}

func (i IBAN) BBAN() string {
	return string(i[4:])
}

func (i IBAN) IsValid() bool {
	if len(i) > IBANMaxLength {
		return false
	}

	cc := countries.ByName(i.CountryCode())
	if !cc.IsValid() {
		return false
	}

	// IBAN must contain only alphanumeric characters.
	if !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(string(i)) {
		return false
	}

	for _, b := range i {
		if (b < '0' || b > '9') && (b < 'A' || b > 'Z') {
			return false
		}
	}

	var payload []byte
	payload = append(payload, []byte(i.BBAN())...)
	payload = append(payload, []byte(i.CountryCode())...)
	payload = append(payload, []byte(i.CheckDigits())...)

	digits := convertToDigits(payload)
	return mod97_10(digits) == 1
}

var genericLetterDigitMap = map[byte][]byte{
	'A': {1, 0}, 'B': {1, 1}, 'C': {1, 2}, 'D': {1, 3}, 'E': {1, 4}, 'F': {1, 5}, 'G': {1, 6},
	'H': {1, 7}, 'I': {1, 8}, 'J': {1, 9}, 'K': {2, 0}, 'L': {2, 1}, 'M': {2, 2}, 'N': {2, 3},
	'O': {2, 4}, 'P': {2, 5}, 'Q': {2, 6}, 'R': {2, 7}, 'S': {2, 8}, 'T': {2, 9}, 'U': {3, 0},
	'V': {3, 1}, 'W': {3, 2}, 'X': {3, 3}, 'Y': {3, 4}, 'Z': {3, 5},
}

func computeCheckDigits(countryCode string, bban string) string {
	var payload []byte
	payload = append(payload, []byte(bban)...)
	payload = append(payload, []byte(countryCode)...)
	payload = append(payload, []byte{'0', '0'}...)

	digits := convertToDigits(payload)
	checkDigits := 98 - mod97_10(digits)
	if checkDigits < 10 {
		return fmt.Sprintf("0%d", checkDigits)
	}

	return fmt.Sprintf("%d", checkDigits)
}

func convertToDigits(payload []byte) []byte {
	var digits []byte
	for _, c := range payload {
		if c >= 'A' && c <= 'Z' {
			digits = append(digits, genericLetterDigitMap[c]...)
			continue
		}

		digits = append(digits, c-'0')
	}

	return digits
}

func mod97_10(digits []byte) int {
	i, n := 0, 0
	for i < len(digits) {
		numNextDigits := 9
		if n > 0 {
			numNextDigits = 8
		}
		if n > 10 {
			numNextDigits = 7
		}

		for numNextDigits > 0 && i < len(digits) {
			n = n*10 + int(digits[i])
			i++
			numNextDigits--
		}

		n = n % 97
	}

	return n
}
