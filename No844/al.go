package no844

// 给定 s 和 t 两个字符串，当它们分别被输入到空白的文本编辑器后，如果两者相等，返回 true 。# 代表退格字符。

// 注意：如果对空文本输入退格字符，文本继续为空。

// 示例 1：

// 输入：s = "ab#c", t = "ad#c"
// 输出：true
// 解释：s 和 t 都会变成 "ac"。
// 示例 2：

// 输入：s = "ab##", t = "c#d#"
// 输出：true
// 解释：s 和 t 都会变成 ""。
// 示例 3：

// 输入：s = "a#c", t = "b"
// 输出：false
// 解释：s 会变成 "c"，但 t 仍然是 "b"。

func back(s string) string {
	res := []rune{}
	for _, r := range s {
		if r != '#' {
			res = append(res, r)
		} else if len(res) > 0 {
			// 把最后一个去掉
			res = res[:len(res)-1]
		}
	}
	return string(res)
}

func backspaceCompare(s string, t string) bool {
	return back(s) == back(t)
}
