package conv

import "strconv"

func FormatInt(num int64, base, bitSize int) string {
	var max int64 = 1<<bitSize - 1
	if num > max {
		return strconv.FormatInt(max, base)
	}
	//负数取补码
	if num < 0 {
		num = max + num + 1
	}
	return strconv.FormatInt(num, base)
}
