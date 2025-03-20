package user

import (
	"context"
	"net/http"

	"github.com/LuisGaravaso/goexpert-auction/configs/rest_err"
	"github.com/LuisGaravaso/goexpert-auction/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	usecase user.UserUsecaseInterface
}

func NewUserController(usecase user.UserUsecaseInterface) *UserController {
	return &UserController{usecase: usecase}
}

func (u *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	if err := uuid.Validate(userId); err != nil {
		restErr := rest_err.NewBadRequestError("invalid fields", rest_err.Causes{
			Field:   "user_id",
			Message: "user id must be a valid UUID",
		})
		c.JSON(restErr.Code, restErr)
		return
	}

	input := user.UserInputDTO{ID: userId}
	userData, err := u.usecase.FindUserById(context.Background(), input)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)

}
