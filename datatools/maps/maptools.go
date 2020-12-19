package maptools

// InMap returns the value of the given key in the map or false.
func InMap(key string, dict map[string]string) bool {
	if _, ok := dict[key]; ok {
		return true
	}
	return false
}
