package tool

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

func UrlImgByte(imgUrl string) ([]byte, error) {
	res, err := http.Get(imgUrl)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	// 读取获取的[]byte数据
	data, _ := ioutil.ReadAll(res.Body)

	return data, nil
}

func UrlImgBase64(imgUrl string) (string, error) {
	res, err := http.Get(imgUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// 读取获取的[]byte数据
	data, _ := ioutil.ReadAll(res.Body)

	imageBase64 := base64.StdEncoding.EncodeToString(data)
	return imageBase64, nil
}
