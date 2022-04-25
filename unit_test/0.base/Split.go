package base

import "strings"

// Split 按照给定的分隔符 sep 对string进行切分，返回字符串切片
func Split(s,sep string )(result []string){
	i := strings.Index(s,sep)
	for i > -1{
		result = append(result,s[:i])
		//s = s[i+1:]
		s = s[i+len(sep):]	// fix bug
		i = strings.Index(s,sep)
	}
	result = append(result,s)
	return result
}
