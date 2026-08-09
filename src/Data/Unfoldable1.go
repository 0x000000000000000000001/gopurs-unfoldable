func Unfoldr1ArrayImpl(isNothing func(interface{}) bool, fromJust func(interface{}) interface{}, fst func(interface{}) interface{}, snd func(interface{}) interface{}, f func(interface{}) interface{}, b interface{}) []interface{} {
	var result []interface{}
	value := b
	for {
		tuple := f(value)
		result = append(result, fst(tuple))
		maybe := snd(tuple)
		if isNothing(maybe) { break }
		value = fromJust(maybe)
	}
	return result
}
