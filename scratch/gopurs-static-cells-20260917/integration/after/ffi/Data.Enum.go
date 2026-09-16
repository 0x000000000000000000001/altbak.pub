func FromCharCode(c int) string { return string(rune(c)) }
func ToCharCode(c string) int { return int([]rune(c)[0]) }
