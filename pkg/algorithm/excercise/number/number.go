package number

import (
	"fmt"
	"math"
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
