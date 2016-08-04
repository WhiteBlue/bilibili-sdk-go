package test

import (
	"github.com/whiteblue/bilibili-go/client"
	"github.com/whiteblue/bilibili-go/client/service"
	"strconv"
	"testing"
)

const (
	APPKEY = ""
	SECRET = ""
)

func TestApiSortRank(t *testing.T) {
	c := client.NewClient(APPKEY, SECRET)
	back, err := c.Rank.SortRank(1, 1, 10, "hot")
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	} else {
		length := len(back.List)
		if length == 0 {
			t.Error("return length is 0")
		}
		for i := 0; i < length; i++ {
			index := strconv.Itoa(i)
			if back.List[index].Title == "" {
				t.Error("api return nil")
			}
			t.Log(back.List[index])
		}
	}
}

func TestWeekBangumi(t *testing.T) {
	c := client.NewClient(APPKEY, SECRET)
	back, err := c.Bangumi.GetWeekList("2")
	if err != nil {
		t.Error(err.Error())
		t.Failed()
	} else {
		if len(back.List) == 0 {
			t.Error("return length is 0")
		}
		for _, element := range back.List {
			if element.Title == "" {
				t.Error("api return nil")
			}
			t.Log(element)
		}
	}
}

func TestBangumiIndex(t *testing.T) {
	c := client.NewClient(APPKEY, SECRET)
	back, err := c.Bangumi.GetIndex()
	if err != nil {
		t.Error(err.Error())
		t.Failed()
	} else {
		if len(back.Banners) == 0 {
			t.Error("return banner length is 0")
		}
		for _, banner := range back.Banners {
			if banner.Title == "" {
				t.Error("api return nil")
			}
			t.Log(banner.Title)
		}
		if len(back.Recommends) == 0 {
			t.Error("return recommends length is 0")
		}
		for _, ele := range back.Recommends {
			if ele.Title == "" {
				t.Error("api return nil")
			}
			t.Log(ele)
		}
	}
}

func TestVideoInfo(t *testing.T) {
	c := client.NewClient(APPKEY, SECRET)
	back, err := c.Video.GetVideoInfo(5495647)
	if err != nil {
		t.Error(err.Error())
		t.Failed()
	} else {
		if back.Title == "" {
			t.Error("return title is nil")
		}
		if len(back.PartList) == 0 {
			t.Error("return partlist length is 0")
		}
		t.Log(back)
	}
}

func TestVideoPath(t *testing.T) {
	c := client.NewClient(APPKEY, SECRET)
	back, err := c.Video.GetVideoPartPath("8932442", 1, "mp4")
	if err != nil {
		t.Error(err.Error())
		t.Failed()
	} else {
		if len(back.List) == 0 {
			t.Error("return list length is 0")
		}
		for _, dUrl := range back.List {
			t.Log(dUrl)
		}
	}
}

func TestSearch(t *testing.T) {
	c := client.NewClient(APPKEY, SECRET)
	back, err := c.Others.Search("fate", 1, 10, "totalrank", "all")
	if err != nil {
		t.Error(err.Error())
		t.Failed()
	} else {
		if len(back.PageInfo) == 0 {
			t.Error("return list length is 0")
		}

		for _, video := range back.Result.Videos {
			t.Log(video)
		}
	}
}

func TestSpInfo(t *testing.T) {
	c := client.NewClient(APPKEY, SECRET)
	back, err := c.Special.GetSpecialInfo(158)
	if err != nil {
		t.Error(err.Error())
		t.Failed()
	} else {
		if back.Title == "" {
			t.Error("return title is nil")
		}
		t.Log(back)
	}
}

func TestSpVideos(t *testing.T) {
	c := client.NewClient(APPKEY, SECRET)
	back, err := c.Special.GetSpecialVideos(158, true)
	if err != nil {
		t.Error(err.Error())
		t.Failed()
	} else {
		if len(back.List) == 0 {
			t.Error("return list is nil")
		}
		for _, ele := range back.List {
			t.Log(ele)
		}
	}
}

func TestUserInfo(t *testing.T) {
	c := client.NewClient(APPKEY, SECRET)
	back, err := c.User.GetUserInfo(591635)
	if err != nil {
		t.Error(err.Error())
		t.Failed()
	} else {
		if back.Name == "" {
			t.Error("api return user name nil")
		}
		t.Log(back)
	}
}

func TestUserVideos(t *testing.T) {
	c := client.NewClient(APPKEY, SECRET)
	back, err := c.User.GetUserVideos(591635, 1, 10)
	if err != nil {
		t.Error(err.Error())
		t.Failed()
	} else {
		if len(back.List) == 0 {
			t.Error("api return empty list")
		}
		if len(back.TypeIndex) == 0 {
			t.Error("type list is empty")
		}
		t.Log(back)
	}
}
