package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符拼接
	s := []string{"111", "222", "333", "444", "555"}
	str := strings.Join(s, ":")
	str1 := strings.Join(s, "\n")
	fmt.Println(str, str1)
	fmt.Println()

	// 字符拆分
	str2 := "hello|my|boy|lufy"
	//strSplit := strings.Split(str2, "|")
	strSplit := strings.SplitN(str2, "|", 2) // 指定分割符号，指定分割次数
	fmt.Println(strSplit, len(strSplit))

	// 字符替换
	str3 := "BVc5sIR3GJkHjE6uFrnqOA@"
	strRep := strings.Replace(str3, "&", "=", -1)
	strRep = strings.Replace(strRep, "@", "==", -1)
	fmt.Println(strRep)
}
