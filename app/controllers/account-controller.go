package controllers

import (
	"net/http"

	"github.com/bayuscodings/telloservice"
	"github.com/bayuscodings/telloservice/app/auth"
	"github.com/bayuscodings/telloservice/app/middlewares"
	"github.com/bayuscodings/telloservice/app/models"
	"github.com/bayuscodings/telloservice/app/services"
)

type AccountController struct {
	*BaseController
	App            *telloservice.ApplicationHandler
	AccountService *services.AccountService
}

func NewAccountController(App *telloservice.ApplicationHandler) *AccountController {
	accountService := services.NewAccountService(App.DB)
	return &AccountController{
		BaseController: NewBaseController(),
		App:            App,
		AccountService: accountService,
	}
}

func (ac *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	userPayload := r.Context().Value(middlewares.UserPayloadKey).(*auth.Payload)
	request := r.Context().Value(middlewares.ValidatedRequestKey).(*models.CreateAccountInputDto)

	request.UserID = userPayload.ID

	data, exception := ac.AccountService.CreateAccount(request)
	if exception != nil {
		exception.Respond(w)
		return
	}

	ac.asApiResponse(w, "Account created successfully", data, 201)
}

// FetchAccounts handles the request to fetch accounts with pagination
func (ac *AccountController) FetchAccounts(w http.ResponseWriter, r *http.Request) {
	userPayload := r.Context().Value(middlewares.UserPayloadKey).(*auth.Payload)
	paginationInput := ac.parsePaginationParams(r)

	result, exceptions := ac.AccountService.FetchAccounts(uint(userPayload.ID), &paginationInput)
	if exceptions != nil {
		exceptions.Respond(w)
		return
	}

	ac.asPaginatedApiResponse(
		w, "Accounts retrieved successfully",
		result.Data,
		result.Pagination,
	)
}

func (ac *AccountController) CreateTransFer(w http.ResponseWriter, r *http.Request) {
	userPayload := r.Context().Value(middlewares.UserPayloadKey).(*auth.Payload)
	request := r.Context().Value(middlewares.ValidatedRequestKey).(*models.CreateTransferDto)

	exception := ac.AccountService.TransferFunds(userPayload.ID, request)
	if exception != nil {
		exception.Respond(w)
		return
	}

	ac.asApiResponse(w, "Transfer created successfully", nil)
}
