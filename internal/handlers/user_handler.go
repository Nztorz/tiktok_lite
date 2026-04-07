package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/Nztorz/tiktok_lite/internal/services"
	"github.com/Nztorz/tiktok_lite/internal/utils"
)

type userRequest struct {
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Bio      *string `json:"bio"`
}

type userInputClean struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

type UserHandler struct {
	service services.UserService
	logger  *log.Logger
}

func NewUserHandler(l *log.Logger, service services.UserService) *UserHandler {
	return &UserHandler{
		logger:  l,
		service: service,
	}
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (h *UserHandler) validateRequest(req *userRequest) (userInputClean, error) {
	if req.Email == "" {
		return userInputClean{}, errors.New("email is required")
	}

	if req.Username == "" {
		return userInputClean{}, errors.New("username is required")
	}

	if len(req.Username) > 50 {
		return userInputClean{}, errors.New("username cannot be greater than 50 characters")
	}

	if !emailRegex.MatchString(req.Email) {
		return userInputClean{}, errors.New("invalid email format")
	}

	if req.Password == "" {
		return userInputClean{}, errors.New("password is required")
	}

	bio := ""
	if req.Bio != nil {
		bio = *req.Bio
	}

	userInput := userInputClean{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Bio:      bio,
	}

	return userInput, nil
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// decode req data
	userData, err := utils.ParseJSON[userRequest](r)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, fmt.Sprintf("it was a problem decoding the user's data %v", err))
		h.logger.Println(err)
		return
	}

	// validate the data
	userInput, err := h.validateRequest(&userData)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		h.logger.Println(err)
		return
	}

	// register
	// here the validated data will be
	// processed and register to db
	user, err := h.service.Register(r.Context(), userInput.Email, userInput.Username, userInput.Password, userInput.Bio)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, fmt.Sprintf("there was a problem creating user %v", err))
		h.logger.Println(err)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, user)
}
