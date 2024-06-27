package sliceutil

import "slices"

type number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// ToMapByKey スライスをmapに変換する
func ToMapByKey[TGroupKey comparable, TValue any](base []TValue, keyFunc func(TValue) TGroupKey) map[TGroupKey]TValue {
	result := make(map[TGroupKey]TValue)
	for _, v := range base {
		result[keyFunc(v)] = v
	}
	return result
}

// ConvertNumberSlice number型のスライスの型を変換する
func ConvertNumberSlice[T1, T2 number](input []T1) []T2 {
	output := make([]T2, len(input))
	for i, v := range input {
		output[i] = T2(v)
	}
	return output
}

// ConvertSlice スライスの型を変換する
func ConvertSlice[TValue any, TOutput any](base []TValue, keyFunc func(TValue) TOutput) []TOutput {
	result := make([]TOutput, len(base))
	for i, v := range base {
		result[i] = keyFunc(v)
	}
	return result
}

// ConvertUniqueSlice スライスの重複を削除する
func ConvertUniqueSlice[T comparable](base []T) []T {
	result := make([]T, len(base))
	for _, v := range base {
		if slices.Contains(result, v) {
			continue
		}
		result = append(result, v)
	}
	return result
}
