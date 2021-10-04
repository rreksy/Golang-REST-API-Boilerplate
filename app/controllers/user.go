package controller

import (
	"golang-blueprint/app/helper"
	"golang-blueprint/app/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService user.Service
}

func NewUserController(userService user.Service) *userController {
	return &userController{userService}
}

func (ctrl *userController) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := ctrl.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// JWT NOT SET
// func (ctrl *userController) Login(c *gin.Context) {
// 	var input user.LoginInput

// 	err := c.ShouldBindJSON(&input)
// 	if err != nil {
// 		errors := helper.FormatValidationError(err)
// 		errorsMessage := gin.H{"errors": errors}

// 		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorsMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	loggedInUser, err := ctrl.userService.Login(input)

// 	if err != nil {
// 		errorsMessage := gin.H{"errors": err.Error()}
// 		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorsMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	token, err := ctrl.authService.GenerateToken(loggedInUser.ID)
// 	if err != nil {
// 		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	formatter := user.FormatUser(loggedInUser, token)

// 	response := helper.APIResponse("Logged In Successfully", http.StatusOK, "success", formatter)

// 	c.JSON(http.StatusOK, response)
// }

func (ctrl *userController) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := ctrl.userService.IsEmailAvailable(input)

	if err != nil {
		errorsMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (ctrl *userController) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser)

	response := helper.APIResponse("Successfully fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
