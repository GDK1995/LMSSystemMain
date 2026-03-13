package main

import (
	"MainService/handlers"
	"MainService/repositories"
	"MainService/services"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var gormDB *gorm.DB

func InitDB() {
	gormUser := os.Getenv("GORM_USER")
	gormPassword := os.Getenv("GORM_PASSWORD")
	gormName := os.Getenv("GORM_NAME")
	gormHost := os.Getenv("GORM_HOST")
	gormPort := os.Getenv("GORM_PORT")

	connection := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", gormUser, gormPassword, gormName, gormHost, gormPort)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal("Connection error for migrations:", err)
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if errTwo := goose.SetDialect("postgres"); errTwo != nil {
		log.Fatal(errTwo)
	}
	log.Println("Launching migrations...")

	if errThree := goose.Up(db, "migrations"); errThree != nil {
		log.Fatal("Migration execution error:", errThree)
	}
	log.Println("Migrations successfully applied")

	gormDB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Error of GORM:", err)
	}
}
func CloseDB() {
	s, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	errTwo := s.Close()
	if errTwo != nil {
		log.Fatal(errTwo)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}

	InitDB()
	defer CloseDB()

	courseRepo := repositories.NewCourseRepository(gormDB)
	courseServ := services.NewCourseService(courseRepo)
	courseHandler := handlers.NewCourseHandler(courseServ)

	chapterRepo := repositories.NewChapterRepository(gormDB)
	chapterServ := services.NewChapterService(chapterRepo)
	chapterHandler := handlers.NewChapterHandler(chapterServ)

	lessonRepo := repositories.NewLessonRepository(gormDB)
	lessonServ := services.NewLessonService(lessonRepo)
	lessonHandler := handlers.NewLessonHandler(lessonServ)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.POST("/course", courseHandler.AddCourseH)
	router.GET("/course", courseHandler.GetCourseH)
	router.GET("/course/:id", courseHandler.GetCourseByIDH)
	router.DELETE("/course/:id", courseHandler.DeleteCourseH)
	router.PATCH("course", courseHandler.UpdateCourseH)

	router.POST("/chapter", chapterHandler.AddChapterH)
	router.GET("/chapter", chapterHandler.GetChaptersH)
	router.GET("/chapter/course/:courseId", chapterHandler.GetChaptersByCourseIDH)
	router.GET("/chapter/:id", chapterHandler.GetChapterByIDH)
	router.DELETE("/chapter/:id", chapterHandler.DeleteChapterH)
	router.PATCH("/chapter", chapterHandler.UpdateChapterH)

	router.POST("/lesson", lessonHandler.AddLessonH)
	router.GET("/lesson", lessonHandler.GetLessonsH)
	router.GET("/lesson/course/:courseId", lessonHandler.GetLessonsByCourseIDH)
	router.GET("/lesson/:id", lessonHandler.GetLessonByIDH)
	router.DELETE("/lesson/:id", lessonHandler.DeleteLessonH)
	router.PATCH("/lesson", lessonHandler.UpdateLessonH)

	router.Run(":8083")
}
