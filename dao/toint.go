package dao

import "strconv"

func MustToInt32(s string) (i int32, err error) {
	ret, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	i = int32(ret)
	return i, nil
}
