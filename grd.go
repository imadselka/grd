package grd

type TryResult[T any] struct {
	result T
	err    error
}

func Try[T any](fn func() (T, error)) *TryResult[T] {
	res, err := fn()
	return &TryResult[T]{result: res, err: err}
}

func (t *TryResult[T]) Then(fn func(T) (T, error)) *TryResult[T] {
	if t.err != nil {
		return t
	}
	res, err := fn(t.result)
	return &TryResult[T]{result: res, err: err}
}

func (t *TryResult[T]) Catch(fn func(error) T) T {
	if t.err != nil {
		return fn(t.err)
	}
	return t.result
}

func (t *TryResult[T]) Finally(fn func()) *TryResult[T] {
	fn()
	return t
}
