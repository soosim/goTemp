package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/postpage", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var (
				key   string = r.PostFormValue("key")
				value string = r.PostFormValue("value")
			)

			fmt.Println("key is : %s\n", key)
			fmt.Println("value is : %s\n", value)
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
