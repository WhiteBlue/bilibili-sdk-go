package lib

import (
	"strconv"
	"errors"
	"strings"
)

//B站分类
var sorts = map[string]int{
	"动画":1,
	"番剧":13,
	"音乐":3,
	"舞蹈":129,
	"娱乐":5,
	"游戏":4,
	"科技":36,
	"鬼畜":119,
	"电影":23,
	"电视剧":11,
}


//取得新番信息
func (this *BClient) GetBangumi() (map[string][]map[string]interface{}, error) {
	params := map[string][]string{
		"_device":{"iphone"},
		"btype":{"2"},
		"platform":{"ios"},
		"type":{"json"},
	}
	json, err := this.Get("http://app.bilibili.com/bangumi/timeline_v2", params)
	if err != nil {
		return nil, err
	}

	if movies, ok := json.Get("list").JSONArray(); ok {
		rMap := make(map[string][]map[string]interface{})

		for _, movie := range movies {
			weekdayInt, _ := movie.Get("weekday").Int()
			weekday := strconv.Itoa(weekdayInt)
			if _, ok := rMap[weekday]; !ok {
				rMap[weekday] = make([]map[string]interface{}, 0)
			}
			mJson, _ := movie.Map()
			rMap[weekday] = append(rMap[weekday], mJson)
		}
		return rMap, nil
	}

	return nil, errors.New("API return error..")
}

//取得总排行
func (this *BClient) GetAllRank() (map[string]interface{}, error) {
	json, err := this.Get("http://api.bilibili.cn/index", nil)
	if err != nil {
		return nil, err
	}
	rList := make(map[string]interface{}, 10)
	for name, key := range sorts {
		if innerMap, ok := json.Get("type" + strconv.Itoa(key)).Map(); ok {
			rList[name] = innerMap
		}
	}
	return rList, nil
}


//取得视频详细信息
func (this *BClient) GetVideoInfo(aid int) (map[string]interface{}, error) {
	params := map[string][]string{
		"appver":{"2310"},
		"build":{"2310"},
		"batch":{"1"},
		"check_area":{"1"},
		"id":{strconv.Itoa(aid)},
		"platform":{"ios"},
		"type":{"json"},
	}
	json, err := this.Get("http://api.bilibili.com/view", params)
	if err != nil {
		return nil, err
	}

	if rMap, ok := json.Map(); ok {
		return rMap, nil
	}
	return nil, errors.New("API return error")
}

/**
 * 取得视频源地址
 *
 * videoType:
 * 	flv:0
 *      hdmp4:1
 *	mp4:2
 * quality:
 *	1
 *	2
 *
 */
func (this *BClient) GetVideo(cid int, quality int, videoType int) (map[string]interface{}, error) {
	var vType string
	switch videoType {
	case 0:vType = "flv"
	case 1:vType = "hdmp4"
	default:vType = "mp4"
	}
	params := map[string][]string{
		"cid":{strconv.Itoa(cid)},
		"quality":{strconv.Itoa(quality)},
		"otype":{"json"},
		"type":{vType},
	}
	json, err := this.Get("http://interface.bilibili.com/playurl", params)
	if err != nil {
		return nil, err
	}
	rMap := make(map[string]interface{})
	if result, ok := json.Get("result").String(); ok && strings.EqualFold(result, "suee") {
		videos, _ := json.Get("durl").JSONArray()
		video := videos[0]
		rMap["url"], _ = video.Get("url").String()
		rMap["accept"], _ = json.Get("accept_format").String()
		rMap["length"], _ = json.Get("timelength").Int()
		rMap["backup"], _ = video.Get("backup_url").Map()

		return rMap, nil
	}else {
		return nil, errors.New("API return error")
	}
}


/**
 * 全站搜索
 *
 * 注意keyword需要urlEncode
 */
func (this *BClient) GetSearch(kWord string, page int, pageSize int, order string, searchType string) (map[string]interface{}, error) {
	params := map[string][]string{
		"keyword":{kWord},
		"page":{strconv.Itoa(page)},
		"pagesize":{strconv.Itoa(pageSize)},
		"device":{"phone"},
		"main_ver":{"v3"},
		"order":{order},
		"platform":{"ios"},
		"search_type":{searchType},
		"source_type":{"0"},
	}
	json, err := this.Get("http://api.bilibili.com/search", params)
	if err != nil {
		return nil, err
	}

	returnMap := make(map[string]interface{}, 5)

	if code, ok := json.Get("code").Int(); ok && code == 0 {
		returnMap["page"], _ = json.Get("page").Map()
		returnMap["pagesize"], _ = json.Get("pagesize").Map()
		returnMap["page_info"], _ = json.Get("pageinfo").Map()
		returnMap["result"], _ = json.Get("result").Map()

		return returnMap, nil;
	}else {
		return nil, errors.New("API return error")
	}
}


//取得专题视频
func (this *BClient) GetSpVideos(spid int, isBangumi int) (map[string]interface{}, error) {
	params := map[string][]string{
		"spid":{strconv.Itoa(spid)},
		"bangumi":{strconv.Itoa(isBangumi)},
		"type":{"json"},
	}
	json, err := this.Get("http://api.bilibili.com/spview", params)
	if err != nil {
		return nil, err
	}

	if rMap, ok := json.Map(); ok {
		return rMap, nil
	}

	return nil, errors.New("API return error")
}


//取得专题信息
func (this *BClient) GetSpInfo(spid int) (map[string]interface{}, error) {
	params := map[string][]string{
		"spid":{strconv.Itoa(spid)},
	}
	json, err := this.Get("http://api.bilibili.cn/sp", params)
	if err != nil {
		return nil, err
	}

	if rMap, ok := json.Map(); ok {
		return rMap, nil
	}
	return nil, errors.New("API return error")
}


//读取分类排行
func (this *BClient) GetSortRank(tid int, page int, pageSize int, order string) (map[string]interface{}, error) {
	params := map[string][]string{
		"appver":{"2310"},
		"build":{"2310"},
		"ios":{"0"},
		"order":{order},
		"page":{strconv.Itoa(page)},
		"pagesize":{strconv.Itoa(pageSize)},
		"platform":{"ios"},
		"tid":{strconv.Itoa(tid)},
		"type":{"json"},
	}
	json, err := this.Get("http://api.bilibili.com/list", params)
	if err != nil {
		return nil, err
	}
	if rMap, ok := json.Map(); ok {
		return rMap, nil
	}
	return nil, errors.New("API return error")
}


//取得用户信息
func (this *BClient) GetUserInfo(mid int) (map[string]interface{}, error) {
	params := map[string][]string{
		"mid":{strconv.Itoa(mid)},
	}
	json, err := this.Get("http://api.bilibili.cn/userinfo", params)
	if err != nil {
		return nil, err
	}
	if rMap, ok := json.Map(); ok {
		return rMap, nil
	}
	return nil, errors.New("API return error")
}


//取得用户视频
func (this *BClient) GetUserVideos(mid int, page int, pageSize int) (map[string]interface{}, error) {
	params := map[string][]string{
		"mid":{strconv.Itoa(mid)},
		"page":{strconv.Itoa(page)},
		"pagesize":{strconv.Itoa(pageSize)},
	}
	json, err := this.Get("http://space.bilibili.com/ajax/member/getSubmitVideos", params)
	if err != nil {
		return nil, err
	}
	if rMap, ok := json.Map(); ok {
		return rMap, nil
	}
	return nil, errors.New("API return error")
}


//番剧首页
func (this *BClient) GetBangumiIndex() (map[string]interface{}, error) {
	params := map[string][]string{
		"platform":{"ios"},
		"build":{"2310"},
		"device":{"phone"},
	}
	json, err := this.Get("http://app.bilibili.com/api/region_ios/13.json", params)
	if err != nil {
		return nil, err
	}
	if rMap, ok := json.Map(); ok {
		return rMap, nil
	}
	return nil, errors.New("API return error")
}

//取得banner和番剧推荐(喂这两个为什么要在一起)
func (this *BClient)  GetBannerInfo() (map[string]interface{}, error) {
	json, err := this.Get("http://app.bilibili.com/x/banner/ver?ver=ios4/ver.ios4.@2x.phone", nil)
	if err != nil {
		return nil, err
	}

	gameUrl, _ := json.Get("game").Get("url").String()
	bangumiUrl, _ := json.Get("bangumi").Get("url").String()

	gameJson, err := this.Get(gameUrl, nil)
	bannerJson, err := this.Get("http://www.bilibili.com/index/slideshow.json", nil)
	bangumiJson, err := this.Get(bangumiUrl, nil)

	if err != nil {
		return nil, err
	}

	rMap := make(map[string]interface{})

	rMap["game"], _ = gameJson.Map()
	rMap["banner"], _ = bannerJson.Map()
	rMap["bangumi"], _ = bangumiJson.Map()

	return rMap, nil
}

//APP启动图
func (this *BClient) GetAPPStartImage() (map[string]interface{}, error) {
	params := map[string][]string{
		"build":{"3170"},
		"channel":{"appstore"},
		"height":{"1334"},
		"width":{"750"},
		"plat":{"1"},
	}
	json, err := this.Get("http://app.bilibili.com/x/splash", params)
	if err != nil {
		return nil, err
	}
	if rMap, ok := json.Map(); ok {
		return rMap, nil
	}
	return nil, errors.New("API return error")
}

//APP首页,暂时没想到有什么用orz
func (this *BClient) GetAPPIndex() (map[string]interface{}, error) {
	params := map[string][]string{
		"build":{"3110"},
		"device":{"phone"},
		"platform":{"ios"},
		"scale":{"2"},
		"actionKey":{"appkey"},
	}
	json, err := this.Get("http://live.bilibili.com/AppIndex/home", params)
	if err != nil {
		return nil, err
	}
	if rMap, ok := json.Map(); ok {
		return rMap, nil
	}
	return nil, errors.New("API return error")
}
