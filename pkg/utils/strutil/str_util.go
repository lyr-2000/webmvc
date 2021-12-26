package strutil

import "strconv"

func IsNotBlank(s string) bool {
	var res bool
	for _, b := range s {
		if b == ' ' || b == '\r' || b == '\n' || b == '\t' {
			continue
		}
		res = true
		break
	}
	return res
}

func ContainsAlpha(s string) bool {
	for _, b := range s {
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			return true
		}
	}
	return false
}




func ParseInt(msg string) int {
	val, _ := strconv.Atoi(msg)
	return val
}
func ParseInt64(s string) int64 {
	parseInt, _ := strconv.ParseInt(s, 10, 64)
	return parseInt
}
func ParseFloat64(s string) float64 {
	var a,_ = strconv.ParseFloat(s,64)
	return a
}
