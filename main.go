package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"taobao/mode"
	"time"
)

func main() {

	r := gin.Default()

	r.GET("/get", func(c *gin.Context) {

		keywords := c.Query("keywords")
		sleepSecond := c.Query("sleepSecond")
		header := c.Request.Header

		splits := strings.Split(keywords, ",")

		log.Println(splits, sleepSecond)

		parseInt, _ := strconv.Atoi(sleepSecond)

		var result []string

		for _, keyword := range splits {
			s := analysis(keyword, header)
			result = append(result, s)
			time.Sleep(time.Duration(parseInt) * time.Second)
		}

		//var keywords = []string{
		//	"麦饭石乌龟锅",
		//	"做豆腐的卤水 家用",
		//	"烤火炉子烧柴 室内",
		//}
		c.JSON(200, gin.H{
			"message": result,
		})

	})

	r.Run("0.0.0.0:8080")

}

func analysis(keyWord string, header http.Header) string {

	escapeKeyWord := url.QueryEscape(keyWord)

	urlStr := "https://s.taobao.com/search?q=" + escapeKeyWord + "&imgfile=&commend=all&ssid=s5-e&search_type=item&sourceId=tb.index&spm=a21bo.2017.201856-taobao-item.1&ie=utf8&initiative_id=tbindexz_20170306&sort=sale-desc"

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
		return err.Error()
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var targetJson string

	for _, lineStr := range strings.Split(string(body), "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}

		contain := strings.Contains(lineStr, "g_page_config")

		if contain {
			targetJson = lineStr
			break
		}
	}

	log.Println(targetJson)

	start := strings.Index(targetJson, "=")

	jsonStr := substr(targetJson, start+2)

	var item mode.Item
	err = json.Unmarshal([]byte(jsonStr), &item)
	if err != nil {
		log.Println(err, "[", jsonStr, "]")
		return err.Error() + jsonStr
	}

	auctions := item.Mods.Itemlist.Data.Auctions

	s := auctions[len(auctions)-1]

	result := keyWord + "--" + s.ViewSales

	log.Println(result)

	return result
}

func substr(input string, start int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	length := len(asRunes)

	return string(asRunes[start : length-1])
}
