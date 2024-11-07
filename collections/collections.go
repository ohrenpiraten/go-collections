package collections

import "github.com/ohrenpiraten/go-collections/predicates"

func Map[T any, R any](collection []T, mapper func(T) R) (result []R) {
	result = make([]R, 0)
	for _, element := range collection {
		result = append(result, mapper(element))
	}
	return result
}

func MapError[T any, R any](collection []T, transform func(T) (R, error)) (result []R, _ error) {
	result = make([]R, 0)
	for _, element := range collection {
		if transformed, err := transform(element); err != nil {
			return result, err
		} else {
			result = append(result, transformed)
		}
	}
	return result, nil
}

func First[T any](collection []T, predicate predicates.Predicate[T]) (result T, found bool) {
	for _, element := range collection {
		if predicate(element) {
			return element, true
		}
	}
	return result, false
}

func Filter[T any](collection []T, predicate predicates.Predicate[T]) (result []T) {
	result = make([]T, 0)
	for _, element := range collection {
		if predicate(element) {
			result = append(result, element)
		}
	}
	return result
}

func Count[T any](collection []T, predicate predicates.Predicate[T]) int {
	return len(Filter(collection, predicate))
}

func AnyMatch[T any](collection []T, predicate predicates.Predicate[T]) bool {
	for _, element := range collection {
		if predicate(element) {
			return true
		}
	}
	return false
}

func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}

func ContainsAll[T comparable](container []T, elements []T) bool {
	if len(container) < len(elements) {
		return false
	}
	for _, element := range elements {
		if !Contains(container, element) {
			return false
		}
	}
	return true
}

func MutualContainment[T comparable](a []T, b []T) bool {
	return ContainsAll(a, b) && ContainsAll(b, a)
}
