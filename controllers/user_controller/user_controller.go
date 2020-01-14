package user_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/navisot/bookstore-users-api/domains/user_model"
	"github.com/navisot/bookstore-users-api/services/users_service"
	"github.com/navisot/bookstore-users-api/utils/errors"
)

func CreateUser(c *gin.Context) {

	var user user_model.User

	if err := c.ShouldBindJSON(&user); err != nil {

		restErr := errors.NewBadRequestError("invalid json body")

		c.JSON(restErr.Status, restErr)

		return
	}

	resultUser, saveErr := users_service.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, resultUser)

}

func GetUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)

	if userErr != nil {
		err := errors.NewBadRequestError("user id must be numeric")
		c.JSON(err.Status, err)
		return
	}

	resultUser, getErr := users_service.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, resultUser)

}
