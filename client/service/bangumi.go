package service

import (
	"encoding/json"
)

type bangumiElement struct {
	Title        string `json:"title"`
	Area         string `json:"area"`
	AreaLimit    int    `json:"arealimit"`
	Attention    int    `json:"attention"`
	BangumiId    int    `json:"bangumi_id"`
	BgmCount     string `json:"bgmcount"`
	Cover        string `json:"cover"`
	SquareCover  string `json:"square_cover"`
	DanmakuCount int    `json:"danmaku_count"`
	Favorites    int    `json:"favorites"`
	IsFinish     int    `json:"is_finish"`
	LastUpdate   string `json:"lastupdate_at"`
	New          bool   `json:"new"`
	PlayCount    int    `json:"play_count"`
	SeasonId     int    `json:"season_id"`
	SpId         int    `json:"spid"`
	Url          string `json:"url"`
	ViewRank     int    `json:"viewRank"`
	Weekday      int    `json:"weekday"`
}

type banner struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Img      string `json:"img"`
	SImg     string `json:"simg"`
	Aid      int    `json:"aid"`
	Type     string `json:"type"`
	Platform int    `json:"platform"`
	Pid      int    `json:"pid"`
}

type recommendBangumiVideo struct {
	Aid         string `json:"aid"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Play        int    `json:"play"`
	Review      int    `json:"review"`
	VideoReview int    `json:"video_review"`
	Favorites   int    `json:"favorites"`
	Mid         int    `json:"mid"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Create      string `json:"create"`
	Pic         string `json:"pic"`
	Coins       int    `json:"coins"`
	Duration    string `json:"duration"`
}

type bangumiActor struct {
	Actor string `json:"actor"`
	Role  string `json:"role"`
}

type bangumiVideo struct {
	Aid        string `json:"av_id"`
	Coins      int    `json:"coins"`
	Cover      string `json:"cover"`
	Danmaku    string `json:"danmaku"`
	Index      string `json:"index"`
	Title      string `json:"index_title"`
	UpdateTime string `json:"update_time"`
}

type bangumiSeason struct {
	Cover        string `json:"cover"`
	IsFinish     string `json:"is_finish"`
	SeasonId     string `json:"season_id"`
	SeasonStatus int `json:"season_status"`
	Title        string `json:"title"`
	TotalCount   string `json:"total_count"`
}

type bangumiInfoResponse struct {
	Actors       []bangumiActor `json:"actor"`
	Alias        string `json:"alias"`
	Area         string `json:"area"`
	BangumiId    string `json:"bangumi_id"`
	BangumiTitle string `json:"bangumi_title"`
	Brief        string `json:"brief"`
	Coins        string `json:"coins"`
	CopyRight    string `json:"copyright"`
	Cover        string `json:"cover"`
	DanmakuCount string `json:"danmaku_count"`
	Episodes     []bangumiVideo `json:"episodes"`
	Evaluate     string `json:"evaluate"`
	Favorites    string `json:"favorites"`
	IsFinish     string `json:"is_finish"`
	JpTitle      string `json:"jp_title"`
	PlayCount    string `json:"play_count"`
	PubTime      string `json:"pub_time"`
	SeasonId     string `json:"season_id"`
	SeasonStatus int `json:"season_status"`
	SeasonTitle  string `json:"season_title"`
	Seasons      []bangumiSeason `json:"seasons"`
	SquareCover  string `json:"squareCover"`
	Staff        string `json:"staff"`
	Title        string `json:"title"`
	TotalCount   string `json:"total_count"`
}

type weekBangumiResponse struct {
	Count string           `json:"count"`
	List  []bangumiElement `json:"list"`
}

type bangumiIndexResponse struct {
	Banners    []banner                `json:"banners"`
	Recommends []recommendBangumiVideo `json:"recommends"`
}

type BangumiService struct {
	BaseService
}

func (b *BangumiService) GetWeekList(bType string) (*weekBangumiResponse, error) {
	retBody, err := b.doRequest("http://app.bilibili.com/bangumi/timeline_v2", map[string]string{
		"_device":  "iphone",
		"btype":    bType,
		"platform": "ios",
		"type":     "json",
	})
	if err != nil {
		return nil, err
	}
	var ret weekBangumiResponse

	json.Unmarshal(retBody, &ret)

	return &ret, nil
}

func (b *BangumiService) GetIndex() (*bangumiIndexResponse, error) {
	retBody, err := b.doRequest("http://app.bilibili.com/api/region_ios/13.json", map[string]string{
		"platform": "ios",
		"device":   "phone",
	})
	if err != nil {
		return nil, err
	}

	var ret struct {
		Content bangumiIndexResponse `json:"result"`
	}

	json.Unmarshal(retBody, &ret)

	return &ret.Content, nil
}

func (b *BangumiService) GetBangumiInfo(seasonId string) (*bangumiInfoResponse, error) {
	retBody, err := b.doRequest("http://bangumi.bilibili.com/api/season_v4", map[string]string{
		"platform":"ios",
		"build":"3940",
		"season_id": seasonId,
		"type":   "bangumi",
	})
	if err != nil {
		return nil, err
	}

	var ret struct {
		Content bangumiInfoResponse `json:"result"`
	}

	json.Unmarshal(retBody, &ret)

	return &ret.Content, nil
}