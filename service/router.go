package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func MakeFailedJsonMap(code string, message string) map[string]string {
	return map[string]string{
		"code":    code,
		"message": message,
	}
}

func ConformRoute(app *BiliBiliApplication) {

	app.Router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Access-Control-Max-Age", "7200")
	})

	app.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "BiliBili-Service-v2.0"})
	})

	app.Router.GET("/allrank", func(c *gin.Context) {
		back := app.Cache.GetCache(INDEX_CACHE)
		c.JSON(200, back)
	})

	app.Router.GET("/toprank", func(c *gin.Context) {
		back := app.Cache.GetCache(ALL_RANK_CACHE)
		c.JSON(200, back)
	})

	app.Router.GET("/view/:aid", func(c *gin.Context) {
		aid := c.Param("aid")
		aidNum, err := strconv.Atoi(aid)
		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM ERROR", err.Error()))
			return
		}
		list, err := app.Client.Video.GetVideoInfo(aidNum)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("VIDEO_NOT_FOUND", err.Error()))
			return
		}
		c.JSON(200, list)
	})

	app.Router.GET("/video/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		quality := c.Query("quality")
		vType := c.DefaultQuery("type", "mp4")
		cidNum, err := strconv.Atoi(cid)
		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM ERROR", err.Error()))
			return
		}
		qualityNum, err := strconv.Atoi(quality)
		if err != nil {
			qualityNum = 1
		}
		list, err := app.Client.Video.GetVideoPartPath(cidNum, qualityNum, vType)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("VIDEO_NOT_FOUND", err.Error()))
			return
		}
		c.JSON(200, list)
	})

	app.Router.GET("/user/:mid", func(c *gin.Context) {
		mid := c.Param("mid")
		midNum, err := strconv.Atoi(mid)
		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM ERROR", err.Error()))
			return
		}
		list, err := app.Client.User.GetUserInfo(midNum)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("USER_NOT_FOUND", err.Error()))
			return
		}
		c.JSON(200, list)
	})

	app.Router.GET("/uservideos/:mid", func(c *gin.Context) {
		mid := c.Param("mid")
		page := c.DefaultQuery("page", "1")
		pageSize := c.DefaultQuery("count", "15")

		midNum, err := strconv.Atoi(mid)
		pageNum, err := strconv.Atoi(page)
		pageSizeNum, err := strconv.Atoi(pageSize)

		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM ERROR", err.Error()))
			return
		}
		list, err := app.Client.User.GetUserVideos(midNum, pageNum, pageSizeNum)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("USER_NOT_FOUND", err.Error()))
			return
		}
		c.JSON(200, list)
	})

	app.Router.POST("/search", func(c *gin.Context) {
		content := c.DefaultPostForm("content", "")
		page := c.DefaultPostForm("page", "1")
		pageSize := c.DefaultPostForm("count", "20")
		order := c.DefaultPostForm("order", "totalrank")
		searchType := c.DefaultPostForm("type", "all")

		var err error
		pageNum, err := strconv.Atoi(page)
		pageSizeNum, err := strconv.Atoi(pageSize)

		if strings.TrimSpace(content) == "" || err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", ""))
			return
		}

		list, err := app.Client.Others.Search(content, pageNum, pageSizeNum, order, searchType)
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
			return
		}

		c.JSON(200, list)
	})

	app.Router.GET("/top/:tid", func(c *gin.Context) {
		tid := c.Param("tid")
		var err error
		tidNum, err := strconv.Atoi(tid)
		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", err.Error()))
			return
		}

		cacheName := SORT_TOP_CACHE + strconv.Itoa(tidNum)
		target := app.Cache.GetCache(cacheName)

		if target == nil {
			c.JSON(404, MakeFailedJsonMap("SORT_NOT_FOUND", ""))
			return
		}

		c.JSON(200, target)
	})

	app.Router.GET("/sort/:tid", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		pageSize := c.DefaultQuery("count", "20")
		tid := c.Param("tid")
		order := c.DefaultQuery("order", "hot")

		var err error
		tidNum, err := strconv.Atoi(tid)
		pageNum, err := strconv.Atoi(page)
		pageSizeNum, err := strconv.Atoi(pageSize)

		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", err.Error()))
			return
		}
		list, err := app.Client.Rank.SortRank(tidNum, pageNum, pageSizeNum, order)
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
			return
		}
		c.JSON(200, list)
	})

	app.Router.GET("/spinfo/:spid", func(c *gin.Context) {
		spid := c.Param("spid")
		spidNum, err := strconv.Atoi(spid)

		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", ""))
		}

		list, err := app.Client.Special.GetSpecialInfo(spidNum)
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
			return
		}
		c.JSON(200, list)
	})

	app.Router.GET("/spvideos/:spid", func(c *gin.Context) {
		spid := c.Param("spid")
		isBangumi := c.DefaultQuery("bangumi", "0")

		var err error
		spidNum, err := strconv.Atoi(spid)
		isBangumiNum, err := strconv.Atoi(isBangumi)
		if err != nil {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", ""))
		}

		var isBangumiBool bool
		if isBangumiNum == 1 {
			isBangumiBool = true
		}

		list, err := app.Client.Special.GetSpecialVideos(spidNum, isBangumiBool)
		if err != nil {
			c.JSON(404, MakeFailedJsonMap("SP_NOT_FOUND", err.Error()))
			return
		}
		c.JSON(200, list)
	})

	app.Router.GET("/bangumi", func(c *gin.Context) {
		back := app.Cache.GetCache(BANGUMI_LIST_CACHE)

		c.JSON(200, back)
	})

	app.Router.GET("/bangumiindex", func(c *gin.Context) {
		back := app.Cache.GetCache(BANGUMI_CACHE)

		c.JSON(200, back)
	})

	app.Router.GET("/appindex", func(c *gin.Context) {
		back := app.Cache.GetCache(LIVE_INDEX_CACHE)

		c.JSON(200, back)
	})
}
