package routes

import (
	"fmt"
	"net/http"

	"github.com/bayuscodings/telloservice"
	_ "github.com/bayuscodings/telloservice/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type RouteBuilder struct {
	router *mux.Router
}

func (routeBuilder *RouteBuilder) MakeRoute(path string, f func(RouteBuilder, *mux.Router)) *RouteBuilder {
	apiRouter := routeBuilder.router.PathPrefix(path).Subrouter()
	f(RouteBuilder{router: apiRouter}, apiRouter)
	return routeBuilder
}

func BuildRoute(App *telloservice.ApplicationHandler) *mux.Router {
	mainRouteBuilder := &RouteBuilder{router: mux.NewRouter()}

	ServeSwagger(mainRouteBuilder.router)
	mainRouteBuilder.MakeRoute("/", func(apiRouteBuilder RouteBuilder, router *mux.Router) {
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Running TelloService...\n")
		})

		UserRoutes(apiRouteBuilder, App)
		AccountRoutes(apiRouteBuilder, App)
	})

	return mainRouteBuilder.router
}

func ServeSwagger(router *mux.Router) {
	// Serve the embedded Swagger UI
	router.PathPrefix("/api-docs/").Handler(httpSwagger.WrapHandler)
}

// applyMiddlewares applies a list of middleware to a given handler
func applyMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
