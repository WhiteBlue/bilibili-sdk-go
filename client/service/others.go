package service

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

type videoTypeInfoElement struct {
	Tid   int    `json:"tid"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type searchItem struct {
	Title            string `json:"title"`
	Cover            string `json:"cover"`
	Uri              string `json:"uri"`
	Params           string `json:"param"`
	Goto             string `json:"goto"`
	Desc             string `json:"desc"`

	MovieActors      string `json:"actors"`
	MovieStaff       string `json:"staff"`
	MoviePubDate     string `json:"screen_date"`
	MovieArea        string `json:"area"`
	MovieLength      int    `json:"length"`

	VideoPlay        int    `json:"play"`
	VideoDanmaku     int    `json:"danmaku"`
	VideoAuthor      string `json:"author"`
	VideoDuration    string `json:"duration"`

	SeasonTotalCount int    `json:"total_count"`
	SeasonDesc       string `json:"cat_desc"`

	UserFans         int `json:"fans"`
	UserSign         int `json:"sign"`
}

type searchItems struct {
	Seasons []searchItem `json:"season"`
	Movies  []searchItem  `json:"movie"`
	Vides   []searchItem  `json:"archive"`
}

type searchNavItem struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
	Pages int    `json:"pages"`
	Type  int    `json:"type"`
}

type searchResponse struct {
	Page  int             `json:"page"`
	Navs  []searchNavItem `json:"nav"`
	Items searchItems     `json:"items"`
}

type searchByTypeResponse struct {
	AllPage int `json:"pages"`
	Items   []searchItem `json:"items"`
}

type BannerElement struct {
	Id     int `json:"id"`
	Name   string `json:"name"`
	Pic    string `json:"pic"`
	Url    string `json:"url"`
	PosNum int `json:"pos_num"`
}

type liveBanner struct {
	Title  string `json:"title"`
	Img    string `json:"img"`
	Remark string `json:"remark"`
	Link   string `json:"link"`
}

type liveElement struct {
	User          struct {
			      Face string `json:"face"`
			      Mid  int    `json:"mid"`
			      Name string `json:"name"`
		      } `json:"owner"`
	Cover         struct {
			      Src string `json:"src"`
		      } `json:"cover"`
	Title         string `json:"title"`
	RoomId        int    `json:"room_id"`
	Online        int    `json:"online"`
	Area          string `json:"area"`
	AreaId        int    `json:"area_id"`
	PlayUrl       string `json:"playurl"`
	AcceptQuality string `json:"accept_quality"`
}

type liveAppIndexResponse struct {
	Banners    []liveBanner `json:"banner"`
	Partitions []struct {
		Partition struct {
				  Id      int    `json:"id"`
				  Name    string `json:"name"`
				  Area    string `json:"area"`
				  SubIcon struct {
						  Src string `json:"src"`
					  } `json:"sub_icon"`
			  } `json:"partition"`
		Lives     []liveElement `json:"lives"`
	} `json:"partitions"`
	Recommend  struct {
			   Lives      []liveElement `json:"lives"`
			   BannerData []liveElement `json:"banner_data"`
		   } `json:"recommend_data"`
}

type OthersService struct {
	BaseService
}

/*
	order:
		"totalrank"
		"click"
		"pubdate"
		"dm"

	searchType:
		"all"
*/
func (o *OthersService) Search(keyword string, page, pageSize int, order string) (*searchResponse, error) {
	//url raw encode
	keywordEncode := strings.Replace(url.QueryEscape(keyword), "+", "%20", -1)
	retBody, err := o.doRequest("http://app.bilibili.com/x/v2/search", map[string]string{
		"keyword":     keywordEncode,
		"pn":        strconv.Itoa(page),
		"ps":    strconv.Itoa(pageSize),
		"device":      "phone",
		"main_ver":    "v3",
		"order":       order,
		"platform":    "ios",
	})
	if err != nil {
		return nil, err
	}

	var ret struct {
		Data searchResponse `json:"data"`
	}

	json.Unmarshal(retBody, &ret)

	return &ret.Data, nil
}

func (o *OthersService) SearchByType(keyword string, page, pageSize int, searchType int) (*searchByTypeResponse, error) {
	//url raw encode
	keywordEncode := strings.Replace(url.QueryEscape(keyword), "+", "%20", -1)
	retBody, err := o.doRequest("http://app.bilibili.com/x/v2/search/type", map[string]string{
		"keyword":     keywordEncode,
		"pn":        strconv.Itoa(page),
		"ps":    strconv.Itoa(pageSize),
		"mobi_app":      "iphone",
		"platform":    "ios",
		"device":       "phone",
		"type": strconv.Itoa(searchType),
	})
	if err != nil {
		return nil, err
	}

	var ret struct {
		Data searchByTypeResponse `json:"data"`
	}

	json.Unmarshal(retBody, &ret)

	return &ret.Data, nil
}

func (o *OthersService) AppIndex() (*liveAppIndexResponse, error) {
	retBody, err := o.doRequest("http://live.bilibili.com/AppIndex/home", map[string]string{
		"device":    "phone",
		"platform":  "ios",
		"scale":     "2",
		"actionKey": "appkey",
	})
	if err != nil {
		return nil, err
	}

	var ret struct {
		Data liveAppIndexResponse `json:"data"`
	}

	json.Unmarshal(retBody, &ret)

	return &ret.Data, nil
}

func (o *OthersService) IndexBanner() ([]BannerElement, error) {
	retBody, err := o.doRequest("http://api.bilibili.com/x/web-show/res/loc", map[string]string{
		"jsonp":"jsonp",
		"pf":"0",
		"id":"23",
	})

	if err != nil {
		return nil, err
	}

	var ret struct {
		Data []BannerElement `json:"data"`
	}

	json.Unmarshal(retBody, &ret)

	return ret.Data, nil
}
