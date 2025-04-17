package controller

import "github.com/shijuvar/huma-surreal-example/model"

// CreateResourceInput represents the request object for HTTP Post /resources
type CreateResourceInput struct {
	Body struct {
		Name        string   `json:"name" doc:"Name of the resource"`
		Description string   `json:"description,omitempty" doc:"Description of the resource"`
		Location    string   `json:"location" doc:"Location of the resource"`
		Category    string   `json:"category" doc:"Category of the resource"`
		Tags        []string `json:"tags,omitempty" doc:"Tags of the resource"`
	}
}

// CreateResourceOutput represents the response object for HTTP Post on /resources
type CreateResourceOutput struct {
	Body struct {
		ID  string `json:"id" doc:"ID of the resource"`
		Err string `json:"err" doc:"Error message for the resource"` // domain error message
	}
}

// UpdateResourceInput represents the request object for HTTP Put on /resources/{id}
type UpdateResourceInput struct {
	ID   string `path:"id" doc:"ID of the resource"`
	Body struct {
		Name        string   `json:"name" doc:"Name of the resource"`
		Description string   `json:"description,omitempty" doc:"Description of the resource"`
		Location    string   `json:"location" doc:"Location of the resource"`
		Category    string   `json:"category" doc:"Category of the resource"`
		Tags        []string `json:"tags,omitempty" doc:"Tags of the resource"`
	}
}

// ResourcesOutput represents the response object for HTTP Get on /resources
type ResourcesOutput struct {
	Body struct {
		Resources []model.Resource `json:"resources" doc:"List of resources"`
		Err       string           `json:"err" doc:"Error message for the resource"`
	}
}

// ResourceIDInput represents the route parameter for HTTP Get and HTTP Delete operations
type ResourceIDInput struct {
	ID string `path:"id" doc:"ID of the resource"`
}

// ResourceByIDOutput represents the response object for HTTP Get on /resources/{id}
type ResourceByIDOutput struct {
	Body struct {
		Resource model.Resource `json:"resource" doc:"Resource of the resource"`
		Err      string         `json:"err" doc:"Error message for the resource"`
	}
}

// ResourceIDOutput represents the response object for both HTTP Put and HTTP Delete on /resources/{id}
type ResourceIDOutput struct {
	Body struct {
		Err string `json:"err" doc:"Error message for the resource"`
	}
}
