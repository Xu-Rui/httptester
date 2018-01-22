package main

import (
	"creator"
	"handler"
)

func main() {

	for {
		funcName, kvMap := creator.ResolveGetUrl("http://1.2.3.4:3/myprojet/caje/Comment?order_id=123&mask=4&type=0&id=456")
		handler.RandomMask(kvMap)
		handler.GetHandler(funcName, kvMap)

		jsonStr := "{firstName = \"John\"}"
		funcName, kvMap = creator.ResolvePostJson("http://1.2.3.4:3/myprojet/caje/Comment", jsonStr)
		handler.RandomMask(kvMap)
		handler.PostHandler(funcName, kvMap)
	}
}

