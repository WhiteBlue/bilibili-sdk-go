package client

import (
	"encoding/json"
	"strconv"
)

type RankService struct {
	BaseService
}

type RankVideoElement struct {
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Uri       string `json:"uri"`
	Param     string `json:"param"`
	Goto      string `json:"goto"`
	Name      string `json:"name"`
	Play      int `json:"play"`
	Reply     int `json:"reply"`
	Favourite int `json:"favourite"`
}

/*
	order:
		"view",
		"senddate",
		"reply",
		"danmaku",
		"favorite",
*/
func (r *RankService) SortRank(tid, page, pageSize int, order string) ([]RankVideoElement, error) {
	retBody, err := r.doRequest("http://app.bilibili.com/x/v2/region/show/child/list", map[string]string{
		"build":    "4040",
		"device":      "phone",
		"mobi_app":      "iphone",
		"platform":      "ios",
		"order":    order,
		"pn":     strconv.Itoa(page),
		"ps": strconv.Itoa(pageSize),
		"rid":      strconv.Itoa(tid),
	})
	if err != nil {
		return nil, err
	}
	var ret struct {
		List []RankVideoElement `json:"data"`
	}

	//delete the 'num' key (mdzz)
	json.Unmarshal(retBody, &ret)

	return ret.List, nil
}
