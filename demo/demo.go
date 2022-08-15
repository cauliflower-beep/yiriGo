package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

// 测试用
var wg sync.WaitGroup

type people struct {
	age    int
	height int
}

type animal struct {
	gentle int
	color  string
}

func main() {
	//s := "abcdef"
	//fmt.Println([]byte(s))

	//fmt.Println(strStr("hello", "ll"))
	//nums := []int{1, 2, 3, 4, 5}
	//fmt.Println(nums[0:0])

	// 数字转字符串
	//fmt.Println([]byte(strconv.Itoa(1234)))

	// 字符串转数字
	//fmt.Println(strconv.Atoi("111"))

	//数组数字相转换
	//plusOne([]int{9, 9, 9})

	// 最长回文子串
	//fmt.Println(longestPalindrome("babad"))

	// 字符串拼接
	//fmt.Println("111 " + "222")

	// 二进制加法
	//fmt.Println(addBinary("11", "1"))

	//groutine同步问题
	//wg.Add(3)
	//for i := 0; i < 3; i++ {
	//	go goA()
	//}
	//fmt.Println("123")
	//wg.Wait()
	//fmt.Println("123456")

	////goroutine返回值问题
	//wg.Add(2)
	//pc := make(chan []people)
	//ac := make(chan []animal)
	//go func() {
	//	arr := []people{
	//		people{1, 12},
	//		people{2, 12},
	//		people{5, 12},
	//		people{4, 12},
	//		people{7, 12},
	//	}
	//	pc <- arr
	//
	//	arr2 := []animal{
	//		animal{1, "red"},
	//		animal{2, "green"},
	//		animal{2, "black"},
	//		animal{1, "white"},
	//		animal{2, "blue"},
	//	}
	//	ac <- arr2
	//	//for i := 0; i < 5; i++ {
	//	//	pc <- arr
	//	//}
	//	wg.Done()
	//}()
	//go func() {
	//	arr := []people{
	//		people{10, 12},
	//		people{11, 12},
	//		people{17, 12},
	//		people{14, 12},
	//		people{13, 12},
	//	}
	//	pc <- arr
	//	arr2 := []animal{
	//		animal{3, "red"},
	//		animal{4, "green"},
	//		animal{3, "black"},
	//		animal{4, "white"},
	//		animal{3, "blue"},
	//	}
	//	ac <- arr2
	//	wg.Done()
	//}()
	//
	//pcT := []people{}
	////acT := []animal{}
	////for i := 0; i < 4; i++ {
	////	select {
	////	case p := <-pc:
	////		fmt.Println("received", p)
	////		pcT = append(pcT, p...)
	////	case a := <-ac:
	////		fmt.Println("received", a)
	////		acT = append(acT, a...)
	////	}
	////}
	////
	////close(pc)
	////close(ac)
	//
	//wg.Wait()
	//close(pc)
	//close(ac)
	////for-range遍历
	//
	//for p := range pc {
	//	fmt.Println(p)
	//	pcT = append(pcT, p...)
	//}
	//
	//for a := range ac {
	//	fmt.Println(a)
	//}
	//
	//fmt.Println(pcT)
	//sort.Sort(newsDetail(pcT))
	//fmt.Println(pcT)

	//s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(s[1:2])
	//fmt.Println(s[:10])

	//快慢指针去重
	//fmt.Println(RemoveDuplicates([]string{"a", "b", "b", "b", "b", "c", "C", "c", "c"}))

	// bbs-go解疑
	p1 := p
	p2 := p
	fmt.Println(p1, p2, p1 == p2)
}

type newsDetail []people

func (n newsDetail) Len() int { return len(n) }

func (n newsDetail) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

func (n newsDetail) Less(i, j int) bool { return n[i].age < n[j].age }

//
func goA() {
	fmt.Println(111)
	wg.Done()
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

/*******************************二进制加法****************************/
func addBinary(a string, b string) string {
	// 转换成10进制
	numA, numB := bin2dec(a), bin2dec(b)

	log.Println(numA, numB)

	// 计算10进制和
	sum := numA + numB
	log.Println(sum)

	// 转换成二进制(顺序取余)
	return dec2bin(sum)
}

// 2进制转10进制
func bin2dec(s string) (num int) {
	num64, _ := strconv.ParseInt(s, 2, 64)
	return int(num64)
}

// 10进制转换为2进制
func dec2bin(n int) (bin string) {
	//for ; n > 0; n /= 2 {
	//	lsb := n % 2
	//	bin = strconv.Itoa(lsb) + bin
	//}
	bin = strconv.FormatInt(int64(n), 2)
	return
}

/********************************end*************************************/

/*******************************数组去重**********************************/
func RemoveDuplicates(str []string) (strWithoutDup []string) {
	// 快慢指针法，有bug，重复元素必须相邻才有效果
	//if len(str) == 0 {
	//	return str
	//}
	//left, right := 0, 1
	//for ; right < len(str); right++ {
	//	if str[left] == str[right] {
	//		continue
	//	}
	//	left++
	//	str[left] = str[right]
	//}
	//return str[:left+1]

	// 利用map去重，空间换时间，避免双重循环
	strMap := make(map[string]string)
	for _, v := range str {
		strMap[v] = v
	}
	for _, v := range strMap {
		strWithoutDup = append(strWithoutDup, v)
	}
	return
}

/***************************************************************************/
