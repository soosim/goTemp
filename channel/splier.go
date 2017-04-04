package main

import (
	"fmt"
	"net/http"
	"time"
)

var urlChannel = make(chan string, 200)

var stagRegExp = regexp.MustCompile(`<a[^>]+[(href)|(HREF)]\s*\t*\n*=\s*\t*\n*[(".+")|('.+')][^>]*>[^<]*</a>`)

func main() {
	go Spy("http://blog.csdn.net")
	for url := range urlChannel {
		fmt.Println("routines num=", runtime.NumGoroutine(), "chanlen=", len(urlChannel))
		go Spy(url)
	}

	fmt.Println("a")
}

func Spy(url string) {
	defer func() {
		if r := recover(); r != nil {

		}
	}()

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-gent", GetRandomAgent())
	cilent := http.DefaultClient
	res, e := cilent.Do(req)
	if e != nil {
		fmt.Errorf("Get 请求%s返回错误：%s", url, e)
		return
	}

	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()

		bodyByte, _ := ioutil.ReadAll(body)
		resStr := string(bodyByte)

		atag := atagRegExp.FindAllString(resStr, -1)
		for _, a := range atag {
			href, _ := GetHref(a)
			if strings.Cotains(href, "article/details/") {
				fmt.Println("***", href)
			} else {
				fmt.Println("xxx", href)
			}
		}
	}
}
func GetRandomAgent() string {
	var userAgent = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
		"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
		"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
		"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
		"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
		"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgent[r.Intn(len(userAgent))]
}
