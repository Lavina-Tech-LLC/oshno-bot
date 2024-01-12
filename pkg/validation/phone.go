package validation

import "regexp"

var uzbPhoneRegex = regexp.MustCompile(`^[+]{1}99{1}[0-9]{10}$`)

// IsPhoneValid validates phone number for Uzbekistan
func IsPhoneValid(p string) bool {
	return uzbPhoneRegex.MatchString(p)
}
