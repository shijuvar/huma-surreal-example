// Package surrealdata provides persistence with SurrealDB for package model
package surrealdata

import (
	surrealdb "github.com/surrealdb/surrealdb.go"
	"github.com/surrealdb/surrealdb.go/pkg/models"

	"github.com/shijuvar/huma-surreal-example/model"
)

type ResourceRepository struct {
	db *surrealdb.DB
}

func NewResourceRepository(db *surrealdb.DB) *ResourceRepository {
	return &ResourceRepository{db: db}
}
func (r ResourceRepository) Create(resource model.Resource) (string, error) {
	result, err := surrealdb.Create[model.Resource](r.db, models.Table("resources"), resource)
	if err != nil {
		return "", err
	}
	id := result.ID.ID.(string) // get the id from RecordID
	return id, nil
}

func (r ResourceRepository) GetAll() ([]model.Resource, error) {
	resources, err := surrealdb.Select[[]model.Resource, models.Table](r.db, models.Table("resources"))
	if err != nil {
		return nil, err
	}
	if len(*resources) == 0 || resources == nil {
		return nil, model.ErrResourcesNotFound
	}
	return *resources, nil
}
func (r ResourceRepository) GetByID(id string) (model.Resource, error) {
	recordID := models.ParseRecordID("resources:" + id)
	resource, err := surrealdb.Select[model.Resource, models.RecordID](r.db, *recordID)
	if err != nil {
		return model.Resource{}, err
	}
	if resource.ID == nil {
		return model.Resource{}, model.ErrResourceIDNotFound
	}
	return *resource, nil
}

func (r ResourceRepository) Delete(id string) error {
	recordID := models.ParseRecordID("resources:" + id)
	result, err := surrealdb.Delete[model.Resource, models.RecordID](r.db, *recordID)
	if err != nil {
		return err
	}
	if result.ID == nil {
		return model.ErrResourceIDNotFound
	}
	return nil
}
func (r ResourceRepository) Update(id string, resource model.Resource) error {
	recordID := models.ParseRecordID("resources:" + id)
	result, err := surrealdb.Update[model.Resource, models.RecordID](r.db, *recordID, resource)
	if err != nil {
		return err
	}
	if result.ID == nil {
		return model.ErrResourceIDNotFound
	}
	return nil
}
