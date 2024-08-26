package server

import (
	"errors"
	"log"
	"net/http"
	"github.com/Flamiry/books.git/internal/domain/models"
	"github.com/Flamiry/books.git/internal/storage"
	"github.com/gin-gonic/gin"
)

type Storage interface {
	CreateTask(models.Task) error
	UpdateTask(string) (models.Task, error)
	AllTasks()([]models.Task, error)
	DeleteTask(string) error
	TaskInfo(string) (models.Task, error)
}

type Server struct {
	host string
	storage Storage
}

func New(host string, storage Storage) *Server {
	return &Server{
		host:    host,
		storage: storage,
	}
}

func (s *Server) Run() error {
	r := gin.Default()
	taskGroup := r.Group("/tasks")
	{
	taskGroup.GET("/all-tasks", s.AllTasksHandler)
	taskGroup.POST("/add-task", s.TaskCreateHandler)
	taskGroup.PUT("/update-task/:id", s.TaskUpdateHandler)
	taskGroup.DELETE("/delete-task/:id", s.TaskDeleteHandler)
	taskGroup.GET("/task-info/:id", s.TaskInfoHandler)
	}
	if err := r.Run(s.host); err != nil {
		return err
	}
	return nil
}

func (s *Server) AllTasksHandler(ctx *gin.Context) {
	tasks, err := s.storage.AllTasks()
	if err != nil {
		if errors.Is(err, storage.ErrListNotFound) {
			ctx.String(http.StatusNoContent, err.Error())
			return
		}  
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (s *Server) TaskCreateHandler(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindBodyWithJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.storage.CreateTask(task); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (s *Server) TaskUpdateHandler (ctx *gin.Context) {
	tid := ctx.Param("id")
	log.Println(tid)

}

func (s *Server) TaskInfoHandler (ctx *gin.Context) {
	tid := ctx.Param("id")
	log.Println(tid)
	task, err := s.storage.TaskInfo(tid)
	if err != nil {
		if errors.Is(err, storage.ErrTaskNotFound) {
			ctx.String(http.StatusNoContent, err.Error())
			return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"errror": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (s *Server) TaskDeleteHandler (ctx *gin.Context) {
	tid := ctx.Param("id")
	if err := s.storage.DeleteTask(tid); err != nil {
		if errors.Is(err, storage.ErrTaskFailedDelete) {
			ctx.String(http.StatusNoContent, err.Error())
			return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
ctx.String(http.StatusOK,"task deleted")
}