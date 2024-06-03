package backspacecompare

// gen 使用栈的特点
func gen(s string) string {
	res := []rune{}

	for _, r := range s {
		if r != '#' {
			res = append(res, r)
		} else if len(res) > 0 {
			res = res[:len(res)-1]
		}

	}

	return string(res)
}

// backSpaceCompare [844] 比较含退格的字符串
func backSpaceCompare(s string, t string) bool {
	return gen(s) == gen(t)
}
