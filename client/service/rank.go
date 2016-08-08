package service

import (
	"encoding/json"
	"strconv"
)

type RankService struct {
	BaseService
}

type sortRankResponse struct {
	Name    string                  `json:"name"`
	List    map[string]videoElement `json:"list"`
	Pages   int                     `json:"pages"`
	Results int                     `json:"results"`
}

/*
	order:
		"default",
		"damku",
		"hot",
*/
func (r *RankService) SortRank(tid, page, pageSize int, order string) (*sortRankResponse, error) {
	retBody, err := r.doRequest("http://api.bilibili.com/list", map[string]string{
		"appver":   "2310",
		"build":    "2310",
		"ios":      "0",
		"order":    order,
		"page":     strconv.Itoa(page),
		"pagesize": strconv.Itoa(pageSize),
		"platform": "ios",
		"tid":      strconv.Itoa(tid),
		"type":     "json",
	})
	if err != nil {
		return nil, err
	}
	var ret sortRankResponse

	//delete the 'num' key (mdzz)
	json.Unmarshal(retBody, &ret)

	delete(ret.List, "num")

	return &ret, nil
}
