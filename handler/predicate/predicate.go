package predicate

type Predicate[T any] interface {
	Apply(t T) bool

	And(other Predicate[T]) Predicate[T]
}

type DefaultPredicate[T any] struct {
	Delegate Predicate[T]
}

func (p *DefaultPredicate[T]) Apply(t T) bool {
	return p.Delegate.Apply(t)
}
func (p *DefaultPredicate[T]) And(other Predicate[T]) Predicate[T] {
	return &AndPredicate[T]{
		Left:  p,
		Right: other,
	}
}

type AndPredicate[T any] struct {
	DefaultPredicate[T]
	Left  Predicate[T]
	Right Predicate[T]
}

func (p *AndPredicate[T]) Apply(t T) bool {
	return func() bool {
		//如果是true,执行右边
		if p.Left.Apply(t) {
			return p.Right.Apply(t)
		}
		//如果是false,直接返回
		return false
	}()
}
