package validation

import "regexp"

var uzbPhoneRegex = regexp.MustCompile(`^[+]{1}99{1}[0-9]{10}$`)
var tgPhoneRegex = regexp.MustCompile(`^\+992\d{9}$`)

// IsPhoneValid validates phone number for Uzbekistan
func IsPhoneValid(p string) bool {
	isUzb := uzbPhoneRegex.MatchString(p)
	if isUzb {
		return true
	}

	return tgPhoneRegex.MatchString(p)
}
