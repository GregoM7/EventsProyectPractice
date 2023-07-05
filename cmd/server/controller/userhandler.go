package controller

import (
	"errors"

	"github.com/GregoM7/EventsProyectPractice/internal/user"
	"github.com/GregoM7/EventsProyectPractice/package/responses"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	ReadAll()
}

type userController struct {
	s user.Service
}

func NewUserController(s user.Service) *userController {
return &userController{s: s}
}

func (h *userController) ReadAll() gin.HandlerFunc {
	return func(c *gin.Context){
		users, _ := h.s.ReadAll()
		if (len(users) == 0) {
			responses.Failure(c, 400, errors.New("There is no Users"))
		}
		responses.Success(c, 200, users)
	}
}