package common

import (
	"strconv"
)

func SanitizePhone(phone string) string {
	if len(phone) == 13 && phone[4] == '9' {
		phone = phone[0:4] + phone[5:]
	}
	return phone
}

func formatDecimal(value int64) string {
	// Convert the decimal part to a string
	decimalStr := strconv.FormatInt(value, 10)

	// Pad with leading zeros if necessary
	if len(decimalStr) == 1 {
		decimalStr = "0" + decimalStr
	}

	return decimalStr
}

func formatInteger(value int64) string {
	// Convert the integer part to a string
	integerStr := strconv.FormatInt(value, 10)

	// Format with thousands separator
	formatted := ""
	for i := len(integerStr) - 1; i >= 0; i-- {
		formatted = string(integerStr[i]) + formatted
		if (len(integerStr)-i)%3 == 0 && i != 0 {
			formatted = "." + formatted
		}
	}

	return formatted
}

func FormatCentsToBRL(cents int64) string {
	// Extract the integer and decimal parts
	integerPart := cents / 100
	decimalPart := cents % 100

	// Format the integer part with thousands separator
	formattedIntegerPart := formatInteger(integerPart)

	// Format the decimal part
	formattedDecimalPart := formatDecimal(decimalPart)

	// Join the integer and decimal parts with the comma separator
	formattedAmount := formattedIntegerPart + "," + formattedDecimalPart

	// Add the BRL symbol
	formattedAmount = "R$ " + formattedAmount

	return formattedAmount
}
