package utils

import (
	"fmt"
	"net/url"
)

// https%3A%2F%2Fexample.com%2F%3Fq%3Dgolang%26page%3D1
func UrlDecode(encodedURL string) (string, error) {
	// 解码URL
	decodedURL, err := url.QueryUnescape(encodedURL)
	if err != nil {
		fmt.Println("URL解码失败:", err)
		return "", err
	}

	decodedString := string(decodedURL)

	return decodedString, nil

}
