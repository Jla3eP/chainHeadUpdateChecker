package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	ping := true
	ticker := time.NewTicker(1 * time.Second)

	request, _ := http.NewRequest(http.MethodGet, "http://localhost:3500/eth/v1alpha1/beacon/chainhead", nil)
	client := http.DefaultClient

	pinger := sync.Once{}

	var lastRespBody string
	for {
		<-ticker.C
		resp, _ := client.Do(request)
		respBodyStr, _ := ioutil.ReadAll(resp.Body)

		if string(respBodyStr) != lastRespBody {
			if ping && lastRespBody != "" {
				pinger.Do(func() {
					go func() {
						for i := 0; i < 15; i++ {
							time.Sleep(80 * time.Millisecond)
							fmt.Print("\a")
						}
					}()
				})
			}
			lastRespBody = string(respBodyStr)
			fmt.Println(lastRespBody + "\n")
		}
	}
}
