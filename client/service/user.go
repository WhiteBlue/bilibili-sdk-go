package service

import (
	"strconv"
	"encoding/json"
)

type userVideosResponse struct {
	List      []UserVideoElement `json:"vlist"`
	TypeIndex map[string]videoTypeInfoElement `json:"tlist"`
}

type UserVideoElement struct {
	Aid         int `json:"aid"`
	Copyright   string `json:"copyright"`
	TypeId      int `json:"typeid"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Play        int `json:"play"`
	Review      int `json:"review"`
	VideoReview int `json:"video_review"`
	Favorites   int `json:"favorites"`
	Mid         int `json:"mid"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Created     string `json:"created"`
	Pic         string `json:"pic"`
	Comment     int `json:"comment"`
	Length      string `json:"length"`
}

type userInfoResponse struct {
	Mid          int `json:"mid"`
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	Rank         int `json:"rank"`
	Face         string `json:"face"`
	Coins        float32 `json:"coins"`
	RegTime      int `json:"regtime"`
	Birthday     string `json:"birthday"`
	Place        string `json:"place"`
	Description  string `json:"description"`
	Attentions   []int `json:"attentions"`
	FansNum      int `json:"fans"`
	FriendNum    int `json:"friend"`
	AttentionNum int `json:"attention"`
	Sign         string `json:"sign"`
}

type UserVideoResponse struct {

}

type UserService struct {
	BaseService
}

func (u *UserService) GetUserInfo(mid int) (*userInfoResponse, error) {
	retBody, err := u.doRequest("http://api.bilibili.cn/userinfo", map[string]string{
		"mid": strconv.Itoa(mid),
	})
	if err != nil {
		return nil, err
	}

	var ret userInfoResponse

	json.Unmarshal(retBody, &ret)

	return &ret, nil
}

func (u *UserService) GetUserVideos(mid, page, pageSize int) (*userVideosResponse, error) {
	retBody, err := u.doRequest("http://space.bilibili.com/ajax/member/getSubmitVideos", map[string]string{
		"mid":      strconv.Itoa(mid),
		"page":     strconv.Itoa(page),
		"pagesize": strconv.Itoa(pageSize),
	})
	if err != nil {
		return nil, err
	}

	var ret struct {
		Data userVideosResponse `json:"data"`
	}
	json.Unmarshal(retBody, &ret)

	return &ret.Data, nil
}