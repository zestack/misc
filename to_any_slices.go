package misc

// ToAnySlices 将指定的类型切片转化成任意类型切片
func ToAnySlices[E any](slices []E) []any {
	result := make([]any, len(slices))
	for i, slice := range slices {
		result[i] = slice
	}
	return result
}
