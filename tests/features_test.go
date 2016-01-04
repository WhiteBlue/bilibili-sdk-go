package tests

import (
	"testing"
	"github.com/whiteblue/bilibili-service/lib"
	"fmt"
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
	r := lib.NewBiliClient()
	back, err := r.GetVideoMp4("5472916","2")
	if err != nil {
		panic(err)
	}
	fmt.Println(back)
}