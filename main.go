package main

import (
	. "./ext"
	"./gin-gonic/gin"
)


type SearchForm struct {
	content    string `form:"content" json:"content" binding:"required"`
	page_count int `form:"page_count" json:"page_count" binding:"required"`
}

func main() {
	client := NewBiliClient()

	gin.SetMode("release")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message":"BiliBili-Html5-v2.0"})
	})

	r.GET("/topinfo", func(c *gin.Context) {
		list, err := client.GetIndex()
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("SERVER_ERROR", err.Error()))
		}
		c.JSON(200, list)
	})

	r.GET("/view/:aid", func(c *gin.Context) {
		aid := c.Param("aid")
		list, err := client.GetVideoInfo(aid)
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("SERVER_ERROR", err.Error()))
		}
		c.JSON(200, list)
	})

	r.GET("/search", func(c *gin.Context) {
		var json SearchForm
		if err := c.BindJSON(&json); err == nil {
			page := c.DefaultQuery("page", "1")
			list, err := client.GetSearch(json.content, page, json.page_count)
			if err != nil {
				c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
			}
			c.JSON(200, list)
		}else {
			c.JSON(400, MakeFailedJsonMap("PARAM_ERROR", err.Error()))
		}
	})

	r.GET("/spinfo/:spid", func(c *gin.Context) {
		spid := c.Param("spid")
		list, err := client.GetSpInfo(spid)
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
		}
		c.JSON(200, list)
	})


	r.GET("/spvideos/:spid", func(c *gin.Context) {
		spid := c.Param("spid")
		isBangumi := c.DefaultQuery("bangumi", "0")
		list, err := client.GetSpVideos(spid, isBangumi)
		if err != nil {
			c.JSON(500, MakeFailedJsonMap("API_RETURN_ERROR", err.Error()))
		}
		c.JSON(200, list)
	})


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
