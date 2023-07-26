package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func printABC(wg *sync.WaitGroup, ch chan int, str string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		<-ch
		fmt.Println(str)
		ch <- 1
	}
}

type OpenApiMsg struct {
	Url string `json:"url"`
}

//func (o *OpenApiMsg) MarshalJSON() ([]byte, error) {
//	urlValue := strings.ReplaceAll(o.Url, "&", "\\u0026")
//	data := struct {
//		Url string `json:"url"`
//	}{
//		Url: urlValue,
//	}
//	return json.Marshal(data)
//}

func main() {
	op := OpenApiMsg{
		Url: "&",
	}
	openApiMsgReqJson, _ := json.Marshal(&op)
	fmt.Println(string(openApiMsgReqJson))

}
