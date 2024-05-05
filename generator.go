package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	CountryCodeNetherlands = "NL"
	CountryCodeGermany     = "DE"
	CountryCodeAustria     = "AT"
	CountryCodeSwitzerland = "CH"
	CountryCodeBelgium     = "BE"
	CountryCodeIreland     = "IE"
	CountryCodeSpain       = "ES"
	CountryCodeFrance      = "FR"
	CountryCodeItaly       = "IT"

	IBANLengthNetherlands = 18
	IBANLengthGermany     = 22
	IBANLengthAustria     = 20
	IBANLengthSwitzerland = 21
	IBANLengthBelgium     = 16
	IBANLengthIreland     = 22
	IBANLengthSpain       = 24
	IBANLengthFrance      = 27
	IBANLengthItaly       = 27
)

var Generators map[string]Generator = map[string]Generator{
	CountryCodeNetherlands: GenericGenerator{CountryCode: CountryCodeNetherlands, Banks: dutchBanks, AccountNumberLength: IBANLengthNetherlands},
	CountryCodeGermany:     GenericGenerator{CountryCode: CountryCodeGermany, Banks: germanBanks, AccountNumberLength: IBANLengthGermany},
	CountryCodeIreland:     GenericGenerator{CountryCode: CountryCodeIreland, Banks: irishBanks, AccountNumberLength: IBANLengthIreland},
	CountryCodeAustria:     GenericGenerator{CountryCode: CountryCodeAustria, Banks: austrianBanks, AccountNumberLength: IBANLengthAustria},
	CountryCodeSwitzerland: GenericGenerator{CountryCode: CountryCodeSwitzerland, Banks: swissBanks, AccountNumberLength: IBANLengthSwitzerland},

	CountryCodeBelgium: BelgiumGenerator{},
	CountryCodeSpain:   SpainGenerator{},
	CountryCodeFrance:  FranceGenerator{},
	CountryCodeItaly:   ItalyGenerator{},
}

const (
	genericAccountNumberCharset = "0123456789"

	ibanChecksumLength            = 2
	belgiumNationalChecksumLength = 2
	spainNationalChecksumLength   = 2
	franceNationalChecksumLength  = 2
	italyNationalChecksumLength   = 1
)

type Generator interface {
	Generate() IBAN
}

// DutchIBANGenerator generates IBANs for countries that do not have
// a checksum for the account number and only use the IBAN checksum.
type GenericGenerator struct {
	CountryCode         string
	Banks               []bank
	AccountNumberLength int
}

func (g GenericGenerator) Generate() IBAN {
	bank := g.Banks[seededRand.Intn(len(g.Banks))]
	accountNumberLength := g.AccountNumberLength - (ibanChecksumLength + len(g.CountryCode) + len(bank.code))
	accountNumber := randomStringFrom(genericAccountNumberCharset, accountNumberLength)

	bban := fmt.Sprintf("%s%s", bank.code, string(accountNumber))
	return NewIBAN(g.CountryCode, bban)
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

func randomStringFrom(charset string, length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

type BelgiumGenerator struct{}

func (g BelgiumGenerator) Generate() IBAN {
	randomBank := belgianBanks[seededRand.Intn(len(belgianBanks))]

	accountNumberLength := IBANLengthBelgium - (len(CountryCodeBelgium) +
		len(randomBank.code) + ibanChecksumLength + belgiumNationalChecksumLength)
	accountNumber := randomStringFrom(genericAccountNumberCharset, accountNumberLength)

	nationalChecksum := g.computeNationalChecksum(randomBank.code, string(accountNumber))

	bban := fmt.Sprintf("%s%s%s", randomBank.code, accountNumber, nationalChecksum)
	return NewIBAN(CountryCodeBelgium, bban)
}

func (g BelgiumGenerator) computeNationalChecksum(bankCode, accountNumber string) string {
	var payload []byte
	payload = append(payload, []byte(bankCode)...)
	payload = append(payload, []byte(accountNumber)...)

	digits := convertToDigits(payload)
	checkDigits := mod97_10(digits)
	if checkDigits == 0 {
		checkDigits = 97
	}

	if checkDigits < 10 {
		return fmt.Sprintf("0%d", checkDigits)
	}

	return fmt.Sprintf("%d", checkDigits)
}

type SpainGenerator struct{}

func (g SpainGenerator) Generate() IBAN {
	randomBank := spanishBanks[seededRand.Intn(len(spanishBanks))]
	accountNumberLength := IBANLengthSpain - (len(CountryCodeSpain) +
		len(randomBank.code) + ibanChecksumLength + spainNationalChecksumLength)

	accountNumber := randomStringFrom(genericAccountNumberCharset, accountNumberLength)
	nationalChecksum := g.computeNationalChecksum(randomBank.code, string(accountNumber))

	bban := fmt.Sprintf("%s%s%s", randomBank.code, nationalChecksum, accountNumber)
	return NewIBAN(CountryCodeSpain, bban)
}

func (g SpainGenerator) computeNationalChecksum(bankCode, accountNumber string) string {
	var firstDigit int
	weights := []int{4, 8, 5, 10, 9, 7, 3, 6}
	bankCodeDigits := convertToDigits([]byte(bankCode))
	for i := 0; i < len(bankCodeDigits); i++ {
		firstDigit += int(bankCodeDigits[i]) * weights[i]
	}
	firstDigit = firstDigit % 11
	if firstDigit > 1 {
		firstDigit = 11 - firstDigit
	}

	var secondDigit int
	weights = []int{1, 2, 4, 8, 5, 10, 9, 7, 3, 6}
	accountNumberDigits := convertToDigits([]byte(accountNumber))
	for i := 0; i < len(accountNumberDigits); i++ {
		secondDigit += int(accountNumberDigits[i]) * weights[i]
	}
	secondDigit = secondDigit % 11
	if secondDigit > 1 {
		secondDigit = 11 - secondDigit
	}

	return fmt.Sprintf("%d%d", firstDigit, secondDigit)
}

type FranceGenerator struct{}

var franceAccountNumberLetterToDigitMapping = map[byte]int{
	'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8, 'I': 9,
	'J': 1, 'K': 2, 'L': 3, 'M': 4, 'N': 5, 'O': 6, 'P': 7, 'Q': 8, 'R': 9,
	'S': 2, 'T': 3, 'U': 4, 'V': 5, 'W': 6, 'X': 7, 'Y': 8, 'Z': 9,
}

func (g FranceGenerator) Generate() IBAN {
	bank := frenchBanks[seededRand.Intn(len(frenchBanks))]
	branch := bank.branchCodes[seededRand.Intn(len(bank.branchCodes))]

	accountNumberLength := IBANLengthFrance - (len(CountryCodeFrance) +
		len(bank.code) + len(branch) + ibanChecksumLength + franceNationalChecksumLength)

	accountNumber := randomStringFrom(genericAccountNumberCharset, accountNumberLength)
	checksum := g.computeNationalChecksum(bank.code, branch, string(accountNumber))

	bban := fmt.Sprintf("%s%s%s%s", bank.code, branch, accountNumber, checksum)
	return NewIBAN(CountryCodeFrance, bban)
}

func (g FranceGenerator) computeNationalChecksum(bankCode, bankBranch, accountNumber string) string {
	var bankCodeNumber uint64
	for _, digit := range g.convertToDigits([]byte(bankCode)) {
		bankCodeNumber = bankCodeNumber*10 + uint64(digit)
	}
	bankCodeNumber = bankCodeNumber * 89

	var bankBranchNumber uint64
	for _, digit := range g.convertToDigits([]byte(bankBranch)) {
		bankBranchNumber = bankBranchNumber*10 + uint64(digit)
	}
	bankBranchNumber = bankBranchNumber * 15

	var accountNumberNumber uint64
	for _, digit := range g.convertToDigits([]byte(accountNumber)) {
		accountNumberNumber = accountNumberNumber*10 + uint64(digit)
	}
	accountNumberNumber = accountNumberNumber * 3

	checkDigits := 97 - (bankCodeNumber+bankBranchNumber+accountNumberNumber)%97
	if checkDigits < 10 {
		return fmt.Sprintf("0%d", checkDigits)
	}

	return fmt.Sprintf("%02d", checkDigits)
}

func (g FranceGenerator) convertToDigits(payload []byte) []int {
	var digits []int
	for _, char := range payload {
		if char >= '0' && char <= '9' {
			digits = append(digits, int(char-'0'))
			continue
		}

		digits = append(digits, franceAccountNumberLetterToDigitMapping[char])
	}

	return digits
}

type ItalyGenerator struct{}

var italyAccountNumberOddLetterToDigitMapping = map[byte]int{
	'0': 1, '1': 0, '2': 5, '3': 7, '4': 9, '5': 13, '6': 15, '7': 17, '8': 19, '9': 21,
	'A': 1, 'B': 0, 'C': 5, 'D': 7, 'E': 9, 'F': 13, 'G': 15, 'H': 17, 'I': 19, 'J': 21,
	'K': 2, 'L': 4, 'M': 18, 'N': 20, 'O': 11, 'P': 3, 'Q': 6, 'R': 8, 'S': 12, 'T': 14,
	'U': 16, 'V': 10, 'W': 22, 'X': 25, 'Y': 24, 'Z': 23,
}

var italyAccountNumberEvenLetterToDigitMapping = map[byte]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
	'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8, 'J': 9,
	'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15, 'Q': 16, 'R': 17, 'S': 18,
	'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23, 'Y': 24, 'Z': 25,
}

func (g ItalyGenerator) Generate() IBAN {
	var (
		bank   = italianBanks[seededRand.Intn(len(italianBanks))]
		branch = bank.branchCodes[seededRand.Intn(len(bank.branchCodes))]
	)

	accountNumberLength := IBANLengthItaly - (len(CountryCodeItaly) +
		len(bank.code) + len(branch) + ibanChecksumLength + italyNationalChecksumLength)
	accountNumber := randomStringFrom(genericAccountNumberCharset, accountNumberLength)

	checkChar := g.computeNationalChecksum(bank.code, branch, string(accountNumber))
	bban := fmt.Sprintf("%s%s%s%s", checkChar, bank.code, branch, accountNumber)

	return NewIBAN(CountryCodeItaly, bban)
}

func (g ItalyGenerator) computeNationalChecksum(bankCode, bankBranch, accountNumber string) string {
	var payload []byte
	payload = append(payload, []byte(bankCode)...)
	payload = append(payload, []byte(bankBranch)...)
	payload = append(payload, []byte(accountNumber)...)

	var sum int
	for _, value := range g.convertToValues(payload) {
		sum += value
	}

	return string(byte(sum%26 + 0x41))
}

func (g ItalyGenerator) convertToValues(payload []byte) []int {
	var digits []int
	for i, char := range payload {
		if (i+1)%2 == 0 {
			digits = append(digits, italyAccountNumberEvenLetterToDigitMapping[char])
			continue
		}

		digits = append(digits, italyAccountNumberOddLetterToDigitMapping[char])
	}

	return digits
}
