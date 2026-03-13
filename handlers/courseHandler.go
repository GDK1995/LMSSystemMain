package handlers

import (
	"MainService/entities"
	"MainService/entitiesDTO"
	"MainService/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseHandler interface {
	AddCourseH(c *gin.Context)
	GetCourseH(c *gin.Context)
	GetCourseByIDH(c *gin.Context)
	DeleteCourseH(c *gin.Context)
	UpdateCourseH(c *gin.Context)
}

type courseHandler struct {
	courseService services.CourseService
}

func NewCourseHandler(courseService services.CourseService) CourseHandler {
	return &courseHandler{courseService: courseService}
}

func (ch *courseHandler) AddCourseH(c *gin.Context) {
	var course entities.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	courseID, errTwo := ch.courseService.AddCourseS(course)
	if errTwo != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errTwo.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"course_id": courseID,
	})
}

func (ch *courseHandler) GetCourseH(c *gin.Context) {
	courses, err := ch.courseService.GetCoursesS()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (ch *courseHandler) GetCourseByIDH(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	course, errTwo := ch.courseService.GetCourseByIDS(uint(id))
	if errTwo != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errTwo.Error(),
		})
		return
	}

	c.JSON(200, course)
}

func (ch *courseHandler) DeleteCourseH(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	errTwo := ch.courseService.DeleteCourseS(uint(id))
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

func (ch *courseHandler) UpdateCourseH(c *gin.Context) {
	var courseDTO entitiesDTO.CourseDTO
	err := c.ShouldBindJSON(&courseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	errTwo := ch.courseService.UpdateCurseS(courseDTO)
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
