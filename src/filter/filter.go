package filter

type Predicate func(int) bool

func Filter(list []int, predicate Predicate) []int {
	var result []int
	for _, e := range list {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}
