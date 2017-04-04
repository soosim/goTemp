package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Entry struct {
	Title     string
	Author    string
	URL       string
	Permalink string
}

type Feed struct {
	Data struct {
		Children []struct {
			Data Entry
		}
	}
}

func main() {
	url := "http://www.reddit.com/r/golang/new.json"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error fetching:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Error Status not OK:", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading body", err)
	}
	var entrier Feed

	if err := json.Unmarshal(body, &entrier); err != nil {
		log.Fatalln("Error decoding JSON", err)
	}

	for _, ed := range entrier.Data.Children {
		entry := ed.Data
		log.Println(">>>")
		log.Println("Title :", entry.Title)
		log.Println("Author :", entry.Author)
		log.Println("URL :", entry.URL)
		log.Println("Permalink :", entry.Permalink)
	}

}
