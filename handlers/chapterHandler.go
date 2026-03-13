package handlers

import (
	"MainService/entities"
	"MainService/entitiesDTO"
	"MainService/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChapterHandler interface {
	AddChapterH(c *gin.Context)
	GetChaptersH(c *gin.Context)
	GetChaptersByCourseIDH(c *gin.Context)
	GetChapterByIDH(c *gin.Context)
	DeleteChapterH(c *gin.Context)
	UpdateChapterH(c *gin.Context)
}

type chapterHandler struct {
	chapterService services.ChapterService
}

func NewChapterHandler(chapterService services.ChapterService) ChapterHandler {
	return &chapterHandler{chapterService: chapterService}
}

func (chh *chapterHandler) AddChapterH(c *gin.Context) {
	var chapter entities.Chapter
	if err := c.ShouldBindJSON(&chapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	chapterID, errTwo := chh.chapterService.AddChapterS(chapter)
	if errTwo != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errTwo.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"chapter id": chapterID,
	})
}

func (chh *chapterHandler) GetChaptersH(c *gin.Context) {
	chapters, err := chh.chapterService.GetChaptersS()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, chapters)
}

func (chh *chapterHandler) GetChaptersByCourseIDH(c *gin.Context) {
	strID := c.Param("courseId")

	id64, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	chapters, errTwo := chh.chapterService.GetChaptersByCourseIDS(uint(id64))
	if errTwo != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errTwo.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, chapters)
}

func (chh *chapterHandler) GetChapterByIDH(c *gin.Context) {
	strID := c.Param("id")

	id64, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	chapter, errTwo := chh.chapterService.GetChapterByIDS(uint(id64))
	if errTwo != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errTwo.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, chapter)
}

func (chh *chapterHandler) DeleteChapterH(c *gin.Context) {
	strID := c.Param("id")

	id64, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	errTwo := chh.chapterService.DeleteChapterS(uint(id64))
	if errTwo != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errTwo.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (chh *chapterHandler) UpdateChapterH(c *gin.Context) {
	var chapterDTO entitiesDTO.ChapterDTO
	if err := c.ShouldBindJSON(&chapterDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	errTwo := chh.chapterService.UpdateChapterS(chapterDTO)
	if errTwo != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errTwo.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
