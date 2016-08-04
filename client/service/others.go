package service

import ()
import (
	"strings"
	"strconv"
	"net/url"
	"encoding/json"
)

type videoTypeInfoElement struct {
	Tid   int `json:"tid"`
	Name  string `json:"name"`
	Count int `json:"count"`
}

type searchPageInfo struct {
	Total      int `json:"total"`
	NumResults int `json:"numResults"`
	Pages      int `json:"pages"`
}

type searchBangumiElement struct {
	SeasonId     int `json:"season_id"`
	BangumiId    int `json:"bangumi_id"`
	SpId         int `json:"spid"`
	Title        string `json:"title"`
	Brief        string `json:"brief"`
	Styles       string `json:"styles"`
	Cv           string `json:"cv"`
	Staff        string `json:"staff"`
	Evaluate     string `json:"evaluate"`
	Cover        string `json:"cover"`
	TypeUrl      string `json:"typeurl"`
	Favorites    int `json:"favorites"`
	IsFinish     int `json:"is_finish"`
	PlayCount    int `json:"play_count"`
	DanmakuCount int `json:"danmaku_count"`
	TotalCount   int `json:"total_count"`
}

type searchTopicElement struct {
	TpId        int `json:"tp_id"`
	TpType      int `json:"tp_type"`
	Mid         int `json:"mid"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Arcurl      string `json:"arcurl"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Click       int `json:"click"`
	Review      int `json:"review"`
	Favourite   int `json:"favourite"`
}

type searchResult struct {
	Videos   []videoElement `json:"video"`
	Bangumis []searchBangumiElement `json:"bangumi"`
	Topics   []searchTopicElement `json:"topic"`
}

type searchResponse struct {
	Page     int `json:"page"`
	PageSize int `json:"pagesize"`
	PageInfo map[string]searchPageInfo `json:"pageinfo"`
	Result   searchResult `json:"result"`
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

