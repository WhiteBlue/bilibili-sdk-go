package service

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

type videoTypeInfoElement struct {
	Tid   int    `json:"tid"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type searchPageInfo struct {
	Total      int `json:"total"`
	NumResults int `json:"numResults"`
	Pages      int `json:"pages"`
}

type searchVideoElement struct {
	Aid         string `json:"aid"`
	Mid         int    `json:"mid"`
	Copyright   string `json:"copyright"`
	TypeId      int    `json:"typeid"`
	TypeName    string `json:"typename"`
	Title       string `json:"title"`
	Play        int    `json:"play"`
	Review      int    `json:"review"`
	VideoReview int    `json:"video_review"`
	Favorites   int    `json:"favorites"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Create      string `json:"create"`
	Pic         string `json:"pic"`
	Duration    string `json:"duration"`
	Comment     int    `json:"comment"`
	RankScore   int    `json:"rank_score"`
	Tag         string `json:"tag"`
	PubDate     int    `json:"pubdate"`
	SendDate    int    `json:"senddate"`
}

type searchBangumiElement struct {
	SeasonId     int    `json:"season_id"`
	BangumiId    int    `json:"bangumi_id"`
	SpId         int    `json:"spid"`
	Title        string `json:"title"`
	Brief        string `json:"brief"`
	Styles       string `json:"styles"`
	Cv           string `json:"cv"`
	Staff        string `json:"staff"`
	Evaluate     string `json:"evaluate"`
	Cover        string `json:"cover"`
	TypeUrl      string `json:"typeurl"`
	Favorites    int    `json:"favorites"`
	IsFinish     int    `json:"is_finish"`
	PlayCount    int    `json:"play_count"`
	DanmakuCount int    `json:"danmaku_count"`
	TotalCount   int    `json:"total_count"`
}

type searchTopicElement struct {
	TpId        int    `json:"tp_id"`
	TpType      int    `json:"tp_type"`
	Mid         int    `json:"mid"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Arcurl      string `json:"arcurl"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Click       int    `json:"click"`
	Review      int    `json:"review"`
	Favourite   int    `json:"favourite"`
}

type searchResult struct {
	Videos   []searchVideoElement   `json:"video"`
	Bangumis []searchBangumiElement `json:"bangumi"`
	Topics   []searchTopicElement   `json:"topic"`
}

type searchResponse struct {
	Page     int                       `json:"page"`
	PageSize int                       `json:"pagesize"`
	PageInfo map[string]searchPageInfo `json:"pageinfo"`
	Result   searchResult              `json:"result"`
}

type liveBanner struct {
	Title  string `json:"title"`
	Img    string `json:"img"`
	Remark string `json:"remark"`
	Link   string `json:"link"`
}

type liveElement struct {
	User struct {
		Face string `json:"face"`
		Mid  int    `json:"mid"`
		Name string `json:"name"`
	} `json:"owner"`
	Cover struct {
		Src string `json:"src"`
	} `json:"cover"`
	Title         string `json:"title"`
	RoomId        int    `json:"room_id"`
	Online        int    `json:"online"`
	Area          string `json:"area"`
	AreaId        int    `json:"area_id"`
	PlayUrl       string `json:"playurl"`
	AcceptQuality string `json:"accept_quality"`
}

type liveAppIndexResponse struct {
	Banners    []liveBanner `json:"banner"`
	Partitions []struct {
		Partition struct {
			Id      int    `json:"id"`
			Name    string `json:"name"`
			Area    string `json:"area"`
			SubIcon struct {
				Src string `json:"src"`
			} `json:"sub_icon"`
		} `json:"partition"`
		Lives []liveElement `json:"lives"`
	} `json:"partitions"`
	Recommend struct {
		Lives      []liveElement `json:"lives"`
		BannerData []liveElement `json:"banner_data"`
	} `json:"recommend_data"`
}

type OthersService struct {
	BaseService
}

/*
	order:
		"totalrank"
		"click"
		"pubdate"
		"dm"

	searchType:
		"all"
*/
func (o *OthersService) Search(keyword string, page, pageSize int, order, searchType string) (*searchResponse, error) {
	//url raw encode
	keywordEncode := strings.Replace(url.QueryEscape(keyword), "+", "%20", -1)
	retBody, err := o.doRequest("http://api.bilibili.com/search", map[string]string{
		"keyword":     keywordEncode,
		"page":        strconv.Itoa(page),
		"pagesize":    strconv.Itoa(pageSize),
		"device":      "phone",
		"main_ver":    "v3",
		"order":       order,
		"platform":    "ios",
		"search_type": searchType,
		"source_type": "0",
	})
	if err != nil {
		return nil, err
	}

	var ret searchResponse

	json.Unmarshal(retBody, &ret)

	return &ret, nil
}

func (o *OthersService) AppIndex() (*liveAppIndexResponse, error) {
	retBody, err := o.doRequest("http://live.bilibili.com/AppIndex/home", map[string]string{
		"device":    "phone",
		"platform":  "ios",
		"scale":     "2",
		"actionKey": "appkey",
	})
	if err != nil {
		return nil, err
	}

	var ret struct {
		Data liveAppIndexResponse `json:"data"`
	}

	json.Unmarshal(retBody, &ret)

	return &ret.Data, nil
}
