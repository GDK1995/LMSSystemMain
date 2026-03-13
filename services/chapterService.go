package services

import (
	"MainService/entities"
	"MainService/entitiesDTO"
	"MainService/mappers"
	"MainService/repositories"
)

type ChapterService interface {
	AddChapterS(chapter entities.Chapter) (uint, error)
	GetChaptersS() ([]entitiesDTO.ChapterDTO, error)
	GetChaptersByCourseIDS(courseID uint) ([]entitiesDTO.ChapterDTO, error)
	GetChapterByIDS(chapterID uint) (entitiesDTO.ChapterDTO, error)
	DeleteChapterS(chapterID uint) error
	UpdateChapterS(updChapter entitiesDTO.ChapterDTO) error
}

type chapterService struct {
	chapterRepository repositories.ChapterRepository
}

func NewChapterService(chapterRepository repositories.ChapterRepository) ChapterService {
	return &chapterService{chapterRepository: chapterRepository}
}

func (chs *chapterService) AddChapterS(chapter entities.Chapter) (uint, error) {
	chapterID, err := chs.chapterRepository.AddChapter(chapter)
	if err != nil {
		return 0, err
	}

	return chapterID, nil
}

func (chs *chapterService) GetChaptersS() ([]entitiesDTO.ChapterDTO, error) {
	chapters, err := chs.chapterRepository.GetChapters()
	if err != nil {
		return []entitiesDTO.ChapterDTO{}, err
	}

	chaptersDTO := mappers.ChaptersToDTO(chapters)

	return chaptersDTO, nil
}

func (chs *chapterService) GetChaptersByCourseIDS(courseID uint) ([]entitiesDTO.ChapterDTO, error) {
	chapters, err := chs.chapterRepository.GetChaptersByCourseID(courseID)
	if err != nil {
		return []entitiesDTO.ChapterDTO{}, err
	}

	chaptersDTO := mappers.ChaptersToDTO(chapters)

	return chaptersDTO, nil
}

func (chs *chapterService) GetChapterByIDS(chapterID uint) (entitiesDTO.ChapterDTO, error) {
	chapter, err := chs.chapterRepository.GetChapterByID(chapterID)
	if err != nil {
		return entitiesDTO.ChapterDTO{}, err
	}

	chapterDTO := mappers.ChapterToDTO(chapter)

	return chapterDTO, nil
}

func (chs *chapterService) DeleteChapterS(chapterID uint) error {
	err := chs.chapterRepository.DeleteChapter(chapterID)
	if err != nil {
		return err
	}
	return nil
}

func (chs *chapterService) UpdateChapterS(updChapter entitiesDTO.ChapterDTO) error {
	chapter, err := chs.chapterRepository.GetChapterByID(updChapter.ID)
	if err != nil {
		return err
	}

	if updChapter.Name != "" && updChapter.Name != chapter.Name {
		chapter.Name = updChapter.Name
	}

	if updChapter.Description != "" && updChapter.Description != chapter.Description {
		chapter.Description = updChapter.Description
	}

	if updChapter.Order != 0 && updChapter.Order != chapter.Order {
		chapter.Order = updChapter.Order
	}

	if updChapter.CourseID != 0 && updChapter.CourseID != chapter.CourseID {
		chapter.CourseID = updChapter.CourseID
	}

	errTwo := chs.chapterRepository.UpdateChapter(chapter)
	if errTwo != nil {
		return errTwo
	}

	return nil
}
