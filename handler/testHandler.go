package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	Server = "http://1.2.3.4:3/myprojet/caje"
)

func GetHandler(iname string, params map[string]string) {
	url := Server + iname + "?"
	for k, v := range params {
		url += k + "=" + string(v) + "&"
	}
	url = strings.TrimSuffix(url, "&")
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	defer response.Body.Close()
}

func PostHandler(iname string, params map[string]string) {
	urlStr := Server + iname
	json := "{"
	for k, v := range params {
		json += "\"" + k + "\":"
		_, err := strconv.Atoi(v)
		if err != nil {
			json += "\"" + string(v) + "\","
		} else {
			json += string(v) + ","
		}
	}
	json = strings.TrimSuffix(json, ",")
	json += "}"
	data := make(url.Values)
	data["params"] = []string{json}
	fmt.Println(urlStr)
	fmt.Println(json)
	response, err := http.PostForm(urlStr, data)
	if err != nil {
		panic(err)
	}
	result, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(result))
	defer response.Body.Close()
}
