package controller

import (
	"github.com/gin-gonic/gin"
	"goblog/service"
	"net/http"
)

func RankByPosts(c *gin.Context) {
	var top int64
	top = 10
	ranks, err := service.RankByPosts(top)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch posts"})
		return
	}

	c.HTML(http.StatusOK, "rank.html", gin.H{
		"ranks": ranks,
	})
}
