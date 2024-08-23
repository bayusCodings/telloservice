package routes

import (
	"net/http"

	"github.com/bayuscodings/telloservice"
	"github.com/bayuscodings/telloservice/app/controllers"
	"github.com/bayuscodings/telloservice/app/middlewares"
	"github.com/bayuscodings/telloservice/app/models"
	"github.com/bayuscodings/telloservice/app/validators"
)

func AccountRoutes(routeBuilder RouteBuilder, App *telloservice.ApplicationHandler) {
	accountController := controllers.NewAccountController(App)
	apiRouter := routeBuilder.router.PathPrefix("/v1").Subrouter()

	apiRouter.Handle("/account",
		applyMiddlewares(
			http.HandlerFunc(accountController.CreateAccount),
			validators.ValidateInput(accountController.Validator, models.CreateAccountInputDto{}),
			middlewares.AuthMiddleware(App.JWT),
		),
	).Methods("POST")

	apiRouter.Handle("/accounts",
		applyMiddlewares(
			http.HandlerFunc(accountController.FetchAccounts),
			middlewares.AuthMiddleware(App.JWT),
		),
	).Methods("GET")

	apiRouter.Handle("/account/transfer",
		applyMiddlewares(
			http.HandlerFunc(accountController.CreateTransFer),
			validators.ValidateInput(accountController.Validator, models.CreateTransferDto{}),
			middlewares.AuthMiddleware(App.JWT),
		),
	).Methods("POST")
}
