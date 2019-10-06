package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Запускаемся. Слушаем порт 80")
	_ = http.ListenAndServe(":80", nil)
}

func handler(iWrt http.ResponseWriter, iReq *http.Request) {
	var lGet = iReq.URL.Path[1:]
	log.Println(iReq)
	if lGet == "" || lGet == "/" {
		lGet = "index.html"
	}
	lData := " products: [{id: 5, url: './img/IMG_2.jpg', cost: '100', category: 'Бр11оши'},{id: 6, url: './img/IMG_2.jpg', cost: '100', category: 'Бр11оши'},{id: 7, url: './img/IMG_2.jpg', cost: '100', category: 'Бр11оши'}]"
	iWrt.Header().Set("Access-Control-Allow-Origin", "*")
	_, _ = fmt.Fprintln(iWrt, lData)
}

func readFile(iFileName string) string {
	lData, err := ioutil.ReadFile(iFileName)
	var lOut string
	if !os.IsNotExist(err) {
		lOut = string(lData)
	} else {
		lOut = "404"
	}
	return lOut
}
