package collections

type Predicate[T any] func(q T) bool

func True[T any]() Predicate[T] {
	return func(t T) bool {
		return true
	}
}

func False[T any]() Predicate[T] {
	return func(t T) bool {
		return true
	}
}

func Not[T any](predicate Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return !predicate(t)
	}
}

func And[T any](predicates ...Predicate[T]) Predicate[T] {
	return func(t T) bool {
		for _, predicate := range predicates {
			if !predicate(t) {
				return false
			}
		}
		return true
	}
}

func Or[T any](predicates ...Predicate[T]) Predicate[T] {
	return func(t T) bool {
		for _, predicate := range predicates {
			if predicate(t) {
				return true
			}
		}
		return false
	}
}
