package services

import (
	"MainService/entities"
	"MainService/entitiesDTO"
	"MainService/mappers"
	"MainService/repositories"
)

type LessonService interface {
	AddLessonS(lesson entities.Lesson) (uint, error)
	GetLessonsS() ([]entitiesDTO.LessonDTO, error)
	GetLessonsByCourseIDS(courseID uint) ([]entitiesDTO.LessonDTO, error)
	GetLessonByIDS(lessonID uint) (entitiesDTO.LessonDTO, error)
	DeleteLessonS(lessonID uint) error
	UpdateLessonS(updLesson entitiesDTO.LessonDTO) error
}

type lessonService struct {
	lessonRepository repositories.LessonRepository
}

func NewLessonService(lessonRepository repositories.LessonRepository) LessonService {
	return &lessonService{lessonRepository: lessonRepository}
}

func (ls *lessonService) AddLessonS(lesson entities.Lesson) (uint, error) {
	lessonID, err := ls.lessonRepository.AddLesson(lesson)
	if err != nil {
		return 0, err
	}

	return lessonID, nil
}

func (ls *lessonService) GetLessonsS() ([]entitiesDTO.LessonDTO, error) {
	lessons, err := ls.lessonRepository.GetLessons()
	if err != nil {
		return []entitiesDTO.LessonDTO{}, err
	}

	lessonDTO := mappers.LessonsToDTO(lessons)

	return lessonDTO, nil
}

func (ls *lessonService) GetLessonsByCourseIDS(courseID uint) ([]entitiesDTO.LessonDTO, error) {
	lessons, err := ls.lessonRepository.GetLessonsByCourseID(courseID)
	if err != nil {
		return []entitiesDTO.LessonDTO{}, err
	}

	lessonsDTO := mappers.LessonsToDTO(lessons)

	return lessonsDTO, nil
}

func (ls *lessonService) GetLessonByIDS(lessonID uint) (entitiesDTO.LessonDTO, error) {
	lesson, err := ls.lessonRepository.GetLessonByID(lessonID)
	if err != nil {
		return entitiesDTO.LessonDTO{}, err
	}

	lessonDTO := mappers.LessonToDTO(lesson)

	return lessonDTO, nil
}

func (ls *lessonService) DeleteLessonS(lessonID uint) error {
	if err := ls.lessonRepository.DeleteLesson(lessonID); err != nil {
		return err
	}

	return nil
}

func (ls *lessonService) UpdateLessonS(updLesson entitiesDTO.LessonDTO) error {
	lesson, err := ls.lessonRepository.GetLessonByID(updLesson.ID)
	if err != nil {
		return err
	}

	if updLesson.Name != "" && updLesson.Name != lesson.Name {
		lesson.Name = updLesson.Name
	}

	if updLesson.Description != "" && updLesson.Description != lesson.Description {
		lesson.Description = updLesson.Description
	}

	if updLesson.Content != "" && updLesson.Content != lesson.Content {
		lesson.Content = updLesson.Content
	}

	if updLesson.Order != 0 && updLesson.Order != lesson.Order {
		lesson.Order = updLesson.Order
	}

	if updLesson.CourseID != 0 && updLesson.CourseID != lesson.CourseID {
		lesson.CourseID = updLesson.CourseID
	}

	errTwo := ls.lessonRepository.UpdateLesson(lesson)
	if errTwo != nil {
		return errTwo
	}

	return nil
}
