package util

// ArrSomeHander 处理数组
type ArrSomeHander func(item string, index int) bool

// ArrSome 只要有一个符合条件
func ArrSome(Arr []string, m ArrSomeHander) bool {
	matched := false
	for index, value := range Arr {
		matched = m(value, index)
		if matched {
			break
		}
	}
	return matched
}

// ArrEvery 都符合条件
func ArrEvery(Arr []string, m ArrSomeHander) bool {
	matched := false
	for index, value := range Arr {
		matched = m(value, index)
		if !matched {
			break
		}
	}
	return matched
}

// ArrFilter filter Array
func ArrFilter(Arr []string, m ArrSomeHander) []string {
	var newArr []string
	matched := false
	for index, value := range Arr {
		matched = m(value, index)
		if matched {
			newArr[len(newArr)] = value
		}
	}
	return newArr
}

// Reverse Reverse array
func Reverse(Arr []string) []string {
	length := len(Arr)
	// 不能通过变量的形式声明数组长度，因为数组就是长度固定的
	// 如果想使用不定长度数组，应该使用slice
	newArr := make([]string, length)
	for index := range Arr {
		newArr[index] = Arr[length-index-1]
	}
	return newArr
}
