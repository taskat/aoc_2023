package maps

import "aoc_2023/utils/types"

func All[KEY comparable, VALUE any](m map[KEY]VALUE, predicate func(KEY, VALUE) bool) bool {
	for key, value := range m {
		if !predicate(key, value) {
			return false
		}
	}
	return true
}

func Any[KEY comparable, VALUE any](m map[KEY]VALUE, predicate func(KEY, VALUE) bool) bool {
	for key, value := range m {
		if predicate(key, value) {
			return true
		}
	}
	return false
}

func Count[KEY comparable, VALUE any](m map[KEY]VALUE) int {
	return len(m)
}

func Contains[KEY comparable, VALUE any](m map[KEY]VALUE, key KEY) bool {
	_, ok := m[key]
	return ok
}

func Filter[KEY comparable, VALUE any](m map[KEY]VALUE, predicate func(KEY, VALUE) bool) map[KEY]VALUE {
	result := make(map[KEY]VALUE)
	for key, value := range m {
		if predicate(key, value) {
			result[key] = value
		}
	}
	return result
}

func Find[KEY comparable, VALUE any](m map[KEY]VALUE, predicate func(KEY, VALUE) bool) (KEY, VALUE, bool) {
	for key, value := range m {
		if predicate(key, value) {
			return key, value, true
		}
	}
	var zeroKey KEY
	var zeroValue VALUE
	return zeroKey, zeroValue, false
}

func FindKey[KEY comparable, VALUE any](m map[KEY]VALUE, predicate func(KEY, VALUE) bool) (KEY, bool) {
	key, _, ok := Find(m, predicate)
	return key, ok
}

func FindValue[KEY comparable, VALUE any](m map[KEY]VALUE, predicate func(KEY, VALUE) bool) (VALUE, bool) {
	_, value, ok := Find(m, predicate)
	return value, ok
}

func ForEach[KEY comparable, VALUE any](m map[KEY]VALUE, action func(KEY, VALUE)) {
	for key, value := range m {
		action(key, value)
	}
}

func Keys[KEY comparable, VALUE any](m map[KEY]VALUE) []KEY {
	return MapToArray(m, func(k KEY, _ VALUE) KEY { return k })
}

func Map[KEY1, KEY2 comparable, VALUE1, VALUE2 any](m map[KEY1]VALUE1, mapper func(KEY1, VALUE1) (KEY2, VALUE2)) map[KEY2]VALUE2 {
	result := make(map[KEY2]VALUE2)
	for key, value := range m {
		newKey, newValue := mapper(key, value)
		result[newKey] = newValue
	}
	return result
}

func MapToArray[KEY comparable, VALUE1, VALUE2 any](m map[KEY]VALUE1, mapper func(KEY, VALUE1) VALUE2) []VALUE2 {
	result := make([]VALUE2, 0, len(m))
	for key, value := range m {
		result = append(result, mapper(key, value))
	}
	return result
}

func Max[KEY comparable, VALUE types.Number](m map[KEY]VALUE) VALUE {
	if len(m) == 0 {
		panic("Cannot find max of empty map")
	}
	var result VALUE
	for _, value := range m {
		if value > result {
			result = value
		}
	}
	return result
}

func Min[KEY comparable, VALUE types.Number](m map[KEY]VALUE) VALUE {
	if len(m) == 0 {
		panic("Cannot find min of empty map")
	}
	var result VALUE
	for _, value := range m {
		if value < result {
			result = value
		}
	}
	return result
}

func None[KEY comparable, VALUE any](m map[KEY]VALUE, predicate func(KEY, VALUE) bool) bool {
	for key, value := range m {
		if predicate(key, value) {
			return false
		}
	}
	return true
}

func Sum[KEY comparable, VALUE types.Summable](m map[KEY]VALUE) VALUE {
	var result VALUE
	for _, value := range m {
		result += value
	}
	return result
}

func Product[KEY comparable, VALUE types.Number](m map[KEY]VALUE) VALUE {
	var result VALUE
	for _, value := range m {
		result *= value
	}
	return result
}

func Values[KEY comparable, VALUE any](m map[KEY]VALUE) []VALUE {
	return MapToArray(m, func(_ KEY, v VALUE) VALUE { return v })
}
