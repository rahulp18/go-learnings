package stringutils

func Reverse(s string) string {
	runes := []rune(s)

	right := len(runes) - 1
	left := 0
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--

	}
	return string(runes)
}
