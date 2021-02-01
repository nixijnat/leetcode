package main

func removeInvalidParentheses(s string) []string {
	// 应该移除的左括号和右括号数量
	leftRm, rightRm := 0, 0
	for i := range s {
		switch s[i] {
		case '(':
			leftRm++ //
		case ')':
			if leftRm == 0 {
				rightRm++ // 不合法，需要删除此括号
			} else {
				leftRm-- // 可以匹配时
			}
		}
	}
	resMap := make(map[string]struct{}, 8)
	helper(s, 0, 0, leftRm, rightRm, "", resMap)
	res := make([]string, 0, 8)
	for k := range resMap {
		res = append(res, k)
	}
	return res
}

// s 原字符串
// index 处理的字符
// count 用来标识 res 是否合法，增加一个 ( 则加一，增加一个 ) 则减一，最终 count==0 则正常。
// leftRm 能够移除的左括号数量
// rightRm 能够移除的右括号数量
// res 已经添加的括号前缀
// resMap 结果集
func helper(s string, index int, count, leftRm, rightRm int, res string, resMap map[string]struct{}) {
	if index == len(s) {
		// 结束
		if count == 0 { // 如果最终的字符串正常，则放入结果
			resMap[res] = struct{}{}
		}
		return
	}
	// 统计相同的字符个数，进行相同处理
	same := 1
	for ; index+same < len(s) && s[index] == s[index+same]; same++ {
	}
	// switch 语句为 移除 index 操作
	switch s[index] {
	case '(':
		if leftRm < same { // 校正要移除的数量
			same = leftRm
		}
		if leftRm > 0 {
			// 当还能去除左括号
			helper(s, index+same, count, leftRm-same, rightRm, res, resMap)
		}
		count++
	case ')':
		if rightRm < same {
			same = rightRm
		}
		if rightRm > 0 {
			// 当还能去除右括号
			helper(s, index+same, count, leftRm, rightRm-same, res, resMap)
		}
		count--
	}
	// 不去除 index 的字符
	// 不能放在 switch 中，不然会遗漏掉 default
	// == 0 表示括号已经正常匹配，可继续寻找剩下括号
	// > 0 表示左括号多余右括号，可以继续寻找
	// < 0 表示右括号多余左括号，已经不正常了，不能继续递归了
	if count >= 0 {
		helper(s, index+1, count, leftRm, rightRm, res+s[index:index+1], resMap)
	}
}
