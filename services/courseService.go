package services

import (
	"MainService/entities"
	"MainService/entitiesDTO"
	"MainService/errorsEntities"
	"MainService/mappers"
	"MainService/repositories"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CourseService interface {
	AddCourseS(course entities.Course) (uint, error)
	GetCoursesS() ([]entitiesDTO.CourseDTO, error)
	GetCourseByIDS(courseID uint) (*entitiesDTO.CourseDTO, error)
	DeleteCourseS(courseID uint) error
	UpdateCurseS(updCourse entitiesDTO.CourseDTO) error
}

type courseService struct {
	courseRepository repositories.CourseRepository
}

func NewCourseService(courseRepository repositories.CourseRepository) CourseService {
	return &courseService{courseRepository: courseRepository}
}

func (cs *courseService) AddCourseS(course entities.Course) (uint, error) {
	logrus.Info("Creating new course")
	logrus.Debugf("Course details: %+v", course)
	courseID, err := cs.courseRepository.AddCourse(course)
	if err != nil {
		logrus.Error("Failed to add course: ", err)
		return 0, errorsEntities.ErrInternalServer
	}

	logrus.Info("Course added successfully")
	return courseID, nil
}

func (cs *courseService) GetCoursesS() ([]entitiesDTO.CourseDTO, error) {
	courses, err := cs.courseRepository.GetCourses()
	if err != nil {
		return nil, errorsEntities.ErrInternalServer
	}

	if len(courses) == 0 {
		return nil, errorsEntities.ErrCourseNotFound
	}

	coursesDTO := mappers.CoursesToDTO(courses)

	return coursesDTO, nil
}

func (cs *courseService) GetCourseByIDS(courseID uint) (*entitiesDTO.CourseDTO, error) {
	course, err := cs.courseRepository.GetCourseByID(courseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorsEntities.ErrCourseNotFound
		}

		return nil, errorsEntities.ErrInternalServer
	}

	courseDTO := mappers.CourseToDTO(course)

	return &courseDTO, nil
}

func (cs *courseService) DeleteCourseS(courseID uint) error {
	err := cs.courseRepository.DeleteCourse(courseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorsEntities.ErrCourseNotFound
		}

		return errorsEntities.ErrInternalServer
	}

	return nil
}

func (cs *courseService) UpdateCurseS(updCourse entitiesDTO.CourseDTO) error {
	course, err := cs.courseRepository.GetCourseByID(updCourse.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorsEntities.ErrCourseNotFound
		}

		return errorsEntities.ErrInternalServer
	}

	if updCourse.Name != "" && course.Name != updCourse.Name {
		course.Name = updCourse.Name
	}

	if updCourse.Description != "" && course.Description != updCourse.Description {
		course.Description = updCourse.Description
	}

	errTwo := cs.courseRepository.UpdateCurse(course)

	if errTwo != nil {
		if errors.Is(errTwo, gorm.ErrRecordNotFound) {
			return errorsEntities.ErrCourseNotFound
		}

		return errorsEntities.ErrInternalServer
	}

	return nil
}
