// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/yishanhe/animalcrossing-api/pkg/database"
	"github.com/yishanhe/animalcrossing-api/pkg/handlers"
	"github.com/yishanhe/animalcrossing-api/restapi/operations"
	"github.com/yishanhe/animalcrossing-api/restapi/operations/bug"
	"github.com/yishanhe/animalcrossing-api/restapi/operations/fish"
)

//go:generate swagger generate server --target ../../animalcrossing-api --name Acnh --spec ../openapi.yaml

func configureFlags(api *operations.AcnhAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AcnhAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.BugGetBugHandler == nil {
		api.BugGetBugHandler = bug.GetBugHandlerFunc(func(params bug.GetBugParams) middleware.Responder {
			return middleware.NotImplemented("operation bug.GetBug has not yet been implemented")
		})
	}
	if api.FishGetFishHandler == nil {
		api.FishGetFishHandler = fish.GetFishHandlerFunc(func(params fish.GetFishParams) middleware.Responder {
			return middleware.NotImplemented("operation fish.GetFish has not yet been implemented")
		})
	}

	api.BugListBugsHandler = bug.ListBugsHandlerFunc(func(params bug.ListBugsParams) middleware.Responder {
		dbClient := database.NewDatabaseClient()
		return handlers.NewListBugs(dbClient).Handle(params)
	})

	api.FishListFishesHandler = fish.ListFishesHandlerFunc(func(params fish.ListFishesParams) middleware.Responder {
		dbClient := database.NewDatabaseClient()
		return handlers.NewListFishes(dbClient).Handle(params)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
