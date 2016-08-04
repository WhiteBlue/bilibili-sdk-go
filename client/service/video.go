package service

import (
	"encoding/json"
	"strconv"
	"strings"
)

type videoElement struct {
	Aid         string `json:"aid"`
	Mid         int    `json:"mid"`
	Copyright   string `json:"copyright"`
	TypeId      int    `json:"typeid"`
	TypeName    string `json:"typename"`
	Title       string `json:"title"`
	SubTitle    string `json:"subtitle"`
	Play        int    `json:"play"`
	Review      int    `json:"review"`
	VideoReview int    `json:"video_review"`
	Favorites   int    `json:"favorites"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Create      string `json:"create"`
	Pic         string `json:"pic"`
	Credit      int    `json:"credit"`
	Coins       int    `json:"coins"`
	Duration    string `json:"duration"`
	Comment     int    `json:"comment"`
	BadGePay    bool   `json:"badgepay"`
}

type videoMidInfo struct {
	Page int `json:"page"`
	Type string `json:"type"`
	Part string `json:"part"`
	Cid  int `json:"cid"`
	Vid  int `json:"vid"`
}

type videoInfoResponse struct {
	Tid         int `json:"tid"`
	TypeName    string `json:"typename"`
	ArcType     string `json:"arctype"`
	Play        string `json:"play"`
	Review      string `json:"review"`
	VideoReview string `json:"video_review"`
	Favorites   string `json:"favorites"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Pic         string `json:"pic"`
	Author      string `json:"author"`
	Mid         string `json:"mid"`
	AuthorFace  string `json:"face"`
	Pages       int `json:"pages"`
	CreatedAt   string `json:"created_at"`
	Coins       string `json:"coins"`
	PartList    map[string]videoMidInfo `json:"list"`
}

type videoDurl struct {
	Length    int `json:"length"`
	Size      int `json:"size"`
	Url       string `json:"url"`
	BackupUrl []string `json:"backup_url"`
}

type videoPathResponse struct {
	result        string `json:"result"`
	Format        string `json:"format"`
	TimeLength    int `json:"timelength"`
	AcceptFormat  string `json:"accept_format"`
	AcceptQuality []int `json:"accept_quality"`
	List          []videoDurl `json:"durl"`
}

type VideoService struct {
	BaseService
}

func (v *VideoService) GetVideoInfo(aid int) (*videoInfoResponse, error) {
	retBody, err := v.doRequest("http://api.bilibili.com/view", map[string]string{
		"batch":      "1",
		"check_area": "1",
		"id":         strconv.Itoa(aid),
		"platform":   "ios",
		"type":       "json",
	})
	if err != nil {
		return nil, err
	}
	var ret videoInfoResponse

	json.Unmarshal(retBody, &ret)

	return &ret, nil
}


/**
	videoType:
		"flv"
		"hdmp4"
		"mp4"

	quality:
		1,2,3

 */
func (v *VideoService) GetVideoPartPath(cid int, quality int, videoType string) (*videoPathResponse, error) {
	retBody, err := v.doRequest("http://interface.bilibili.com/playurl", map[string]string{
		"cid":     strconv.Itoa(cid),
		"quality": strconv.Itoa(quality),
		"otype":   "json",
		"type":    videoType,
	})
	if err != nil {
		return nil, err
	}
	var ret videoPathResponse

	json.Unmarshal(retBody, &ret)

	if strings.EqualFold(ret.result, "suee") {
		return nil, &ApiError{Msg:"api return error..."}
	}

	return &ret, nil
}