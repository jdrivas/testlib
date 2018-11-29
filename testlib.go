package testlib

// Reverse returns its argument string reversed run-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Hello returns a cannonical hello string.
func Hello(s string) string {
	return "Hello World"
}
