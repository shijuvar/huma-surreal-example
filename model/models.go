package model

import (
	"errors"

	"github.com/surrealdb/surrealdb.go/pkg/models"
)

var (
	ErrResourceIDNotFound = errors.New("resource not found for the id")
	ErrResourcesNotFound  = errors.New("resources not found")
)

type Resource struct {
	ID          *models.RecordID `json:"id,omitempty"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Location    string           `json:"location"`
	Category    string           `json:"category"`
	Tags        []string         `json:"tags"`
}

type Repository interface {
	Create(Resource) (string, error)
	Update(string, Resource) error
	GetAll() ([]Resource, error)
	GetByID(string) (Resource, error)
	Delete(string) error
}
