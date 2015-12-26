package tests
import (
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"github.com/bitly/go-simplejson"
	"errors"
	"strconv"
)

const (
	BASE_URL = "http://localhost:8080"
)

type Worker struct {
	client *http.Client
	videos map[string][]map[string]interface{}
}

var w *Worker;

func (this *Worker) DoGet(url string) (*simplejson.Json, error) {
	back, err := this.client.Get(BASE_URL + url);
	if err != nil {
		return nil, err
	}
	json, err := simplejson.NewFromReader(back.Body)
	if (err != nil) {
		return nil, err
	}
	if back.StatusCode != 200 {
		return nil, errors.New("返回状态码: " + strconv.Itoa(back.StatusCode) + " || message:" + json.MustString("msessage"))
	}
	return json, nil;
}

func NewWorker() *Worker {
	return &Worker{client:http.DefaultClient, videos:make(map[string][]map[string]interface{}, 10)}
}


func Md5(formal string) string {
	h := md5.New()
	h.Write([]byte(formal))
	return hex.EncodeToString(h.Sum(nil))
}


func (this *Worker) RefreshVideos() error {
	back, err := w.DoGet("/topinfo")
	if err != nil {
		return err
	}
	if jsonMap, err := back.Map(); err != nil {

	}else {
		for sort, array := range jsonMap {
			if videoArray, ok := array.([]interface{}); ok {
				array := make([]map[string]interface{}, 10)
				for _, li := range videoArray {
					if video, ok := li.(map[string]interface{}); ok {
						array = append(array, video)
					}else {
						return errors.New("Return array is sth wrong")
					}
				}
				this.videos[sort] = array
			}else {
				return errors.New(sort + " is not an array")
			}
		}
	}
	return nil
}


func init() {
	w = NewWorker();
}
