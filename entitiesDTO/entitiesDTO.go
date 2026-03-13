package entitiesDTO

import "MainService/entities"

type CourseDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ChapterDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	CourseID    uint   `json:"course_id"`
	Course      entities.Course
}

type LessonDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Order       int    `json:"order"`
	CourseID    uint   `json:"course_id"`
	Course      entities.Course
}
