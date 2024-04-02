package utility

import "strconv"

func ConvertRollsToString(rolls []int) string {
	s := ""
	for _, v := range rolls {
		c := strconv.Itoa(v)
		s += c
		s += ", "
	}
	return s
}