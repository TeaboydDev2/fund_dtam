package helper

import "strconv"

type Int interface {
	~int | ~int32 | ~int64
}

func StrToInt[T Int](str string, v T) T {
	switch any(v).(type) {
	case int:
		i, err := strconv.Atoi(str)
		if err != nil {
			return v
		}
		return any(i).(T)
	case int32:
		i, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return v
		}
		return T(int32(i))
	case int64:
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return v
		}
		return T(i)
	default:
		return v
	}
}

func StrToBool(str string) bool {

	b, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}

	return b
}
