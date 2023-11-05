package misc

import "cmp"

// Clamp 取值上下限，限制取值返回
func Clamp[T cmp.Ordered](val, min, max T) T {
	if val > max {
		val = max
	}
	if val < min {
		val = min
	}
	return val
}
