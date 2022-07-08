package main

import "fmt"

// 测试用

func main() {
	//s := "abcdef"
	//fmt.Println([]byte(s))

	//fmt.Println(strStr("hello", "ll"))
	//nums := []int{1, 2, 3, 4, 5}
	//fmt.Println(nums[0:0])

	// 数字转字符串
	//fmt.Println([]byte(strconv.Itoa(1234)))

	//数组数字相转换
	//plusOne([]int{9, 9, 9})

	// 最长回文子串
	//fmt.Println(longestPalindrome("babad"))

	// 字符串拼接
	fmt.Println("111 " + "222")

}

// str匹配
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	hi := 0
	for len(haystack)-hi >= len(needle) {
		if haystack[hi:len(needle)+hi] != needle {
			hi++
			continue
		} else {
			return hi
		}
	}
	return -1
}

// 数组组成数字，数字加1再返回数组，如输入 [1,2,3],输出[1,2,4]
//func plusOne(digits []int) []int {
//	//计算输入数组表示的数字
//	num := 0
//	for i := 0; i < len(digits); i++ {
//		num += digits[i] * int(math.Pow(10, float64(len(digits)-i-1)))
//	}
//	num++
//
//	// 将加1后的数字还原成数组
//	length := len(strconv.Itoa(num))
//	numList := []int{}
//
//	for i := 0; i < length; i++ {
//		numList = append(numList, num/int(math.Pow(10, float64(length-i-1))))
//		num -= numList[i] * int(math.Pow(10, float64(length-i-1)))
//	}
//
//	return numList
//
//}

func plusOne(digits []int) []int {
	// 思路一：不同类型转来转去
	// //计算输入数组表示的数字
	// var num float64
	// for i:=0;i<len(digits);i++{
	//     num += float64(digits[i])*math.Pow(10,float64(len(digits)-i-1))
	// }
	// num ++

	// // 将加1后的数字还原成数组
	// length := len(strconv.Itoa(int(num)))
	// numList := []int{}

	// for i:=0;i<length;i++{
	//     numList = append(numList,int(num/math.Pow(10,float64(length-i-1))))
	//     num -= float64(numList[i]) * math.Pow(10,float64(length-i-1))
	// }
	// return numList

	//思路二：直接用数组来计算
	if len(digits) == 1 {
		if digits[0] == 9 {
			return []int{1, 0}
		} else {
			digits[0] = digits[0] + 1
			return digits
		}

	}
	digits[len(digits)-1] += 1
	for i := len(digits) - 1; i > 0; i-- {

		if digits[i] >= 10 {
			digits[i] = digits[i] % 10
			digits[i-1] += 1
		} else {
			break
		}
	}
	if digits[0] >= 10 {
		digits[0] = digits[0] % 10
		res := make([]int, 1)
		res[0] = 1
		res = append(res, digits...)
		return res
	} else {
		return digits
	}

}

// 求最长回文子串
func longestPalindrome(s string) string {
	/*
	   从第一个字符开始遍历，找出与他回文的所有字串，
	   并比较长短，记录最长的回文子串
	   直到遍历完所有的字符。
	*/

	/*
	   既然是找最长的，感觉从长往短找会快一些。
	   采用双指针，从两边往中间缩，直到找到回文子串。
	*/
	if len(s) == 0 {
		return ""
	}
	if isPalindrome(s) {
		return s
	}
	s_str := []byte(s)
	left := 0
	right := len(s) - 1
	for left < right {
		if isPalindrome(string(s_str[left:right])) {
			fmt.Println(string(s_str[left:right]))
			return string(s_str[left:right])
		} else if isPalindrome(string(s_str[left+1 : right+1])) {
			fmt.Println(string(s_str[left+1 : right+1]))
			return string(s_str[left+1 : right+1])
		}
		left++
		right--
	}
	return ""

}

func isPalindrome(s string) bool {
	// 反转字符串
	str := []byte(s)
	l := 0
	r := len(s) - 1
	for l < r {
		str[l], str[r] = str[r], str[l]
		l++
		r--
	}
	s_rev := string(str)
	if s == s_rev {
		return true
	} else {
		return false
	}
}
