package common

import "regexp"

func FormatPhone(input string) string {
	reg, _ := regexp.Compile("[^0-9]+")
	return reg.ReplaceAllString(input, "")
}
