package common

import "fmt"

func SanitizePhone(phone string) string {
	if len(phone) == 13 && phone[4] == '9' {
		fmt.Println(phone)
		fmt.Println(phone[0:4])
		fmt.Println(phone[5:])
		fmt.Println(phone[0:4] + phone[5:])
		phone = phone[0:4] + phone[5:]
	}
	return phone
}
