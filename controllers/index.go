package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

func IndexGet(c *gin.Context) {
	user, uid := Visitor(c)
	topStory := models.HomeTimeline(uid)

	v := DataState{
		Page:     "index",
		TopStory: topStory,
	}
	dataState, _ := json.Marshal(&v)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"user":      user,
		"topStory":  topStory,
		"dataState": string(dataState),
	})
	return
}
