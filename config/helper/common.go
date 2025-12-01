package config_helper

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ParseString[T any](s string, v T) T {
	switch any(v).(type) {
	case int:
		i, err := strconv.Atoi(s)
		if err != nil {
			return v
		}
		return any(i).(T)
	case bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return v
		}
		return any(b).(T)
	case float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return v
		}
		return any(f).(T)
	case string:
		return any(s).(T)
	default:
		return v
	}
}
