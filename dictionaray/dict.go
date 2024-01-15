package dictionaray

import "github.com/ohrenpiraten/go-collections/predicates"

func Values[K comparable, T any](collection map[K]T) (result []T) {
	result = make([]T, 0)
	for _, element := range collection {
		result = append(result, element)
	}
	return result
}

func Keys[K comparable, T any](collection map[K]T) (result []K) {
	result = make([]K, 0)
	for key, _ := range collection {
		result = append(result, key)
	}
	return result
}

func FilterValues[K comparable, T any](collection map[K]T, predicate predicates.Predicate[T]) (result []T) {
	result = make([]T, 0)
	for _, element := range collection {
		if predicate(element) {
			result = append(result, element)
		}
	}
	return result
}
