package main

import (
	j"encoding/json"
	"strings"
	"errors"
	"strconv"
)

//取得新番信息
func (this *RClient) GetBangumi(btype string) (map[int][]interface{}, error) {
	params := map[string][]string{
		"_device":{"iphone"},
		"btype":{btype},
		"platform":{"ios"},
		"type":{"json"},
	}
	json, err := this.doGet("http://app.bilibili.com/bangumi/timeline_v2?" + DoEncrypt(params))
	if err != nil {
		return nil, err
	}
	array := json.Get("list").MustArray()
	rList := make(map[int][]interface{}, len(array))
	for _, obj := range array {
		if li, ok := obj.(map[string]interface{}); ok {
			if day, ok := li["weekday"].(j.Number); ok {
				if dayInt64, _ := day.Int64(); dayInt64 >= 0 {
					dayInt := int(dayInt64)
					if rList[dayInt] == nil {
						rList[dayInt] = make([]interface{}, 0, 20)
					}
					rList[dayInt] = append(rList[dayInt], li)
				}else {
					if rList[-1] == nil {
						rList[-1] = make([]interface{}, 0, 20)
					}
					rList[-1] = append(rList[-1], li)
				}
			}
		}
	}
	return rList, nil
}

//取得首页内容
func (this *RClient) GetIndex() (map[string][]interface{}, error) {
	json, err := this.doGet("http://api.bilibili.cn/index")
	if err != nil {
		return nil, err
	}
	rList := make(map[string][]interface{}, 10)
	for name, key := range this.sorts {
		if innerMap, err := json.Get(key).Map(); err == nil {
			for order, obj := range innerMap {
				if !strings.EqualFold(order, "num") {
					if _, ok := rList[name]; ok {
						rList[name] = append(rList[name], obj)
					}else {
						rList[name] = make([]interface{}, 0, 9)
						rList[name] = append(rList[name], obj)
					}
				}
			}
		}
	}
	return rList, nil
}

//取得视频详细信息
func (this *RClient) GetVideoInfo(aid string) (map[string]interface{}, error) {
	params := map[string][]string{
		"batch":{"1"},
		"check_area":{"1"},
		"id":{aid},
		"platform":{"ios"},
		"type":{"json"},
	}
	json, err := this.doGet("http://api.bilibili.com/view?" + DoEncrypt(params))
	if err != nil {
		return nil, err
	}
	rMap, err := json.Map()
	if err != nil {
		return rMap, nil
	}
	return nil, err
}


//mp4视频源
func (this *RClient) GetVideoMp4(cid string, quality string) (map[string]interface{}, error) {
	params := map[string][]string{
		"cid":{cid},
		"quality":{quality},
		"otype":{"json"},
		"type":{"mp4"},
	}
	json, err := this.doGet("http://interface.bilibili.com/playurl?" + DoEncrypt(params))
	if err != nil {
		return nil, err
	}
	rMap := make(map[string]interface{})
	if result, err := json.Get("result").String(); err == nil && strings.EqualFold(result, "suee") {
		video := json.Get("durl").MustArray()[0]
		videoObj, _ := video.(map[string]interface{})
		rMap["url"] = videoObj["url"]
		rMap["size"] = videoObj["size"]
		rMap["backup"] = videoObj["backup_url"]
		return rMap, nil
	}else {
		return nil, errors.New("API return error")
	}
}

//flv视频源
func (this *RClient) GetVideoFlv(cid string, quality string) (map[string]interface{}, error) {
	params := map[string][]string{
		"cid":{cid},
		"quality":{quality},
		"otype":{"json"},
		"type":{"flv"},
	}
	json, err := this.doGet("http://interface.bilibili.com/playurl?" + DoEncrypt(params))
	if err != nil {
		return nil, err
	}
	rMap := make(map[string]interface{})
	if result, err := json.Get("result").String(); err == nil && strings.EqualFold(result, "suee") {
		video := json.Get("durl").MustArray()[0]
		videoObj, _ := video.(map[string]interface{})
		rMap["url"] = videoObj["url"]
		rMap["size"] = videoObj["size"]
		rMap["backup"] = videoObj["backup_url"]
		return rMap, nil
	}else {
		return nil, errors.New("API return error")
	}
}


//搜索功能
func (this *RClient) GetSearch(kWord string, page string, count int) (map[string]interface{}, error) {
	params := map[string][]string{
		"keyword":{kWord},
		"page":{page},
		"pagesize":{strconv.Itoa(count)},
	}
	json, err := this.doGet("http://api.bilibili.com/search?" + DoEncrypt(params))
	if err != nil {
		return nil, err
	}
	rMap, err := json.Map()
	if err != nil {
		return nil, err
	}
	return rMap, nil
}


//取得专题视频
func (this *RClient) GetSpVideos(spid string, isBangumi string) (map[string]interface{}, error) {
	params := map[string][]string{
		"spid":{spid},
		"bangumi":{isBangumi},
		"type":{"json"},
	}
	json, err := this.doGet("http://api.bilibili.com/spview?" + IphoneEncrypt(params))
	if err != nil {
		return nil, err
	}
	rMap, err := json.Map()
	if err != nil {
		return nil, err
	}
	return rMap, nil
}


//取得专题信息
func (this *RClient) GetSpInfo(spid string) (map[string]interface{}, error) {
	params := map[string][]string{
		"spid":{spid},
	}
	json, err := this.doGet("http://api.bilibili.cn/sp?" + IphoneEncrypt(params))
	if err != nil {
		return nil, err
	}
	rMap, err := json.Map()
	if err != nil {
		return nil, err
	}
	return rMap, nil
}


