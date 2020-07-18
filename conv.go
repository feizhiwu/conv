package conv

import (
	"strconv"
)

func FormatInt(num int64, base, bitSize int) string {
	var max int64 = 1<<bitSize - 1
	if num > max {
		return strconv.FormatInt(max, base)
	}
	//负数取补码
	if num < 0 {
		num = max + num + 1
		if num < 0 {
			num = 0
		}
	}
	return fill(strconv.FormatInt(num, base), strconv.FormatInt(max, base))
}

func fill(s, l string) string {
	if len(s) < len(l) {
		n := len(l) - len(s)
		for i := 0; i < n; i++ {
			s = "0" + s
		}
	}
	return s
}
