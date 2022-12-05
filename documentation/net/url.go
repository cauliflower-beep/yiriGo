package main

import (
	"fmt"
	"net/url"
)

/*
	工作中遇到一个问题，需要在url中拼接带有空格的参数。
	参考了这篇文章解决：
	https://lequ7.com/guan-yu-golangurl-zhong-de-kong-ge-jia-hao-jiu-jing-ying-gai-shi-yong-he-zhong-fang-shi-bian-ma.html
	总结:
	url中不能显示的蕴含空格，这是共识。
	但空格以何种模式存在，在不同的规范中不完全一致，以致于不同的语言有不同的实现。
	go中推荐使用 url.PathEscape 来编码
*/

func main() {
	urlRaw := "http://172.29.2.104:29999/busapi/emall/equity/queryUserEquity?app_key=cairenhui1004&sign=8f62c226239bb28c96834bb0c81b6559&equity_code=FXYJ&start_datetime=2022-12-5 15:18:12&end_datetime=2022-12-5 15:18:22&status=1"
	urlEscape := url.PathEscape(urlRaw)
	fmt.Println(urlEscape)
	urlQuery := url.QueryEscape(urlRaw)
	fmt.Println(urlQuery)
}
