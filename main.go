package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/whiteblue/bilibili-service/lib"
	"strings"
	"net/url"
)


func main() {
	client := NewBiliClient()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message":"BiliBili-Html5-v2.0"})
	})

	//首页信息
	r.GET("/topinfo", func(c *gin.Context) {
		list, err := client.GetIndex()
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("SERVER_ERROR", err.Error()))
		}
		c.JSON(200, list)
	})

	//视频信息
	r.GET("/view/:aid", func(c *gin.Context) {
		aid := c.Param("aid")
		list, err := client.GetVideoInfo(aid)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("VIDEO_NOT_FOUND", err.Error()))
		}
		c.JSON(200, list)
	})


	//mp4视频源
	r.GET("/video/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		quailty := c.DefaultQuery("quailty", "1")
		list, err := client.GetVideoMp4(cid, quailty)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("VIDEO_NOT_FOUND", err.Error()))
		}
		c.JSON(200, list)
	})


	//flv视频源
	r.GET("/videoflv/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		quailty := c.DefaultQuery("quailty", "1")
		list, err := client.GetVideoFlv(cid, quailty)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("VIDEO_NOT_FOUND", err.Error()))
		}
		c.JSON(200, list)
	})

	//搜索
	r.GET("/search", func(c *gin.Context) {
		content := c.Query("content")
		//rawurlencode编码
		content = strings.Replace(url.QueryEscape(content), "+", "%20", -1)
		page := c.DefaultQuery("page", "1")
		count := c.DefaultQuery("count", "20")
		if !strings.EqualFold(content, "")  && IsNumber(page) && IsNumber(count) {
			list, err := client.GetSearch(content, page, count)
			if err != nil {
				c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
			}
			c.JSON(200, list)
		}else {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", "request param error.."))
		}
	})

	//分类排行
	r.GET("/sort/:tid", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		count := c.DefaultQuery("count", "20")
		tid := c.Param("tid")
		order := c.DefaultQuery("order", "hot")
		if IsNumber(page)&&IsNumber(count) {
			list, err := client.GetSortInfo(tid, page, count, order)
			if err != nil {
				c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
			}
			c.JSON(200, list)
		}else {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", "request param error.."))
		}
	})

	//专题页面
	r.GET("/spinfo/:spid", func(c *gin.Context) {
		spid := c.Param("spid")
		list, err := client.GetSpInfo(spid)
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
		}
		c.JSON(200, list)
	})


	//专题页面(根据title)
	r.GET("/spinfo", func(c *gin.Context) {
		title := c.Param("title")
		if strings.EqualFold(title, "") {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", "title is nil..."))
		}
		list, err := client.GetSPByName(title)
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
		}
		c.JSON(200, list)
	})


	//专题视频
	r.GET("/spvideos/:spid", func(c *gin.Context) {
		spid := c.Param("spid")
		isBangumi := c.DefaultQuery("bangumi", "0")
		list, err := client.GetSpVideos(spid, isBangumi)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("SP_NOT_FOUND", err.Error()))
		}
		c.JSON(200, list)
	})


	//新番获取
	r.GET("/bangumi", func(c *gin.Context) {
		btype := c.DefaultQuery("btype", "2")
		list, err := client.GetBangumi(btype)
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("SERVER_ERROR", err.Error()))
		}
		c.JSON(200, list)
	})

	r.Run(":8080")
}
