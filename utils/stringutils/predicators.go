package stringutils

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsEmpty(s string) bool {
	return s == ""
}

func IsNotEmpty(s string) bool {
	return s != ""
}
