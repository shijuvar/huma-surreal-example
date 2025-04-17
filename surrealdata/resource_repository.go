package surrealdata

import (
	surrealdb "github.com/surrealdb/surrealdb.go"
	"github.com/surrealdb/surrealdb.go/pkg/models"

	"github.com/shijuvar/huma-surreal-example/model"
)

type ResourceRespository struct {
	db *surrealdb.DB
}

func NewResourceRespository(db *surrealdb.DB) *ResourceRespository {
	return &ResourceRespository{db: db}
}
func (r ResourceRespository) Create(resource model.Resource) (string, error) {
	result, err := surrealdb.Create[model.Resource](r.db, models.Table("resources"), resource)
	if err != nil {
		return "", err
	}
	id := result.ID.ID.(string) // get the id from RecordID
	return id, nil
}

func (r ResourceRespository) GetAll() ([]model.Resource, error) {
	resources, err := surrealdb.Select[[]model.Resource, models.Table](r.db, models.Table("resources"))
	if err != nil {
		return nil, err
	}
	return *resources, nil
}
func (r ResourceRespository) GetByID(id string) (model.Resource, error) {
	recordID := models.ParseRecordID("resources:" + id)
	resource, err := surrealdb.Select[model.Resource, models.RecordID](r.db, *recordID)
	if err != nil {
		return model.Resource{}, err
	}
	if resource.ID == nil {
		return model.Resource{}, model.ErrResourceNotFound
	}
	return *resource, nil
}

func (r ResourceRespository) Delete(id string) error {
	recordID := models.ParseRecordID("resources:" + id)
	result, err := surrealdb.Delete[model.Resource, models.RecordID](r.db, *recordID)
	if err != nil {
		return err
	}
	if result.ID == nil {
		return model.ErrResourceNotFound
	}
	return nil
}
func (r ResourceRespository) Update(id string, resource model.Resource) error {
	recordID := models.ParseRecordID("resources:" + id)
	result, err := surrealdb.Update[model.Resource, models.RecordID](r.db, *recordID, resource)
	if err != nil {
		return err
	}
	if result.ID == nil {
		return model.ErrResourceNotFound
	}
	return nil
}
