package repositories

import (
	"MainService/entities"

	"gorm.io/gorm"
)

type CourseRepository interface {
	AddCourse(course entities.Course) (uint, error)
	GetCourses() ([]entities.Course, error)
	GetCourseByID(courseID uint) (entities.Course, error)
	DeleteCourse(courseID uint) error
	UpdateCurse(updCourse entities.Course) error
}

type courseRepository struct {
	gormDB *gorm.DB
}

func NewCourseRepository(gormDB *gorm.DB) CourseRepository {
	return &courseRepository{gormDB: gormDB}
}

func (cr *courseRepository) AddCourse(course entities.Course) (uint, error) {
	err := cr.gormDB.Create(&course).Error
	if err != nil {
		return 0, err
	}

	return course.ID, nil
}

func (cr *courseRepository) GetCourses() ([]entities.Course, error) {
	var courseList []entities.Course
	err := cr.gormDB.Find(&courseList).Error
	if err != nil {
		return []entities.Course{}, err
	}

	return courseList, nil
}

func (cr *courseRepository) GetCourseByID(courseID uint) (entities.Course, error) {
	var course entities.Course
	err := cr.gormDB.First(&course, courseID).Error
	if err != nil {
		return entities.Course{}, err
	}

	return course, nil
}

func (cr *courseRepository) DeleteCourse(courseID uint) error {
	err := cr.gormDB.Delete(&entities.Course{}, courseID).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *courseRepository) UpdateCurse(updCourse entities.Course) error {
	err := cr.gormDB.Save(&updCourse).Error
	if err != nil {
		return err
	}

	return nil
}
