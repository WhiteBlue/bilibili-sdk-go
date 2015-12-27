package tests

import (
	"testing"
	"github.com/whiteblue/bilibili-service/lib"
)


func TestIndex(t *testing.T) {
	back, err := w.DoGet("/topinfo")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if jsonMap, err := back.Map(); err != nil {
		t.Fatalf(err.Error())
	}else {
		for sort, array := range jsonMap {
			if videoArray, ok := array.([]interface{}); ok {
				t.Log(sort, ":", len(videoArray))
			}else {
				t.Error(sort, "is not an array")
			}
		}
	}
}


func TestGetVideo(t *testing.T) {
	client := lib.NewBiliClient()

	client.GetBangumi("2")
}