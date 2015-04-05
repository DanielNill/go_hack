package web

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Close = true
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return body, err
}

func StartServer(r *mux.Router) {
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
