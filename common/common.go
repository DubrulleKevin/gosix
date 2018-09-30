package common

// PanicIfError panics if error "e" is not nil.
func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

/*// StringSliceContains returns true if string "s" is found in strings' slice "ss", and false otherwise.
func StringSliceContains(ss []string, s string) bool {
	for _, el := range ss {
		if el == s {
			return true
		}
	}

	return false
}
*/