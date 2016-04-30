package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/whiteblue/bilibili-service/lib"
	"strings"
	"net/url"
	"strconv"
	"time"
)

func MakeFailedJsonMap(code string, message string) map[string]string {
	return map[string]string{
		"code":code,
		"message":message,
	}
}

func main() {
	client := NewClient()
	cache, err := NewCache(client)
	if err != nil {
		panic(err)
	}

	//Init schedule
	sche := InitSchedule(2 * time.Hour, cache.FreshCache)
	go sche.Start()
	defer sche.Stop()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//CORS header
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
	})

	//index info
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message":"BiliBili-Service-v3.0"})
	})

	//首页信息
	r.GET("/allrank", func(c *gin.Context) {
		back := cache.GetStaticCache(LABEL_ALL_RANK)
		c.JSON(200, back)
	})

	//视频信息
	r.GET("/view/:aid", func(c *gin.Context) {
		aid := c.Param("aid")
		aidNum, err := strconv.Atoi(aid)
		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM ERROR", err.Error()))
			return
		}
		list, err := cache.GetVideoInfo(aidNum)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("VIDEO_NOT_FOUND", err.Error()))
			return
		}
		c.JSON(200, list)
	})


	//视频源地址
	r.GET("/video/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		quality := c.Query("quality")
		vType := c.Query("type")
		cidNum, err := strconv.Atoi(cid)
		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM ERROR", err.Error()))
			return
		}
		vTypeNum, err := strconv.Atoi(vType)
		if err != nil {
			vTypeNum = 2
		}
		qualityNum, err := strconv.Atoi(quality)
		if err != nil {
			qualityNum = 1
		}
		list, err := cache.GetVideoLink(cidNum, qualityNum, vTypeNum)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("VIDEO_NOT_FOUND", err.Error()))
			return
		}
		c.JSON(200, list)
	})

	//搜索
	r.POST("/search", func(c *gin.Context) {
		content := c.PostForm("content")
		content = strings.Replace(url.QueryEscape(content), "+", "%20", -1)
		page := c.DefaultPostForm("page", "1")
		pageSize := c.DefaultPostForm("count", "20")
		order := c.DefaultPostForm("order", "totalrank");
		searchType := c.DefaultPostForm("type", "all");

		var err error
		pageNum, err := strconv.Atoi(page)
		pageSizeNum, err := strconv.Atoi(pageSize)

		if !strings.EqualFold(content, "")  && err == nil {
			list, err := client.GetSearch(content, pageNum, pageSizeNum, order, searchType)
			if err != nil {
				c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
				return
			}
			c.JSON(200, list)
		}else {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", err.Error()))
		}
	})

	//分类排行
	r.GET("/sort/:tid", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		pageSize := c.DefaultQuery("count", "20")
		tid := c.Param("tid")
		order := c.DefaultQuery("order", "hot")

		var err error
		tidNum, err := strconv.Atoi(tid)
		pageNum, err := strconv.Atoi(page)
		pageSizeNum, err := strconv.Atoi(pageSize)

		if err == nil {
			list, err := client.GetSortRank(tidNum, pageNum, pageSizeNum, order)
			if err != nil {
				c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
				return
			}
			c.JSON(200, list)
		}else {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", err.Error()))
		}
	})

	//专题页面
	r.GET("/spinfo/:spid", func(c *gin.Context) {
		spid := c.Param("spid")

		spidNum, err := strconv.Atoi(spid)

		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", ""))
		}

		list, err := client.GetSpInfo(spidNum)
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
			return
		}
		c.JSON(200, list)
	})


	//专题视频
	r.GET("/spvideos/:spid", func(c *gin.Context) {
		spid := c.Param("spid")
		isBangumi := c.DefaultQuery("bangumi", "0")

		var err error
		spidNum, err := strconv.Atoi(spid)
		isBangumiNum, err := strconv.Atoi(isBangumi)

		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", ""))
		}

		list, err := client.GetSpVideos(spidNum, isBangumiNum)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("SP_NOT_FOUND", err.Error()))
			return
		}
		c.JSON(200, list)
	})


	//新番获取
	r.GET("/bangumi", func(c *gin.Context) {
		back := cache.GetStaticCache(LABEL_BANGUMI_LIST)

		c.JSON(200, back)
	})


	//新番首页推荐
	r.GET("/bangumiindex", func(c *gin.Context) {
		back := cache.GetStaticCache(LABEL_BANGUMI_INDEX)

		c.JSON(200, back)
	})


	//IOS首页
	r.GET("/appindex", func(c *gin.Context) {
		back := cache.GetStaticCache(LABEL_APP_INDEX)

		c.JSON(200, back)
	})

	r.Run(":8080")
}
