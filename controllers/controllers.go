package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/danielnill/go_hack/models"
	"github.com/danielnill/go_hack/web"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

func FrontPageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := web.Get("https://hacker-news.firebaseio.com/v0/topstories.json")

	if err != nil {
		fmt.Println(err)
	}

	var ids []int
	if err = json.Unmarshal(body, &ids); err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup
	var items [30]models.Item
	for i, id := range ids[0:30] {
		wg.Add(1)
		go models.GetAndAddItem(&items, id, i, &wg)
	}
	wg.Wait()

	body, err = json.Marshal(items)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func DiscussionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var wp sync.WaitGroup
	item := models.GetItem(vars["id"])
	body := item.BuildTree(&wp).ToJSON()
	wp.Wait()
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
