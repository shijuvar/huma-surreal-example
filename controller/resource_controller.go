package controller

import (
	"context"
	"errors"

	"github.com/shijuvar/huma-surreal-example/model"
)

type ResourceController struct {
	repository model.Repository
}

func NewResourceController(repository model.Repository) *ResourceController {
	return &ResourceController{
		repository: repository,
	}
}

func (controller *ResourceController) Create(ctx context.Context, input *CreateResourceInput) (*CreateResourceOutput, error) {
	resource := model.Resource{
		Name:        input.Body.Name,
		Description: input.Body.Description,
		Location:    input.Body.Location,
		Category:    input.Body.Category,
		Tags:        input.Body.Tags,
	}
	id, err := controller.repository.Create(resource)
	if err != nil {
		return nil, err
	}
	response := &CreateResourceOutput{}
	response.Body.ID = id
	return response, nil
}

func (controller *ResourceController) GetAll(ctx context.Context, input *struct{}) (*ResourcesOutput, error) {
	resources, err := controller.repository.GetAll()
	if err != nil {
		return nil, err
	}
	response := &ResourcesOutput{}
	response.Body.Resources = resources
	return response, nil
}

func (controller *ResourceController) GetByID(ctx context.Context, input *ResourceIDInput) (*ResourceByIDOutput, error) {
	response := &ResourceByIDOutput{}
	resources, err := controller.repository.GetByID(input.ID)
	if err != nil {
		if errors.Is(err, model.ErrResourceNotFound) {
			response.Body.Err = model.ErrResourceNotFound.Error()
			return response, nil
		}
		return nil, err
	}
	response.Body.Resource = resources
	return response, nil
}

func (controller *ResourceController) DeleteByID(ctx context.Context, input *ResourceIDInput) (*ResourceIDOutput, error) {
	err := controller.repository.Delete(input.ID)
	if err != nil {
		if errors.Is(err, model.ErrResourceNotFound) {
			response := &ResourceIDOutput{}
			response.Body.Err = model.ErrResourceNotFound.Error()
			return response, nil
		}
		return nil, err
	}
	return nil, nil
}

func (controller *ResourceController) Update(ctx context.Context, input *UpdateResourceInput) (*ResourceIDOutput, error) {
	resource := model.Resource{
		Name:        input.Body.Name,
		Description: input.Body.Description,
		Location:    input.Body.Location,
		Category:    input.Body.Category,
		Tags:        input.Body.Tags,
	}
	err := controller.repository.Update(input.ID, resource)
	if err != nil {
		if errors.Is(err, model.ErrResourceNotFound) {
			response := &ResourceIDOutput{}
			response.Body.Err = model.ErrResourceNotFound.Error()
			return response, nil
		}
		return nil, err
	}
	return nil, nil
}
