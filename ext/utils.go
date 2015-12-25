package ext

import (
	"strings"
	"sort"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"github.com/bitly/go-simplejson"
	"strconv"
	"errors"
)

const (
	APP_KEY = "876fe0ebd0e67a0f"
	SIGN_KEY = "4ebafd7c4951b366"
	SIGN_SECRET = "8cb98205e9b2ad3669aad0fce12a4c13"
	IPHONE_SECRET = "01db463dd35de8856ad3bba497b18ed8"
)


type RClient struct {
	client *http.Client
	sorts  map[string]string
}

func NewBiliClient() *RClient {
	return &RClient{http.DefaultClient,
		map[string]string{
			"动画":"type1",
			"番剧":"type13",
			"音乐":"type3",
			"舞蹈":"type129",
			"娱乐":"type5",
			"游戏":"type4",
			"科技":"type36",
			"鬼畜":"type119",
			"电影":"type23",
			"电视剧":"type11",
		}}
}


func MakeFailedJsonMap(code string, message string) map[string]string {
	return map[string]string{
		"code":code,
		"message":message,
	}
}

//map=>queryString
func HttpBuildQuery(params map[string][]string) string {
	buffer := make([]string, 0, len(params))
	for key, values := range params {
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
	list := make([]string, 0, len(values))
	sList := make(map[string][]string, len(values))
	for key, _ := range values {
		list = append(list, key)
	}
	sort.Strings(list)
	for _, key := range list {
		sList[key] = values[key]
	}
	queryString := HttpBuildQuery(sList)
	return queryString, Md5(queryString + secret)
}

//拼接QueryString和sign
func DoEncrypt(values map[string][]string) string {
	if (values == nil) {
		values = make(map[string][]string, 1)
	}
	values["appkey"] = []string{SIGN_KEY}
	query, sign := encodeSign(values, SIGN_SECRET)
	return query + "&sign=" + sign
}

//app端的sign加密
func IphoneEncrypt(values map[string][]string) string {
	if (values == nil) {
		values = make(map[string][]string, 1)
	}
	query, sign := encodeSign(values, IPHONE_SECRET)
	return query + "&sign=" + sign
}

//simple Md5 encrypt
func Md5(formal string) string {
	h := md5.New()
	h.Write([]byte(formal))
	return hex.EncodeToString(h.Sum(nil))
}

func (this *RClient) do(req *http.Request) (*simplejson.Json, error) {
	return judgeError(this.client.Do(req))
}


func (this *RClient) doGet(url string) (*simplejson.Json, error) {
	return judgeError(this.client.Get(url))
}

func (this *RClient) doPost(uri string, values map[string][]string) (*simplejson.Json, error) {
	return judgeError(this.client.PostForm(uri, values))
}

//API返回异常判断
func judgeError(resp *http.Response, err error) (*simplejson.Json, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("API return http code = " + strconv.Itoa(resp.StatusCode))
	}
	json, err := simplejson.NewFromReader(resp.Body)
	if (err != nil) {
		return nil, err
	}
	rCode := json.Get("code").MustInt(0)
	if rCode != 0 {
		return nil, errors.New("API return code = " + strconv.Itoa(rCode))
	}
	return json, nil
}


