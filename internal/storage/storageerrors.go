package storage

import (
	"errors"
	errtext "github.com/Flamiry/books.git/internal/domain/errors"
)

var ErrListNotFound = errors.New(errtext.ListNotFoundError)
var ErrTaskNotFound = errors.New(errtext.TaskNotFoundError)
var ErrTaskFailedCreate = errors.New(errtext.TaskFailedCreateError)
var ErrTaskFailedUpdate = errors.New(errtext.TaskFailedUpdateError) 
var ErrTaskFailedDelete = errors.New(errtext.TaskFailedDeleteError)
