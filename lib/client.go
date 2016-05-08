package lib

import (
	"net/http"
	"strconv"
	"errors"
	"io/ioutil"
	"sort"
	"strings"
)

const (
	APP_KEY = "4ebafd7c4951b366"
	APP_SECRET = "8cb98205e9b2ad3669aad0fce12a4c13"
)

type BClient struct {
	client *http.Client
}

func NewClient() (*BClient) {
	transport := http.Transport{
		DisableKeepAlives: true,
	}

	return &BClient{&http.Client{Transport:&transport}}

}

//API返回异常判断
func judgeError(resp *http.Response, err error) (*JSON, error) {
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("API return http code = " + strconv.Itoa(resp.StatusCode))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body);
	if err != nil {
		return nil, err
	}

	json, err := ParseJSON(bodyBytes)
	if err != nil {
		return nil, err
	}

	//API code
	if rCode, ok := json.Get("code").Int(); ok {
		if rCode != 0 {
			return nil, errors.New("API return code = " + strconv.Itoa(rCode))
		}
	}

	return json, nil
}


//Params map=>queryString
func httpBuildQuery(params map[string][]string) string {
	//对key升序排序
	list := make([]string, 0, len(params))
	buffer := make([]string, 0, len(params))
	for key, _ := range params {
		list = append(list, key)
	}
	sort.Strings(list)
	for _, key := range list {
		values := params[key]
		for _, value := range values {
			buffer = append(buffer, key)
			buffer = append(buffer, "=")
			buffer = append(buffer, value)
			buffer = append(buffer, "&")
		}
	}
	buffer = buffer[:len(buffer) - 1]
	return strings.Join(buffer, "")
}

//B站的sign加密
func encodeSign(values map[string][]string, secret string) (string, string) {
	queryString := httpBuildQuery(values)
	return queryString, Md5(queryString + secret)
}

//拼接QueryString和sign
func doEncrypt(values map[string][]string) string {
	if (values == nil) {
		values = make(map[string][]string, 1)
	}
	values["appkey"] = []string{APP_KEY}
	query, sign := encodeSign(values, APP_SECRET)
	return query + "&sign=" + sign
}

//encrypt for post method
func postEncrypt(values map[string][]string) {
	if (values == nil) {
		return
	}
	_, sign := encodeSign(values, APP_SECRET)
	values["appkey"] = []string{APP_KEY}
	values["sign"] = []string{sign}
}

func (this *BClient) Get(url string, params map[string][]string) (*JSON, error) {
	if params != nil {
		url = url + "?" + doEncrypt(params)
	}
	return judgeError(this.client.Get(url))
}

func (this *BClient) Post(url string, params map[string][]string) (*JSON, error) {
	//postEncrypt(params)
	return judgeError(this.client.PostForm(url, params))
}

