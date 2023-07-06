package controller

import (
	"errors"
	"os"

	"github.com/GregoM7/EventsProyectPractice/internal/domain"
	"github.com/GregoM7/EventsProyectPractice/internal/domain/dto"
	"github.com/GregoM7/EventsProyectPractice/internal/user"
	"github.com/GregoM7/EventsProyectPractice/package/responses"
	"github.com/gin-gonic/gin"
)


type userController struct {
	s user.Service
}

func NewUserController(s user.Service) *userController {
	return &userController{s: s}
}

func (h *userController) ReadAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := h.s.ReadAll()
		if len(users) == 0 {
			responses.Failure(c, 400, errors.New("There is no Users"))
		}
		responses.Success(c, 200, users)
	}
}

func (h *userController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			responses.Failure(c, 401, errors.New("Token Not Found"))
			return
		}

		if token != os.Getenv("TOKEN") {
			responses.Failure(c, 401, errors.New("Invalid Token"))
			return
		}
		var userdto dto.UserInsert
		err := c.ShouldBindJSON(&userdto)
		if err != nil {
			responses.Failure(c, 400, errors.New("Invalid User"))
			return
		}
		userFinal := CompleteUser(userdto)
		err = h.s.CreateUser(userFinal)
		if err != nil {
			responses.Failure(c, 400, errors.New(err.Error()))
			return
		}
		responses.Success(c, 201, "User Created")

	}
}

// Mapper
func CompleteUser(dto dto.UserInsert) domain.User {
	userReturn := domain.User{}
	if dto.Role == "ADMIN" {
		userReturn.Role = dto.Role
	} else {
		userReturn.Role = "USER"
	}
	userReturn.Username = dto.Username
	userReturn.Password = dto.Password

	return userReturn
}
