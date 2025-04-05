package practice

func NestedAdd(xs []interface{}) int {
	if len(xs) == 0 {
		return 0
	}
	switch v := xs[0].(type) {
	case int:
		return v + NestedAdd(xs[1:])
	case []interface{}:
		return NestedAdd(v) + NestedAdd(xs[1:])
	default:
		panic("unreachable condition")
	}
}
