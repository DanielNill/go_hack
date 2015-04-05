package models

import (
	"encoding/json"
	"fmt"
	"github.com/danielnill/go_hack/web"
	"strconv"
	"sync"
)

type Item struct {
	By       string `json:"by"`
	Id       int    `json:"id"`
	Kids     []int  `json:"kids"`
	Parent   int    `json:"parent"`
	Title    string `json:"title"`
	Text     string `json"text"`
	Time     int    `json:"time"`
	ItemType string `json:"type"`
	Children []*Item
}

func GetItem(id string) Item {
	item := Item{}
	body, err := web.Get("https://hacker-news.firebaseio.com/v0/item/" + id + ".json")
	if err != nil {
		fmt.Println(err)
	}

	if err = json.Unmarshal(body, &item); err != nil {
		fmt.Println(err)
	}
	return item
}

func (self *Item) ToJSON() []byte {
	body, err := json.Marshal(self)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func GetAndAddItem(items *[30]Item, id int, index int, wg *sync.WaitGroup) {
	item := GetItem(strconv.Itoa(id))
	items[index] = item
	wg.Done()
}

func (self *Item) BuildTree(wp *sync.WaitGroup) *Item {
	wp.Add(1)
	for _, id := range self.Kids {
		item := GetItem(strconv.Itoa(id))
		self.Children = append(self.Children, &item)
		go item.BuildTree(wp)
	}
	wp.Done()
	return self
}
