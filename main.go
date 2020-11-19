package main

import (
	"awesomeProject/mode"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	var keywords = []string{
		"麦饭石乌龟锅",
		"做豆腐的卤水 家用",
		"烤火炉子烧柴 室内",
	}


	for _, keyword := range keywords {
		analysis(keyword)
	}

}

func analysis(keyWord string) {

	escapeKeyWord := url.QueryEscape(keyWord)

	urlStr := "https://s.taobao.com/search?q=" + escapeKeyWord + "&imgfile=&commend=all&ssid=s5-e&search_type=item&sourceId=tb.index&spm=a21bo.2017.201856-taobao-item.1&ie=utf8&initiative_id=tbindexz_20170306&sort=sale-desc"

	method := "GET"

	client := &http.Client{
	}
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
	req.Header.Add("cookie", " _samesite_flag_=true; cookie2=1b6942071f1cdb2b594cc518ee20b955; t=14de9fffe0bdf37ddb935132b30ba6f5; _tb_token_=ee0f7bbab9378; xlly_s=1; sgcookie=E1003aGpkVfI4EKty%2B2TqS%2BPTddeJhJ%2F0zfDEdGdelDZfHHI18dkC0XDz%2Bckkp47E7btobzPu1hkrKqbUAO92mvMWA%3D%3D; unb=2363180808; uc3=vt3=F8dCufwsB9cyvzFfDCg%3D&nk2=EF8d6D87H2r2Op%2F9GkXxuGU%3D&lg2=VT5L2FSpMGV7TQ%3D%3D&id2=UUtNgwq20OCazQ%3D%3D; csg=8bf9311d; lgc=salah%5Cu795E%5Cu5947%5Cu7684%5Cu5C0F%5Cu5C3E%5Cu5DF4; cookie17=UUtNgwq20OCazQ%3D%3D; dnk=salah%5Cu795E%5Cu5947%5Cu7684%5Cu5C0F%5Cu5C3E%5Cu5DF4; skt=78d44291fa2c9fdc; existShop=MTYwNTcwOTMyMg%3D%3D; uc4=nk4=0%40EoZib8p25U4avRrYYe4ubfW%2FiFRbZPs7DmxR1g%3D%3D&id4=0%40U2l3yMzMraLu1KHygHCFc2Edpp1A; tracknick=salah%5Cu795E%5Cu5947%5Cu7684%5Cu5C0F%5Cu5C3E%5Cu5DF4; _cc_=VFC%2FuZ9ajQ%3D%3D; _l_g_=Ug%3D%3D; sg=%E5%B7%B483; _nk_=salah%5Cu795E%5Cu5947%5Cu7684%5Cu5C0F%5Cu5C3E%5Cu5DF4; cookie1=W5iTmWxTLlBy3DXqu7ZgwFvN7qzUF0D%2FkJF11KCdYi8%3D; enc=HF%2BYFw%2BZATJ87UZkGEfDwzDdlLhqNSt%2BvhH%2B%2F4zE%2F8YK4hjTQCKFvA3H1A65T7bzVET9QsyKlPVer2A5P3gsgw%3D%3D; hng=CN%7Czh-CN%7CCNY%7C156; thw=cn; mt=ci=11_1; uc1=cookie21=WqG3DMC9Fb5mPLIQoVXj&pas=0&cookie16=W5iHLLyFPlMGbLDwA%2BdvAGZqLg%3D%3D&cookie14=Uoe0aDvcrKNsyA%3D%3D&existShop=false&cookie15=VT5L2FSpMGV7TQ%3D%3D; JSESSIONID=016690D74674551D0D9019D88DCC96E3; cna=xabHF8GmW3ACATonI7Wpx0BH; isg=BEBAP_82C7ZDl_bBGYWCbWY0EccSySSTaRul0rrRDNvuNeBfYtn0IxYHSZ31ntxr; tfstk=cMnlB0iDeNXIHDzm53ZSBa1bLA-OwySUsDoj0mSdd0KCqg5cPZ7wFlJ0ZbmdR; l=eBPjtwMeQflVAW3BBOfanurza77OSIRYYuPzaNbMiOCPs2fB51dOWZ75M_Y6C3GVh6bWR3J42eyyBeYBYQAonxv92j-la_kmn; JSESSIONID=7B02F83B372BC589865227AE2F5DF2D9")
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
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var result string

	for _, lineStr := range strings.Split(string(body), "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}

		contain := strings.Contains(lineStr, "g_page_config")

		if contain {
			result = lineStr
			break
		}
	}

	//log.Println(result)

	start := strings.Index(result, "=")

	jsonStr := substr(result, start+2)

	//log.Println(jsonStr)

	var item mode.Item
	err = json.Unmarshal([]byte(jsonStr), &item)
	if err != nil {
		log.Println(err)
		return
	}

	auctions := item.Mods.Itemlist.Data.Auctions

	s := auctions[len(auctions)-1]
	log.Println(keyWord, "-->", s.ViewSales)
}

func substr(input string, start int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	length := len(asRunes)

	return string(asRunes[start : length-1])
}
