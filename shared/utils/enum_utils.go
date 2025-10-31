package utils

type Enum interface {
	comparable
}

func StringValue[T Enum](e T, values map[T]string) string {
	return values[e]
}

func EnumValue[T Enum](s string, values map[T]string, noneValue T) T {
	for k, v := range values {
		if v == s {
			return k
		}
	}

	return noneValue
}
