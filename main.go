package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

const localUrl = "http://ip2c.org/"

func main() {
	var ip string
	flag.StringVar(&ip, "ip", "", "IP位址")
	flag.Parse()
	isIP := net.ParseIP(ip)

	if len(ip) == 0 || isIP == nil {

		panic("請輸入正確IP格式")
	}

	response, err := http.Get(localUrl + ip)
	defer response.Body.Close()
	//網站發生錯誤
	if err != nil {
		fmt.Printf("Curl Error: %s\n", err)
	}
	if response.StatusCode != 200 {
		fmt.Printf("Http Code Error: %d\n", response.StatusCode)
	} else {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		localInfo := strings.Split(bodyString, ";")
		fmt.Println("IP:" + ip + "\nLocation:" + localInfo[1] + "," + localInfo[3])
	}
}
