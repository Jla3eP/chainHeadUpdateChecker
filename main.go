package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ping := true
	ticker := time.NewTicker(1 * time.Second)

	request, _ := http.NewRequest(http.MethodGet, "http://localhost:3500/eth/v1alpha1/beacon/chainhead", nil)
	client := http.DefaultClient

	var lastRespBody string
	for {
		<-ticker.C
		resp, _ := client.Do(request)
		respBodyStr, _ := ioutil.ReadAll(resp.Body)

		if string(respBodyStr) != lastRespBody {
			lastRespBody = string(respBodyStr)
			fmt.Println(lastRespBody + "\n")
			if ping {
				fmt.Print("\a")
			}
		}
	}
}
