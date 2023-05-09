package common

import "regexp"

func RemoveNonNumeric(input string) string {
	reg, _ := regexp.Compile("[^0-9]+")
	return reg.ReplaceAllString(input, "")
}
