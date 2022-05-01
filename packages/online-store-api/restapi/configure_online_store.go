// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	common "piteroni/online-store-api"
	"piteroni/online-store-api/database"
	"piteroni/online-store-api/restapi/auth"
	"piteroni/online-store-api/restapi/handlers/admin_login_handler"
	"piteroni/online-store-api/restapi/operations"
)

//go:generate swagger generate server --target ../../app --name OnlineStore --spec ../../oas/online-store.json --principal interface{}

func configureFlags(api *operations.OnlineStoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OnlineStoreAPI) http.Handler {
	db, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	logger := common.NewLogger(os.Stdout)

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	authenticator := admin_login_handler.NewAdminAuthenticator(db)

	api.PostAdminLoginHandler = &admin_login_handler.AdminLoginHandler{
		JWTTokenGenerator:  &auth.JWTTokenGenerator{},
		AdminAuthenticator: authenticator,
		AppLogger:          logger,
	}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
