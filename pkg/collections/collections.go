package collections

func MapValues[K comparable, V any](dataMap map[K]V) []V {
	values := make([]V, 0, len(dataMap))

	for _, d := range dataMap {
		values = append(values, d)
	}

	return values
}

func MapKeys[K comparable, V any](dataMap map[K]V) []K {
	keys := make([]K, 0, len(dataMap))

	for k, _ := range dataMap {
		keys = append(keys, k)
	}

	return keys
}
