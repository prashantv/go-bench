package main

// MapOfMap is a map that has anotehr map as a value.
type MapOfMap map[string]map[string]interface{}

// GetWithOkCheck checks whether the first key exists using the v, ok lookup return.
func (m MapOfMap) GetWithOkCheck(k1, k2 string) interface{} {
	if innerMap, ok := m[k1]; ok {
		return innerMap[k2]
	}
	return nil
}

// GetWithCheck checks whether the first key exists by checking if the value is nil.
func (m MapOfMap) GetWithCheck(k1, k2 string) interface{} {
	if innerMap := m[k1]; innerMap != nil {
		return innerMap[k2]
	}
	return nil
}

// Get does not do any checks and relies on nil map lookup.
func (m MapOfMap) Get(k1, k2 string) interface{} {
	return m[k1][k2]
}
