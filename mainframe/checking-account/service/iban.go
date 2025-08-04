package service

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GenerateNewIBAN(abi string, cab string, accountId string) string {
	// Generate a new account code
	accountCode := "000000" + fmt.Sprintf("%f", rand.Float64())[2:8]

	// Generate an IBAN from given values
	ibban := generateItalianBBAN(abi, cab, accountCode)
	iban := generateEuropeanIBAN("IT", ibban)

	return string(iban)
}

func generateItalianBBAN(abi string, cab string, accountCode string) []byte {
	// Char convetion values
	var oddBBANValues = []int16{1, 0, 5, 7, 9, 13, 15, 17, 19, 21, 1, 0, 5, 7, 9, 13, 15, 17,
		19, 21, 2, 4, 8, 18, 20, 11, 3, 6, 8, 12, 14, 16, 10, 22, 25, 24, 23}
	var evenBBANValues = []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7,
		8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

	// Construct Italian BBAN used to compute CIN
	bban := fmt.Appendf(nil, " %s%s%s", abi, cab, accountCode)

	var sumOdd int16 = 0
	var sumEven int16 = 0

	// Sum each convertion value depending on oddness of the char index
	for i := 1; i < len(bban); i++ {
		if i%2 == 1 {
			sumOdd += oddBBANValues[bban[i]-48]
		} else {
			sumEven += evenBBANValues[bban[i]-48]
		}
	}

	// Convert the sum back to the corresponding ascii char
	bban[0] = byte((sumOdd+sumEven)%26 + 65)

	return bban
}

func generateEuropeanIBAN(nation string, bban []byte) []byte {
	// Construct European BBAN used to compute CIN
	ebban := fmt.Appendf(bban, "%s00", nation)

	// Convert all characters to numbers using their ASCII and convertion table
	var numericEbban []byte
	for i := range ebban {
		if ebban[i] >= 65 {
			numericEbban = fmt.Appendf(numericEbban, "%d", ebban[i]-55)
		} else {
			numericEbban = append(numericEbban, ebban[i])
		}
	}

	// Compute CIN using MOD97
	var cin int
	for {
		if len(numericEbban) > 8 {
			value, _ := strconv.Atoi(string(numericEbban[:8]))
			mod := value % 97
			numericEbban = numericEbban[8:]
			numericEbban = append([]byte(strconv.Itoa(mod)), numericEbban...)
		} else {
			value, _ := strconv.Atoi(string(numericEbban))
			cin = 98 - value%97
			break
		}
	}

	return fmt.Appendf(nil, "%s%02d%s", nation, cin, bban)
}
