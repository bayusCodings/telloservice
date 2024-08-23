package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bayuscodings/telloservice"
	"github.com/bayuscodings/telloservice/app/auth"
	"github.com/bayuscodings/telloservice/app/exceptions"
	"github.com/bayuscodings/telloservice/app/middlewares"
	"github.com/bayuscodings/telloservice/app/models"
	"github.com/bayuscodings/telloservice/app/services"
)

type UserController struct {
	*BaseController
	App         *telloservice.ApplicationHandler
	UserService *services.UserService
}

func NewUserController(App *telloservice.ApplicationHandler) *UserController {
	userService := services.NewUserService(App.DB, App.JWT)
	return &UserController{
		BaseController: NewBaseController(),
		App:            App,
		UserService:    userService,
	}
}

// @Summary      Create a new user
// @Description  Creates a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  models.CreateUserInputDto  true  "User to create"
// @Success      201  {object}  models.ApiResponse[models.UserResponseDto]
// @Router       /user [post]
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := r.Context().Value(middlewares.ValidatedRequestKey).(*models.CreateUserInputDto)

	data, exception := uc.UserService.CreateUser(request)
	if exception != nil {
		exception.Respond(w)
		return
	}

	uc.asApiResponse(w, "User created successfully", data, http.StatusCreated)
}

// @Summary      User login
// @Description  Validates user credentials and log user in
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  models.UserLoginInputDto  true  "User to create"
// @Success      200  {object}  models.ApiResponse[models.LoginResponseDto]
// @Router       /user/login [post]
func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	request := r.Context().Value(middlewares.ValidatedRequestKey).(*models.UserLoginInputDto)

	data, exception := uc.UserService.Login(request)
	if exception != nil {
		exception.Respond(w)
		return
	}

	uc.asApiResponse(w, "User authenticated successfully", data)
}

// @Summary      Get current user
// @Description  Gets the currently logged-in user details
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  models.ApiResponse[models.UserResponseDto]
// @Router       /user/me [get]
func (uc *UserController) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	userPayload := r.Context().Value(middlewares.UserPayloadKey).(*auth.Payload)

	data, exception := uc.UserService.FetchUserById(uint(userPayload.ID))
	if exception != nil {
		exception.Respond(w)
		return
	}

	uc.asApiResponse(w, "User retrieved successfully", data)
}

// @Summary      Get user by ID
// @Description  Retrieve a user's details by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.ApiResponse[models.UserResponseDto]
// @Router       /user/{id} [get]
func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL parameters
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		exception := exceptions.NewBadRequestException("Invalid user ID")
		exception.Respond(w)
		return
	}

	data, exception := uc.UserService.FetchUserById(uint(id))
	if exception != nil {
		exception.Respond(w)
		return
	}

	uc.asApiResponse(w, "User retrieved successfully", data)
}
