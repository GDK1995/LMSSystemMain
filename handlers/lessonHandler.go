package handlers

import (
	"MainService/entities"
	"MainService/entitiesDTO"
	"MainService/errorsEntities"
	"MainService/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LessonHandler interface {
	AddLessonH(c *gin.Context)
	GetLessonsH(c *gin.Context)
	GetLessonsByChapterIDH(c *gin.Context)
	GetLessonByIDH(c *gin.Context)
	DeleteLessonH(c *gin.Context)
	UpdateLessonH(c *gin.Context)
}

type lessonHandler struct {
	lessonService services.LessonService
}

func NewLessonHandler(lessonService services.LessonService) LessonHandler {
	return &lessonHandler{lessonService: lessonService}
}

func (lh *lessonHandler) AddLessonH(c *gin.Context) {
	var lesson entities.Lesson
	if err := c.BindJSON(&lesson); err != nil {
		c.Error(errorsEntities.ErrBadRequest)
		return
	}

	lessonID, errTwo := lh.lessonService.AddLessonS(lesson)
	if errTwo != nil {
		c.Error(errTwo)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"lesson_id": lessonID,
	})
}

func (lh *lessonHandler) GetLessonsH(c *gin.Context) {
	lessons, err := lh.lessonService.GetLessonsS()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, lessons)
}

func (lh *lessonHandler) GetLessonsByChapterIDH(c *gin.Context) {
	strID := c.Param("chapterId")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.Error(errorsEntities.ErrBadRequest)
		return
	}

	lessons, errTwo := lh.lessonService.GetLessonsByChapterIDS(uint(id))
	if errTwo != nil {
		c.Error(errTwo)
		return
	}

	c.JSON(http.StatusOK, lessons)
}

func (lh *lessonHandler) GetLessonByIDH(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.Error(errorsEntities.ErrBadRequest)
		return
	}

	lesson, errTwo := lh.lessonService.GetLessonByIDS(uint(id))
	if errTwo != nil {
		c.Error(errTwo)
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (lh *lessonHandler) DeleteLessonH(c *gin.Context) {
	strID := c.Param("id")
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.Error(errorsEntities.ErrBadRequest)
		return
	}

	errTwo := lh.lessonService.DeleteLessonS(uint(id))
	if errTwo != nil {
		c.Error(errTwo)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (lh *lessonHandler) UpdateLessonH(c *gin.Context) {
	var lessonDTO entitiesDTO.LessonDTO
	if err := c.BindJSON(&lessonDTO); err != nil {
		c.Error(errorsEntities.ErrBadRequest)
		return
	}

	errTwo := lh.lessonService.UpdateLessonS(lessonDTO)
	if errTwo != nil {
		c.Error(errTwo)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
