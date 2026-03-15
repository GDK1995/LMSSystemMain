package services

import (
	"MainService/entities"
	"MainService/entitiesDTO"
	"MainService/errorsEntities"
	"MainService/mappers"
	"MainService/repositories"
	"errors"

	"gorm.io/gorm"
)

type LessonService interface {
	AddLessonS(lesson entities.Lesson) (uint, error)
	GetLessonsS() ([]entitiesDTO.LessonDTO, error)
	GetLessonsByChapterIDS(chapterID uint) ([]entitiesDTO.LessonDTO, error)
	GetLessonByIDS(lessonID uint) (*entitiesDTO.LessonDTO, error)
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
		return 0, errorsEntities.ErrInternalServer
	}

	return lessonID, nil
}

func (ls *lessonService) GetLessonsS() ([]entitiesDTO.LessonDTO, error) {
	lessons, err := ls.lessonRepository.GetLessons()
	if err != nil {
		return nil, err
	}

	if len(lessons) == 0 {
		return nil, errorsEntities.ErrLessonNotFound
	}

	lessonDTO := mappers.LessonsToDTO(lessons)

	return lessonDTO, nil
}

func (ls *lessonService) GetLessonsByChapterIDS(chapterID uint) ([]entitiesDTO.LessonDTO, error) {
	lessons, err := ls.lessonRepository.GetLessonsByChapterID(chapterID)
	if err != nil {
		return nil, err
	}

	if len(lessons) == 0 {
		return nil, errorsEntities.ErrLessonNotFound
	}

	lessonsDTO := mappers.LessonsToDTO(lessons)

	return lessonsDTO, nil
}

func (ls *lessonService) GetLessonByIDS(lessonID uint) (*entitiesDTO.LessonDTO, error) {
	lesson, err := ls.lessonRepository.GetLessonByID(lessonID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsEntities.ErrLessonNotFound
		}

		return nil, errorsEntities.ErrInternalServer
	}

	lessonDTO := mappers.LessonToDTO(lesson)

	return &lessonDTO, nil
}

func (ls *lessonService) DeleteLessonS(lessonID uint) error {
	err := ls.lessonRepository.DeleteLesson(lessonID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorsEntities.ErrLessonNotFound
		}

		return errorsEntities.ErrInternalServer
	}
	return nil
}

func (ls *lessonService) UpdateLessonS(updLesson entitiesDTO.LessonDTO) error {
	lesson, err := ls.lessonRepository.GetLessonByID(updLesson.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorsEntities.ErrLessonNotFound
		}

		return errorsEntities.ErrInternalServer
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

	if updLesson.ChapterID != 0 && updLesson.ChapterID != lesson.ChapterID {
		lesson.ChapterID = updLesson.ChapterID
	}

	errTwo := ls.lessonRepository.UpdateLesson(lesson)
	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return errorsEntities.ErrLessonNotFound
		}

		return errorsEntities.ErrInternalServer
	}

	return nil
}
