package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func StructToReader(v interface{}) io.Reader {
	vByte, _ := json.Marshal(v)
	return bytes.NewReader(vByte)
}

type Request struct {
	Route   string // 接口路径
	Method  string // 接口方法
	Header  map[string]string
	Params  io.Reader   // 入参
	Code    int         // 预期返回值状态码
	Res     interface{} // 出参data数据
	Timeout time.Duration
}

func Run(r *Request) (map[string]interface{}, interface{}, error) {
	client := &http.Client{Timeout: r.Timeout}
	// 利用 httptest 包生成 request
	req, err := http.NewRequest(r.Method, r.Route, r.Params)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	for k, v := range r.Header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal([]byte(string(body)), &result)
	if err != nil {
		return nil, nil, err
	}

	//判断接口调用http状态码是否是200
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http error code: %d", resp.StatusCode)
		return nil, nil, err
	}
	if code, ok := result["code"]; ok {
		//判断接口返回值状态码是否跟预期一致
		if int(code.(float64)) != r.Code {
			err = fmt.Errorf("res error code: %d", int(code.(float64)))
			return nil, nil, err
		}
	}
	if data, ok := result["data"]; ok {
		dataByte, _ := json.Marshal(data)
		json.Unmarshal(dataByte, &r.Res)
	}
	return result, r.Res, nil
}
