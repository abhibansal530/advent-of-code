package utils

func MapHasValue[M ~map[K]V, K, V comparable](m M, toFind V) bool {
	for _, v := range m {
		if v == toFind {
			return true
		}
	}
	return false
}
