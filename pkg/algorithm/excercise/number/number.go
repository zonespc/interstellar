package number

import (
	"fmt"
	"math"
	"strconv"
)

//ReverseDigits 翻转32位整型数各位数字
func ReverseDigits(n int) int {
	m := int(0)
	for t := n; math.Abs(float64(t)) > 0; t /= 10 {
		m = m*10 + int(math.Abs(float64(t)))%10
	}
	if m > math.MaxInt32 || m < math.MinInt32 {
		return 0
	}
	if n > 0 {
		return m
	}
	return -m

}

//Atoi 字符串转整型
func Atoi(str string) int {
	flag := 1
	res := 0
	slen := len(str)
	for i := 0; i < slen; i++ {
		if str[i] == '-' || str[i] == '+' || (str[i] >= '0' && str[i] <= '9') {
			if str[i] == '-' || str[i] == '+' {
				if str[i] == '-' {
					flag = -1
				}
				i++
			}
			for i < slen {
				if str[i] > '9' || str[i] < '0' {
					i = len(str)
					break
				}
				if flag*(res*10+int(str[i]-'0')) > 1<<31-1 {
					return 1<<31 - 1
				}
				if flag*(res*10+int(str[i]-'0')) < -1<<31 {
					return -1 << 31
				}
				res = res*10 + int(str[i]-'0')
				fmt.Println(res)
				i++
			}
		} else {
			if str[i] != ' ' {
				return 0
			}
		}
	}
	if flag*res > 1<<31-1 {
		return 1<<31 - 1
	}
	if flag*res < -1<<31 {
		return -1 << 31
	}
	return flag * res
}

/*
//IsNumber 判断字符串是否为有效十进制数字
func IsNumber(s string) bool {
	res := false
	flag := false
	slen := len(s)
	for i := 0; i < slen; i++ {
		if s[i] == '-' || s[i] == '+' || (s[i] >= '0' && s[i] <= '9') {
			if s[i] == '-' || s[i] == '+' {
				flag = true
				i++
			}
			for i < slen {
				if s[i] == '-' || s[i] == '+' {
					return false
				}

				if s[i] > '9' || s[i] < '0' {
					i = len(s)
					break
				}
				if flag*(res*10+int(s[i]-'0')) > 1<<31-1 {
					return 1<<31 - 1
				}
				if flag*(res*10+int(s[i]-'0')) < -1<<31 {
					return -1 << 31
				}
				res = res*10 + int(s[i]-'0')
				fmt.Println(res)
				i++
			}
		} else {
			if s[i] != ' ' {
				return 0
			}
		}
	}
	if flag*res > 1<<31-1 {
		return 1<<31 - 1
	}
	if flag*res < -1<<31 {
		return -1 << 31
	}
	return flag * res
}
*/

//IsPalindrome 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	a := []int{}
	for x > 0 {
		a = append(a, x%10)
		x = x / 10
	}
	for i, j := 0, len(a)-1; i < j; {
		if a[i] != a[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//IsPalindromeByStr 判断一个整数是否是回文数
func IsPalindromeByStr(x int) bool {
	a := strconv.Itoa(x)
	fmt.Println(a)
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[len(a)-i-1] {
			return false
		}
	}
	return true
}

//IsMatchByRecall 回溯法判断是否匹配
/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。\
*/
func IsMatchByRecall(s string, p string) bool {
	if len(p) <= 0 {
		return len(s) <= 0
	}
	firstMatch := (len(s) > 0) && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return IsMatchByRecall(s, p[2:]) || (firstMatch && IsMatchByRecall(s[1:], p))
	}
	return firstMatch && IsMatchByRecall(s[1:], p[1:])
}

//Map2DInt 二维整数集合
type Map2DInt map[string]bool

//Insert 二维整数集合插入操作
func (m Map2DInt) Insert(a, b int, r bool) {
	m[strconv.Itoa(a)+strconv.Itoa(b)] = r
}

//IsExist 判断是否存在
func (m Map2DInt) IsExist(a, b int) bool {
	_, ok := m[strconv.Itoa(a)+strconv.Itoa(b)]
	return ok
}

//Get 返回元素
func (m Map2DInt) Get(a, b int) (bool, bool) {
	val, ok := m[strconv.Itoa(a)+strconv.Itoa(b)]
	return val, ok
}

//IsMatchByDP 根据动态规划判断是否匹配
func IsMatchByDP(s string, p string) bool {
	m := Map2DInt{}
	return dp(0, 0, s, p, m)
}

func dp(i, j int, s, p string, m Map2DInt) bool {
	if !m.IsExist(i, j) {
		ans := false
		if j >= len(p) {
			ans = i >= len(s)
		} else {
			firstMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				ans = dp(i, j+2, s, p, m) || firstMatch && dp(i+1, j, s, p, m)
			} else {
				ans = firstMatch && dp(i+1, j+1, s, p, m)
			}
		}
		m.Insert(i, j, ans)
	}
	val, ok := m.Get(i, j)
	if ok {
		return val
	}
	return false
}

