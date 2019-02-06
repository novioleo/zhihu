package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PostQuestion(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var question models.Question
	if err := decoder.Decode(&question); err != nil {
		log.Println("controllers.PostQuestion(): ", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if question.Title == "" || len(question.TopicURLTokens) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
		return
	}
	uid := VisitorID(c)
	if uid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
		return
	}
	if err := models.InsertQuestion(&question, uid); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":    true,
		"questionId": question.ID,
	})
}

func FollowQuestion(c *gin.Context) {
	qid := c.Param("id")
	if qid == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	uid := VisitorID(c)
	if err := models.FollowQuestion(qid, uid); err != nil {
		c.JSON(http.StatusNotFound, nil)
		log.Println("controllers.FollowQuestion(): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"succeed": true,
	})
}

func UnfollowQuestion(c *gin.Context) {
	qid := c.Param("id")
	if qid == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	uid := VisitorID(c)
	if err := models.UnfollowQuestion(qid, uid); err != nil {
		c.JSON(http.StatusNotFound, nil)
		log.Println("controllers.UnfollowQuestion(): ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"succeed": true,
	})
}

func QuestionFollowers(c *gin.Context) {
	qid := c.Param("id")
	if qid == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	uid := VisitorID(c)
	page := &models.Page{
		Session: sessions.Default(c),
	}
	offset, _ := strconv.Atoi(c.Request.FormValue("offset"))
	followers := page.QuestionFollowers(qid, offset, uid)
	if followers == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"paging": page.Paging,
		"data":   followers,
	})
}

func QuestionComments(c *gin.Context) {
	qid := c.Param("id")
	if qid == "" {
		log.Println("controlloers:QuestionComments(): no qestion id")
		c.JSON(http.StatusNotFound, nil)
		return
	}

	uid := VisitorID(c)
	page := &models.Page{
		Session: sessions.Default(c),
	}
	offset, _ := strconv.Atoi(c.Request.FormValue("offset"))
	comments := page.QuestionComments(qid, offset, uid)
	if comments == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"paging": page.Paging,
		"data":   comments,
	})
}

func PostQuestionComment(c *gin.Context) {
	qid := c.Param("id")
	if qid == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	uid := VisitorID(c)
	content := struct {
		Content string `json:"content"`
	}{}
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&content); err != nil {
		log.Println("controllers.PostQuestionComment(): ", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	comment, err := models.InsertQuestionComment(qid, content.Content, uid)
	if err != nil {
		log.Println("controllers.PostQuestionComment(): ", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, comment)
}

func DeleteQuestionComment(c *gin.Context) {
	qid := c.Param("id")
	cid := c.Param("cid")
	if qid == "" || cid == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	uid := VisitorID(c)
	if err := models.DeleteQuestionComment(qid, cid, uid); err != nil {
		log.Println("controllers.DeleteQuestionComment(): ", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func LikeQuestionComment(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	uid := VisitorID(c)
	if err := models.LikeQuestionComment(cid, uid); err != nil {
		log.Println("controllers.LikeQuestionComment(): ", err)
		if err.Error() == "reply is zero" {
			c.JSON(http.StatusBadRequest, nil)
		} else {
			c.JSON(http.StatusInternalServerError, nil)
		}
		return
	}
	c.JSON(http.StatusOK, nil)
}

func UndoLikeQuestionComment(c *gin.Context) {
	cid := c.Param("cid")
	if cid == "" {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	uid := VisitorID(c)
	if err := models.UndoLikeQuestionComment(cid, uid); err != nil {
		log.Println("controllers.UndoLikeQuestionComment(): ", err)
		if err.Error() == "reply is zero" {
			c.JSON(http.StatusBadRequest, nil)
		} else {
			c.JSON(http.StatusInternalServerError, nil)
		}
		return
	}
	c.JSON(http.StatusOK, nil)
}
