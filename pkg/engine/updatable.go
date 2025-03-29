package engine

type IUpdatable interface {
	Update(...any) error
}

func CheckUpdatable(obj any) (IUpdatable, bool) {
	result, ok := obj.(IUpdatable)
	return result, ok
}
