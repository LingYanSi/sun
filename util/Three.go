package util

// Three 三则运算
func Three(b bool, trueValue interface{}, falseValue interface{}) interface{} {
	if b {
		return trueValue
	}
	return falseValue
}
