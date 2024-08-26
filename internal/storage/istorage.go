package storage

import (
	//"fmt"
	//"log"
	"github.com/Flamiry/books.git/internal/domain/models"
	//"github.com/google/uuid"
)

type IStorage struct {
	taskMap map[string]models.Task
}

func New() *IStorage {
	tslice := make(map[string]models.Task)
	return &IStorage{taskMap: tslice}
}
func (s *IStorage) CreateTask(task models.Task) error {
		if _, ok := s.taskMap[task.TID]; !ok {
			return ErrTaskFailedCreate
		}
		s.taskMap[task.TID] = task
		return nil
}

func (s *IStorage) UpdateTask(tId, title, description, status string) (models.Task, error) {
    t, ok := s.taskMap[tId]
    if !ok {
        return models.Task{}, ErrTaskNotFound
    }
    t.Title = title
    t.Description = description
    t.Status = status
    s.taskMap[tId] = t
    return t, nil
}

func (s *IStorage) AllTasks() ([]models.Task, error){
		if len(s.taskMap) == 0 {
			return []models.Task{}, ErrTaskNotFound
	} 
	return []models.Task{}, nil
}

func (s *IStorage) DeleteTask(TID string) error {
	_, ok := s.taskMap[TID]
	if !ok {
		return ErrTaskNotFound
	}
	delete(s.taskMap, TID)
	return nil
}

func (s *IStorage) TaskInfo(TID string) (models.Task, error) {
	task, ok := s.taskMap[TID]
	if !ok {
		return models.Task{}, ErrTaskNotFound
}
	return task, nil
}