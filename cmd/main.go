package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"github.com/surrealdb/surrealdb.go"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"

	"github.com/shijuvar/huma-surreal-example/controller"
	"github.com/shijuvar/huma-surreal-example/surrealdata"
)

// Options for the CLI.
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

func addRoutes(api huma.API, resController *controller.ResourceController) {
	// Register GET /resources
	huma.Register(api, huma.Operation{
		OperationID: "get-resource",
		Method:      http.MethodGet,
		Path:        "/resources",
		Summary:     "Get all resources",
		Description: "Get all resources",
		Tags:        []string{"Resources"},
	}, resController.GetAll)

	// Register GET /resources/{id}
	huma.Register(api, huma.Operation{
		OperationID: "get-resource-id",
		Method:      http.MethodGet,
		Path:        "/resources/{id}",
		Summary:     "Get a resource by its ID",
		Description: "Get a resource",
		Tags:        []string{"Resources"},
	}, resController.GetByID)

	// Register POST /resources
	huma.Register(api, huma.Operation{
		OperationID:   "post-resource",
		Method:        http.MethodPost,
		Path:          "/resources",
		Summary:       "Post a resource",
		Tags:          []string{"Resources"},
		DefaultStatus: http.StatusCreated,
	}, resController.Create)

	// Register DELETE /resources/{id}
	huma.Register(api, huma.Operation{
		OperationID:   "delete-resource",
		Method:        http.MethodDelete,
		Path:          "/resources/{id}",
		Summary:       "Delete a resource by its ID",
		Description:   "Delete a resource",
		Tags:          []string{"Resources"},
		DefaultStatus: http.StatusNoContent,
	}, resController.DeleteByID)

	// Register PUT /resources/{id}
	huma.Register(api, huma.Operation{
		OperationID:   "update-resource",
		Method:        http.MethodPut,
		Path:          "/resources/{id}",
		Summary:       "Update a resource by its ID",
		Description:   "Update a resource",
		Tags:          []string{"Resources"},
		DefaultStatus: http.StatusNoContent,
	}, resController.Update)
}

func main() {
	db := getSurrealDB()
	defer db.Close()
	resourceRepository := surrealdata.NewResourceRespository(db)
	resController := controller.NewResourceController(resourceRepository)

	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create a new router & API
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("ResourceManager API", "1.0.0"))

		addRoutes(api, resController)

		// Tell the CLI how to start your server.
		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			server := &http.Server{
				Addr:    fmt.Sprintf(":%d", options.Port),
				Handler: router,
			}
			server.ListenAndServe()
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}

func getSurrealDB() *surrealdb.DB {
	// Connect to SurrealDB
	db, err := surrealdb.New("ws://localhost:8000")
	if err != nil {
		panic(err)
	}
	// Set the namespace and database
	if err = db.Use("resourceNS", "resourceDB"); err != nil {
		panic(err)
	}

	// Sign in to authentication `db`
	authData := &surrealdb.Auth{
		Username: "root",
		Password: "root",
	}
	_, err = db.SignIn(authData)
	if err != nil {
		panic(err)
	}
	slog.Info("Successfully logged in")
	return db
}
