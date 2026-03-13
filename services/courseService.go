package services

import (
	"MainService/entities"
	"MainService/entitiesDTO"
	"MainService/mappers"
	"MainService/repositories"
)

type CourseService interface {
	AddCourseS(course entities.Course) (uint, error)
	GetCoursesS() ([]entitiesDTO.CourseDTO, error)
	GetCourseByIDS(courseID uint) (entitiesDTO.CourseDTO, error)
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
	courseID, err := cs.courseRepository.AddCourse(course)
	if err != nil {
		return 0, err
	}

	return courseID, nil
}

func (cs *courseService) GetCoursesS() ([]entitiesDTO.CourseDTO, error) {
	courses, err := cs.courseRepository.GetCourses()
	if err != nil {
		return []entitiesDTO.CourseDTO{}, err
	}

	coursesDTO := mappers.CoursesToDTO(courses)

	return coursesDTO, nil
}

func (cs *courseService) GetCourseByIDS(courseID uint) (entitiesDTO.CourseDTO, error) {
	course, err := cs.courseRepository.GetCourseByID(courseID)
	if err != nil {
		return entitiesDTO.CourseDTO{}, err
	}

	courseDTO := mappers.CourseToDTO(course)

	return courseDTO, nil
}

func (cs *courseService) DeleteCourseS(courseID uint) error {
	err := cs.courseRepository.DeleteCourse(courseID)
	if err != nil {
		return err
	}

	return nil
}

func (cs *courseService) UpdateCurseS(updCourse entitiesDTO.CourseDTO) error {
	course, err := cs.courseRepository.GetCourseByID(updCourse.ID)
	if err != nil {
		return err
	}

	if updCourse.Name != "" && course.Name != updCourse.Name {
		course.Name = updCourse.Name
	}

	if updCourse.Description != "" && course.Description != updCourse.Description {
		course.Description = updCourse.Description
	}

	errTwo := cs.courseRepository.UpdateCurse(course)
	if errTwo != nil {
		return errTwo
	}

	return nil
}
