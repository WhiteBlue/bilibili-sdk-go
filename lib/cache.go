package lib

import (
	"github.com/hashicorp/golang-lru"
	"sync"
	"strings"
	"strconv"
	"fmt"
	"errors"
)

const (
	VIDEO_CACHE_SIZE = 1000
	VIDEO_LINK_CACHE_SIZE = 2000

//番剧首页推荐
	LABEL_BANGUMI_INDEX = "bangumi_index"
//APP首页推荐
	LABEL_APP_INDEX = "app_index"
//番剧列表
	LABEL_BANGUMI_LIST = "bangumi_list"
//热门番剧
	LABEL_BANGUMI_HOT = "app_bangumi_hot"
//分类排行
	LABEL_ALL_RANK = "all_rank"
//总排行
	LABEL_TOP_RANK = "top_rank"
//APP Banner
	LABEL_BANNER = "app_banner"
//APP启动图
	LABEL_START_IMAGE = "app_start_image"
)

type BCache struct {
	cacheMap  map[string]interface{}
	//视频缓存
	videoInfo *lru.Cache
	//视频地址缓存
	videoLink *lru.Cache
	//RW锁
	lock      *sync.RWMutex

	client    *BClient
}

//新建缓存
func NewCache(client *BClient) (*BCache, error) {
	//视频信息
	videoCache, err := lru.New(VIDEO_CACHE_SIZE)
	if err != nil {
		return nil, err
	}
	//视频源
	linkCache, err := lru.New(VIDEO_LINK_CACHE_SIZE)
	if err != nil {
		return nil, err
	}

	cacheMap := make(map[string]interface{})

	lock := new(sync.RWMutex)

	cache := &BCache{cacheMap:cacheMap, videoInfo:videoCache, videoLink:linkCache, client:client, lock:lock}

	flag := cache.FreshCache()

	if flag {
		return cache, nil
	}
	return nil, errors.New("cache init error")
}

//重置缓存
func (this *BCache) FreshCache() bool {
	//Write Lock
	this.lock.Lock()
	defer this.lock.Unlock()

	var back interface{}
	var err error
	if back, err = this.client.GetAllRank(); err == nil {
		this.cacheMap[LABEL_ALL_RANK] = back
	}

	if back, err = this.client.GetAPPIndex(); err == nil {
		this.cacheMap[LABEL_APP_INDEX] = back
	}

	if back, err = this.client.GetBangumi(); err == nil {
		this.cacheMap[LABEL_BANGUMI_LIST] = back
	}

	if back, err = this.client.GetBangumiIndex(); err == nil {
		this.cacheMap[LABEL_BANGUMI_INDEX] = back
	}

	if backMap, err := this.client.GetBannerInfo(); err == nil {
		this.cacheMap[LABEL_BANNER] = backMap["banner"]
		this.cacheMap[LABEL_BANGUMI_HOT] = backMap["bangumi"]
	}

	if back, err := this.client.GetAPPStartImage(); err == nil {
		this.cacheMap[LABEL_START_IMAGE] = back
	}

	if back, err := this.client.GetSortRank(-1, 1, 10, "hot"); err == nil {
		this.cacheMap[LABEL_TOP_RANK] = back
	}


	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (this *BCache) GetVideoInfo(aid int) (map[string]interface{}, error) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if video, ok := this.videoInfo.Get(aid); ok {
		return video.(map[string]interface{}), nil
	}
	back, err := this.client.GetVideoInfo(aid)
	if err != nil {
		return nil, err
	}
	this.videoInfo.Add(aid, back)
	return back, nil
}

func (this *BCache) GetVideoLink(cid int, quality int, vType int) (map[string]interface{}, error) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	cName := strings.Join([]string{strconv.Itoa(cid), strconv.Itoa(quality), strconv.Itoa(vType)}, "_")
	if video, ok := this.videoLink.Get(cName); ok {
		return video.(map[string]interface{}), nil
	}
	back, err := this.client.GetVideo(cid, quality, vType)
	if err != nil {
		return nil, err
	}
	this.videoLink.Add(cName, back)
	return back, nil
}

//Get cache with read lock
func (this *BCache) GetStaticCache(name string) interface{} {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.cacheMap[name]
}

func (this *BCache) SetStaticCache(name string, value interface{}) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	this.cacheMap[name] = value
}