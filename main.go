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

		splits := strings.Split(keywords, ",")

		log.Println(splits, sleepSecond)

		parseInt, _ := strconv.Atoi(sleepSecond)

		var result []string

		for _, keyword := range splits {
			s := analysis(keyword)
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

func analysis(keyWord string) string {

	escapeKeyWord := url.QueryEscape(keyWord)

	urlStr := "https://s.taobao.com/search?q=" + escapeKeyWord + "&imgfile=&commend=all&ssid=s5-e&search_type=item&sourceId=tb.index&spm=a21bo.2017.201856-taobao-item.1&ie=utf8&initiative_id=tbindexz_20170306&sort=sale-desc"

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, urlStr, nil)

	if err != nil {
		fmt.Println(err)
	}
	//req.Header.Add("", "authority: s.taobao.com")
	//req.Header.Add("", "method: GET")
	//req.Header.Add("", "path: /search?q=%E5%86%9C%E6%9D%91%E6%9F%B4%E7%81%AB%E7%81%B6%E5%AE%B6%E7%94%A8%E7%83%A7%E6%9C%A8%E6%9F%B4%E8%8A%82%E8%83%BD&imgfile=&commend=all&ssid=s5-e&search_type=item&sourceId=tb.index&spm=a21bo.2017.201856-taobao-item.1&ie=utf8&initiative_id=tbindexz_20170306&sort=sale-desc")
	//req.Header.Add("", "scheme: https")
	req.Header.Add("accept", " text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Add("accept-encoding", " gzip, deflate, br")
	req.Header.Add("accept-language", " zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,ja;q=0.6,zh-TW;q=0.5,und;q=0.4")
	req.Header.Add("cache-control", " max-age=0")
	req.Header.Add("cookie", "thw=cn; UM_distinctid=17534259f4b484-0adba98985f4cb-376b4502-1fa400-17534259f4c96; CNZZDATA1279311025=1149418505-1602894843-https%253A%252F%252Fwww.taobao.com%252F%7C1605265423; t=2c063a1025d7329ac64d4d93c9764e72; xlly_s=1; hng=CN%7Czh-CN%7CCNY%7C156; enc=B1I%2BwKyQvGYBgpjshpEjf%2BGWgJ9mkW9dbPxsDvONTI3pfX82ykZbgUrzrviGTDVyBnptQq4DxezmWjpvVlLRLA%3D%3D; cna=IbAQGHH9fTwCAXtltulpUtCi; sgcookie=E100y4td6oGSvxJfW9Us3YZrc1EJqlcmoiMYkjP%2B%2F7zT2to97UbWPy3Ng%2BROiE%2BIcry1PTNDJEwgjNgSXzIf1d%2B3QA%3D%3D; uc3=lg2=UIHiLt3xD8xYTw%3D%3D&nk2=F5REOtIWfqrLbtM%3D&vt3=F8dCufwqQPuCmvZQPXI%3D&id2=UUphydtzbn68usiH1A%3D%3D; lgc=tb142554118; uc4=id4=0%40U2grEhU66ueQ0vkgnHeYfTNeYYOYGxhS&nk4=0%40FY4PaJKhammAfqdf84pgOGuGLndsoA%3D%3D; tracknick=tb142554118; _cc_=URm48syIZQ%3D%3D; mt=ci=6_1; v=0; _tb_token_=738137117ee3; alitrackid=www.taobao.com; CNZZDATA1277903761=1914218829-1602895946-https%253A%252F%252Fwww.taobao.com%252F%7C1605955990; CNZZDATA1279154029=1345400558-1602894962-https%253A%252F%252Fwww.taobao.com%252F%7C1605960540; _samesite_flag_=true; cookie2=174efd634213a39bb5a47e320e3414cc; tfstk=cEmNB7wb_hKNWiEXylZV5Yth8GkOawDinMy7Sq8XlVJeY9q_asvhMJvCMJyfGPUG.; l=eBLDkugcO-vr7h6dBOfZourza77TjIRAguPzaNbMiOCPOMfB5fzdWZ718QY6CnGVh6xBR3rEQAfvBeYBqgI-nxvOvhLyCKkmn; _m_h5_tk=a23a1e88184e936880d1b4984934765b_1605968747200; _m_h5_tk_enc=c5815499793f45f383b9598f8fca805b; uc1=cookie14=Uoe0aDVuy5%2FNYw%3D%3D; JSESSIONID=67A383299B53B0AF38C89C74D522E41C; lastalitrackid=www.taobao.com; isg=BDw8QZS6T3YVhHswoYvWp2qFDdruNeBfca0MTBa9XycK4d1rPkc772NZxQmZqRi3")
	req.Header.Add("sec-fetch-dest", " document")
	req.Header.Add("sec-fetch-mode", " navigate")
	req.Header.Add("sec-fetch-site", " same-origin")
	req.Header.Add("sec-fetch-user", " ?1")
	req.Header.Add("upgrade-insecure-requests", " 1")
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("user-agent", " Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")

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

	//log.Println(jsonStr)

	var item mode.Item
	err = json.Unmarshal([]byte(jsonStr), &item)
	if err != nil {
		log.Println(err,jsonStr)
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
