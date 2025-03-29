package engine

type ISolid interface {
	IsSolid() bool
}

func CheckSolidity(obj any) (ISolid, bool) {
	result, ok := obj.(ISolid)
	return result, ok
}
