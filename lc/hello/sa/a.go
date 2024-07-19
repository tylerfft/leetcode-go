package main

// (i(u(wa)e)h)

// huawei

// a(bcdefghijkl(mno)p)q
// a p mno lkjihgfedcb q

// abcdefghijklmnopq

func Resolve(str string) string {
	d := []byte(str)
	prev, back := 0, len(d)-1
	for prev < back {
		for prev < back && d[prev] != '(' && d[prev] != ')' {
			prev++
		}
		for prev < back && d[back] != '(' && d[back] != ')' {
			back--
		}
		prev++
		back--
		reverse(d, prev, back)
	}
	rst := []byte{}
	for i := 0; i < len(d); i++ {
		if d[i] != '(' && d[i] != ')' {
			rst = append(rst, d[i])
		}
	}
	return string(rst)
}
func reverse(d []byte, l, r int) {
	for l < r {
		d[l], d[r] = d[r], d[l]
		l++
		r--
	}
}

// 反转每对括号间的子串
// 题目描述
// 给出一个字符串 s（仅含有小写英文字母和括号）。
// 请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。

// 注意，您的结果中不应包含任何括号。

// 输入
// 输入为一行带有括号的字符串(只包含英文小写字母和左右小括号)
// 最大长度不会超过10000个字符

// 输出
// 反转括号内字符串并输出（只有英文小写字母）
