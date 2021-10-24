package filter

type Predicate[T any] func(T) bool

func Filter[T any](list []T, predicate Predicate[T]) (result []T) {
	for _, e := range list {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return
}
