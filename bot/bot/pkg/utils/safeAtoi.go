package utils

import "strconv"

func SafeAtoi(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return res
}
