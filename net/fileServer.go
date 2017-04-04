package main

import "net/http"

func main() {
	h := http.FileServer(http.Dir("D:/BaiduNetdiskDownload/新建文件夹"))
	http.ListenAndServe(":8888", h)
}
