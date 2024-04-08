package utility

import "strconv"

func ConvertRollsToString(rolls []Roll) string {
	s := ""
	for _, v := range rolls {
		c := strconv.Itoa(v.roll)
		s += c
		s += ", "
	}
	return s
}