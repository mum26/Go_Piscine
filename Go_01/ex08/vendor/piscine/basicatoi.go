package piscine

func StrLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func BasicAtoi(s string) int {
	runes := []rune(s)
	len := StrLen(s)
	num := 0
	i := 0
	for i < len && '0' <= runes[i] && runes[i] <= '9' {
		num = num * 10 + int(runes[i] - '0')
		i++
	}
	return num
}