package repositories

import (
	"MainService/entities"

	"gorm.io/gorm"
)

type LessonRepository interface {
	AddLesson(lesson entities.Lesson) (uint, error)
	GetLessons() ([]entities.Lesson, error)
	GetLessonsByCourseID(courseID uint) ([]entities.Lesson, error)
	GetLessonByID(lessonID uint) (entities.Lesson, error)
	DeleteLesson(lessonID uint) error
	UpdateLesson(updLesson entities.Lesson) error
}

type lessonRepository struct {
	gormDB *gorm.DB
}

func NewLessonRepository(gormDB *gorm.DB) LessonRepository {
	return &lessonRepository{gormDB: gormDB}
}

func (lr *lessonRepository) AddLesson(lesson entities.Lesson) (uint, error) {
	if err := lr.gormDB.Create(&lesson).Error; err != nil {
		return 0, err
	}

	return lesson.ID, nil
}

func (lr *lessonRepository) GetLessons() ([]entities.Lesson, error) {
	var lessons []entities.Lesson

	if err := lr.gormDB.Find(&lessons).Error; err != nil {
		return []entities.Lesson{}, err
	}

	return lessons, nil
}

func (lr *lessonRepository) GetLessonsByCourseID(courseID uint) ([]entities.Lesson, error) {
	var lessons []entities.Lesson

	if err := lr.gormDB.Where("course_id = ?", courseID).Find(&lessons).Error; err != nil {
		return []entities.Lesson{}, err
	}

	return lessons, nil
}

func (lr *lessonRepository) GetLessonByID(lessonID uint) (entities.Lesson, error) {
	var lesson entities.Lesson

	if err := lr.gormDB.First(&lesson, lessonID).Error; err != nil {
		return entities.Lesson{}, err
	}

	return lesson, nil
}

func (lr *lessonRepository) DeleteLesson(lessonID uint) error {
	if err := lr.gormDB.Delete(&entities.Lesson{}, lessonID).Error; err != nil {
		return err
	}

	return nil
}

func (lr *lessonRepository) UpdateLesson(updLesson entities.Lesson) error {
	if err := lr.gormDB.Save(&updLesson).Error; err != nil {
		return err
	}

	return nil
}
