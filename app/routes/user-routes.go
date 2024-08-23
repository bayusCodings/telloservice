package routes

import (
	"net/http"

	"github.com/bayuscodings/telloservice"
	"github.com/bayuscodings/telloservice/app/controllers"
	"github.com/bayuscodings/telloservice/app/middlewares"
	"github.com/bayuscodings/telloservice/app/models"
	"github.com/bayuscodings/telloservice/app/validators"
)

func UserRoutes(routeBuilder RouteBuilder, App *telloservice.ApplicationHandler) {
	userController := controllers.NewUserController(App)
	apiRouter := routeBuilder.router.PathPrefix("/v1").Subrouter()

	apiRouter.Handle("/user",
		applyMiddlewares(
			http.HandlerFunc(userController.CreateUser),
			validators.ValidateInput(userController.Validator, models.CreateUserInputDto{}),
		),
	).Methods("POST")

	apiRouter.Handle("/user/login",
		applyMiddlewares(
			http.HandlerFunc(userController.LoginUser),
			validators.ValidateInput(userController.Validator, models.UserLoginInputDto{}),
		),
	).Methods("POST")

	apiRouter.Handle("/user/me",
		applyMiddlewares(
			http.HandlerFunc(userController.GetCurrentUser),
			middlewares.AuthMiddleware(App.JWT),
		),
	).Methods("GET")

	apiRouter.Handle("/user/{id}", http.HandlerFunc(userController.GetUserById)).Methods("GET")
}
