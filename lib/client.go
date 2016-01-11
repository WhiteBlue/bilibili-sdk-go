package lib

import (
	"strings"
	"errors"
	j "encoding/json"
	"strconv"
	"net/url"
)

//取得新番信息
func (this *RClient) GetBangumi(btype string) (map[string][]interface{}, error) {
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
	rMap, err := json.Map()
	if err != nil {
		return nil, err
	}

	returnMap := make(map[string][]interface{}, 8)

	if list, ok := rMap["list"].([]interface{}); ok {
		for _, obj := range list {
			innerMap, _ := obj.(map[string]interface{})
			weekdayObj, _ := (innerMap["weekday"]).(j.Number)
			weekday, _ := strconv.Atoi(weekdayObj.String())
			if weekday < 0 {
				weekday = -1
			}
			day := strconv.Itoa(weekday)
			if dayList, ok := returnMap[day]; ok {
				dayList = append(dayList, innerMap)
				returnMap[day] = dayList
			}else {
				dayList = make([]interface{}, 0, 30)
				dayList = append(dayList, innerMap)
				returnMap[day] = dayList
			}
		}
	}
	return returnMap, nil
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
		"appver":{"2310"},
		"build":{"2310"},
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
	if err == nil {
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
		rMap["accept"] = json.Get("accept_format").MustString()
		return rMap, nil
	}else {
		return nil, errors.New("API return error")
	}
}

//flv视频源
func (this *RClient) GetVideoFlv(cid string) (map[string]interface{}, error) {
	params := map[string][]string{
		"cid":{cid},
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
		return rMap, nil
	}else {
		return nil, errors.New("API return error")
	}
}


//搜索功能
func (this *RClient) GetSearch(kWord string, page string, count string, order string, searchType string) (map[string]interface{}, error) {
	params := map[string][]string{
		"keyword":{kWord},
		"page":{page},
		"pagesize":{count},
		"device":{"phone"},
		"main_ver":{"v3"},
		"order":{order},
		"platform":{"ios"},
		"search_type":{searchType},
		"source_type":{"0"},
	}
	json, err := this.doGet("http://api.bilibili.com/search?" + DoEncrypt(params))
	if err != nil {
		return nil, err
	}

	returnMap := make(map[string]interface{}, 5)

	if json.Get("code").MustInt() == 0 {
		returnMap["page"] = json.Get("page")
		returnMap["pagesize"] = json.Get("pagesize")
		returnMap["page_info"] = json.Get("pageinfo")
		returnMap["result"] = json.Get("result")

		return returnMap, nil;
	}else {
		return nil, errors.New("API return code : " + strconv.Itoa(json.Get("code").MustInt()))
	}
}


//取得专题视频
func (this *RClient) GetSpVideos(spid string, isBangumi string) (map[string]interface{}, error) {
	params := map[string][]string{
		"spid":{spid},
		"bangumi":{isBangumi},
		"type":{"json"},
	}
	json, err := this.doGet("http://api.bilibili.com/spview?" + DoEncrypt(params))
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
	json, err := this.doGet("http://api.bilibili.cn/sp?" + DoEncrypt(params))
	if err != nil {
		return nil, err
	}
	rMap, err := json.Map()
	if err != nil {
		return nil, err
	}
	return rMap, nil
}


func (this *RClient) GetSPByName(name string) (map[string]interface{}, error) {
	name = strings.Replace(url.QueryEscape(name), "+", "%20", -1)
	params := map[string][]string{
		"title":{name},
	}
	json, err := this.doGet("http://api.bilibili.cn/sp?" + DoEncrypt(params))
	if err != nil {
		return nil, err
	}
	rMap, err := json.Map()
	if err != nil {
		return nil, err
	}
	return rMap, nil
}


//读取分类排行
func (this *RClient) GetSortInfo(tid string, page string, count string, order string) (map[string]interface{}, error) {
	params := map[string][]string{
		"appver":{"2310"},
		"build":{"2310"},
		"ios":{"0"},
		"order":{order},
		"page":{page},
		"pagesize":{count},
		"platform":{"ios"},
		"tid":{tid},
		"type":{"json"},
	}
	json, err := this.doGet("http://api.bilibili.com/list?" + DoEncrypt(params))
	if err != nil {
		return nil, err
	}
	rMap, err := json.Map()
	if err == nil {
		return rMap, nil
	}
	return nil, err
}


func (this *RClient) GetIndexInfo() (map[string]interface{}, error) {
	params := map[string][]string{
		"platform":{"ios"},
		"build":{"2310"},
		"device":{"phone"},
	}
	json, err := this.doGet("http://app.bilibili.com/api/region_ios/13.json?" + DoEncrypt(params))
	if err != nil {
		return nil, err
	}
	rMap, err := json.Map()
	if err == nil {
		return rMap, nil
	}
	return nil, err
}

