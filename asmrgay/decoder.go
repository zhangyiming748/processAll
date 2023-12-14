package asmrgay

import (
	"fmt"
	"net/url"
)

func decoder(OriginalUrl string) string {
	//fmt.Println(OriginalUrl)
	encodeurl := url.QueryEscape(OriginalUrl)
	fmt.Println(encodeurl)
	//decodeurl, err := url.QueryUnescape(OriginalUrl)
	//if err != nil {
	//	fmt.Println(err)
	//}
	return encodeurl
}
func encoder(OriginalUrl string) string {
	decodeurl, err := url.QueryUnescape(OriginalUrl)
	if err != nil {
		fmt.Println(err)
	}
	return decodeurl
}
