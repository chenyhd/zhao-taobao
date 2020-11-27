package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {

	domain := flag.String("domain", "http://127.0.0.1:8080", "发送请求的域名")
	headerFile := flag.String("header", "./header.txt", "请求头文件路径")
	dataFile := flag.String("data", "./data.txt", "数据文件路径")
	outputFile := flag.String("output", "./output.txt", "输出文件路径")
	outputFileTimestamp := flag.String("timestamp", "true", "输出文件路径是否追加时间戳")
	flag.Parse()
	data, err := ioutil.ReadFile(*dataFile)
	check(err)

	split := strings.Split(string(data), "\n")

	var headerText, _ = ioutil.ReadFile(*headerFile)

	header := http.Header{}
	for _, b := range strings.Split(string(headerText), "\n") {
		i := strings.Split(b, ": ")
		header.Add(i[0], i[1])
	}

	log.Println(header)

	var output []byte

	for i, s := range split {
		log.Print("正在处理第", i, "个, 值为[", s, "]")
		log.Print(string(headerText))

		urlStr := *domain + "/get?keywords=" + s + "&sleepSecond=2"

		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, urlStr, nil)

		if err != nil {
			fmt.Println(err)
		}

		// Delete encoding header
		header.Del("Accept-Encoding")

		req.Header = header

		res, err := client.Do(req)

		if err != nil {
			log.Println(err)
			return
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)

		log.Println(body)

		output = append(output[:], body...)

	}

	b, _ := strconv.ParseBool(*outputFileTimestamp)

	if b {
		*outputFile += strconv.FormatInt(time.Now().Unix(), 10)
	}

	_ = ioutil.WriteFile(*outputFile, output, 0644)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
