package main

import (
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type Dog struct {
	Name  string
	Color string
}

func main() {
	mc := memcache.New("127.0.0.1:11211")

	fetchItem, err := mc.Get("dog")

	if err != memcache.ErrCacheMiss {
		if err != nil {
			fmt.Println("Error fetching from memcache", err)
		} else {
			fmt.Println("Cache hit!")
			dog, err := DecodeData(fetchItem.Value)
			if err != nil {
				fmt.Println("Error")
			} else {
				fmt.Println("Dog name is: ", dog.Name)
			}
		}
	}

	spot := Dog{Name: "Spot", Color: "brown"}

	setItem := memcache.Item{key: "dog", Value: EncodeData(spot), Expiration: 300}
	err = mc.Set(&setItem)

	if err != nil {
		fmt.Println("Error")
	}
}

func DecodeData(raw []byte) (dog Dog, err error) {
	err = json.Unmarshal(raw, &dog)
}

func EncodeData(dog Dog) []byte {
	enc, err := json.Marshal(dog)
	if err != nil {
		fmt.Println("Error encoding")
	}
	return enc
}
