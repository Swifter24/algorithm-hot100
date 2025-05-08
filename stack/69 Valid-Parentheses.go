package stack

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//每个右括号都有一个对应的相同类型的左括号。
//
//示例 1：
//输入：s = "()"
//输出：true
//
//示例 2：
//输入：s = "()[]{}"
//输出：true
//
//示例 3：
//输入：s = "(]"
//输出：false
//
//示例 4：
//输入：s = "([])"
//输出：true

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	mp := map[rune]rune{'(': ')', '[': ']', '{': '}'}
	st := []rune{}
	for _, ch := range s {
		if v := mp[ch]; v > 0 {
			st = append(st, v)
		} else {
			if len(st) == 0 || st[len(st)-1] != ch {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}
