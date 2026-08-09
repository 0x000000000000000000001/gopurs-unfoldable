func UnfoldrArrayImpl(isNothing func(interface{}) bool, fromJust func(interface{}) interface{}, fst func(interface{}) interface{}, snd func(interface{}) interface{}, f func(interface{}) interface{}, b interface{}) []interface{} {
	var result []interface{}
	value := b
	for {
		maybe := f(value)
		if isNothing(maybe) {
			break
		}
		tuple := fromJust(maybe)
		result = append(result, fst(tuple))
		value = snd(tuple)
	}
	if result == nil {
		return []interface{}{}
	}
	return result
}
