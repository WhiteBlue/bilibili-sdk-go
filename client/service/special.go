package service

import (
	"encoding/json"
	"strconv"
)

type specialVideoElement struct {
	Aid   int    `json:"aid"`
	Cid   int    `json:"cid"`
	Cover string `json:"cover"`
	Title string `json:"title"`
	Click int    `json:"click"`
	Page  int    `json:"page"`
}

type specialVideosResponse struct {
	Count   int                   `json:"count"`
	Results int                   `json:"results"`
	List    []specialVideoElement `json:"list"`
}

type specialInfoResponse struct {
	SpId         int    `json:"spid"`
	Title        string `json:"title"`
	CreateAt     string `json:"create_at"`
	UpdateAt     string `json:"lastupdate_at"`
	Alias        string `json:"alias"`
	Cover        string `json:"cover"`
	IsBangumi    int    `json:"isbangumi"`
	IsBangumiEnd int    `json:"isbangumi_end"`
	BangumiDate  string `json:"bangumi_date"`
	Description  string `json:"description"`
	View         int    `json:"view"`
	VideoView    int    `json:"video_view"`
	Favourite    int    `json:"favourite"`
	Attention    int    `json:"attention"`
}

type SpecialService struct {
	BaseService
}

func (s *SpecialService) GetSpecialInfo(spid int) (*specialInfoResponse, error) {
	retBody, err := s.doRequest("http://api.bilibili.cn/sp", map[string]string{
		"spid": strconv.Itoa(spid),
	})
	if err != nil {
		return nil, err
	}

	var ret specialInfoResponse

	json.Unmarshal(retBody, &ret)

	return &ret, nil
}

/*
	isBangumi:
		the result is "bangumi" or other videos
*/
func (s *SpecialService) GetSpecialVideos(spid int, isBangumi bool) (*specialVideosResponse, error) {
	retType := 0
	if isBangumi {
		retType = 1
	}
	retBody, err := s.doRequest("http://api.bilibili.com/spview", map[string]string{
		"spid":    strconv.Itoa(spid),
		"bangumi": strconv.Itoa(retType),
		"type":    "json",
	})
	if err != nil {
		return nil, err
	}

	var ret specialVideosResponse

	json.Unmarshal(retBody, &ret)

	return &ret, nil
}
