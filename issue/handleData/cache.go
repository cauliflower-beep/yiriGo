package main

import (
	"encoding/json"
	"fmt"
)

/*
背景：
	PMsg系统接收自选股更新消息，拿到uid之后，需要去dcache中获取该uid的全部自选股列表
	解析缓存数据如下：
*/

/*
反引号可以定义包含 json 格式的大字符串
注意要保留前后的 "[" 及 "]" 才可以正常解码，
不要自作聪明的把这两个符号 trim 掉(从某不知名csdn评论中看到的，谢谢你，无名侠).
*/
var cacheData = `[
{"scode":"688016","platform":"android","groupid":0,"marketid":1,"priority":0,"position":"AAAAAP////sYMDF8MDB8Njg4MDE2fGVtcDEyMzQ1Njc4OXVpZA==","status":1,"deletetime":0,"updatetime":1660893317723,"createtime":1660619136382},
{"scode":"000001","platform":"android","groupid":0,"marketid":0,"priority":1,"position":"AAAAAP////swMHwwMHwwMDAwMDF8ZW1wMTIzNDU2Nzg5dWlk","status":1,"deletetime":0,"updatetime":1660893317723,"createtime":1660275048462},
{"scode":"510050","platform":"android","groupid":0,"marketid":1,"priority":2,"position":"AAAAAP////swMXwwMHw1MTAwNTB8MDAwMDEwMDEwMDAwMDE5","status":1,"deletetime":0,"updatetime":1660893317723,"createtime":1659662148785},
{"scode":"600275","platform":"android","groupid":0,"marketid":1,"priority":3,"position":"AAAAAP////swMXwwMHw2MDAyNzV8MDAwMDEwMDEwMDAwMDE5","status":1,"deletetime":0,"updatetime":1660893317723,"createtime":1660180008081},
{"scode":"00001","platform":"android","groupid":0,"marketid":64,"priority":4,"position":"AAAAATY0fDAwfDAwMDAwMXxlbXAxMjM0NTY3ODl1aWQ=","status":1,"deletetime":0,"updatetime":1660893317723,"createtime":1660097160939},
{"scode":"000666","platform":"android","groupid":0,"marketid":0,"priority":5,"position":"AAAAAzAwfDAwfDAwMDY2NnwwMDAwMTAwMTAwMDAwMTk=","status":1,"deletetime":0,"updatetime":1660893317723,"createtime":1659938685072}
]`

type T struct {
	Scode string `json:"scode"`
	//Platform   string `json:"platform"`
	//Groupid    int    `json:"groupid"`
	//Marketid   int    `json:"marketid"`
	//Priority   int    `json:"priority"`
	//Position   string `json:"position"`
	//Status     int    `json:"status"`
	//Deletetime int64  `json:"deletetime"`
	//Updatetime int64  `json:"updatetime"`
	//Createtime int64  `json:"createtime"`
}

func main() {
	t_list := []T{}
	//cache_trim := strings.Trim(cacheData, "[")
	//cache_trim2 := strings.Trim(cache_trim, "]")
	if err := json.Unmarshal([]byte(cacheData), &t_list); err != nil {
		fmt.Println(err)
	} else {
		for _, rec := range t_list {
			fmt.Println(rec.Scode)
		}
	}
}