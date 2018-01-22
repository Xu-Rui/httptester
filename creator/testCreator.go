package creator

import (
	"strings"
)

func ResolveGetUrl(url string)  (string,map[string]string) {
	kvmap := make(map[string]string)
	temp := strings.Split(url,"?")
	funName := strings.Split(temp[0],"/")[5]
	params := strings.Split(temp[1],"&")
	for _,str := range params {
		temp = strings.Split(str,"=")
		kvmap[temp[0]] = temp[1]
	}
	return funName,kvmap
}

func ResolvePostJson(url string,jsonStr string) (string,map[string]string) {
	kvmap := make(map[string]string)
	funName := strings.Split(url,"/")[5]
	temp := strings.Trim(strings.Trim(jsonStr,"{"),"}")
	strArray := strings.Split(temp,",")
	for _,v := range strArray{
		ptemp :=strings.SplitN(v,":",2)
		kvmap[strings.Trim(ptemp[0],"\"")] = strings.Trim(ptemp[1],"\"")
	}
	return funName,kvmap
}

