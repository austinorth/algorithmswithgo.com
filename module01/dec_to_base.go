package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	var res string
	if dec != 0 {
		r := dec % base
		res = fmt.Sprintf("%X%s", r, res)
	} else {
		return res
	}
	return DecToBase(dec/base, base) + res
}
