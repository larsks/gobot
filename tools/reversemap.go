package tools

// ReverseMap returns a map with keys and values reversed.
func ReverseMap[K comparable, V comparable](inputMap map[K]V) map[V]K {
	reversedMap := make(map[V]K, len(inputMap))
	for k, v := range inputMap {
		reversedMap[v] = k
	}
	return reversedMap
}
