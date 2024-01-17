package predicate

type Predicate[T any] interface {
	Apply(t T) bool
	//And(other Predicate[T]) Predicate[T]
}
