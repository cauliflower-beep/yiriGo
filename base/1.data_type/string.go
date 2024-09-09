package main

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
)

// string 转 []map
func str2map(s string) (res []map[string]interface{}) {
	_ = json.Unmarshal([]byte(s), &res)
	return
}

// map 转 string
func map2string(m []map[string]interface{}) (s string) {
	bs, _ := json.Marshal(m)
	s = string(bs)
	return
}

// slice 转 string
func slice2string(descs []string) (s string) {
	s = ""
	for i, desc := range descs {
		if i != len(descs)-1 {
			s += desc + "、"
			continue
		}
		s += desc
	}
	return
}
func main() {
	// string 转 []map
	//m := str2map("[{\"000001\":{\"stockName\":\"xxx\",\"lastRiskLevel\":1,\"desc\":\"xxxx\",\"update\":1234}}," +
	//	"{\"000002\":{\"stockName\":\"xxx\",\"lastRiskLevel\":1,\"desc\":\"xxxx\",\"update\":1234}}]")
	//m := str2map("[{\"000001\":{\"stockName\":\"草帽海贼团\",\"lastRiskLevel\":1,\"thisRiskLevel\":1,\"desc\":\"资金问题\",\"update\":1666828808}}]")
	//fmt.Println(len(m), m)

	// []map 转 string
	//s := map2string(m)
	//fmt.Println(s)

	// 切片转字符串
	descs := []string{"路飞", "索隆", "山治"}
	fmt.Println(slice2string(descs))

	fmt.Println(utf8.RuneCountInString("你好"))
}
