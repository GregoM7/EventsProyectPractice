package controller

import (
	"errors"
	"fmt"

	"github.com/GregoM7/EventsProyectPractice/internal/event"
	"github.com/GregoM7/EventsProyectPractice/internal/user"
	"github.com/GregoM7/EventsProyectPractice/package/responses"
	"github.com/gin-gonic/gin"
)

type eventController struct {
	s event.Service
	u user.Service
}

func NewEventController(s event.Service, u user.Service) *eventController {
	return &eventController{s: s, u: u}
}

func (h *eventController) ReadAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		username := c.GetHeader("Username")
		user, err := h.u.GetUser(username)
		if err != nil {
			responses.Failure(c, 400, errors.New("Need a user."))
			return
		}
		fmt.Print(user)
		if user.Role == "ADMIN" {
			events, _ := h.s.ReadAllEvents()
			if len(events) == 0 {
				responses.Failure(c, 400, errors.New("There is no Events"))
				return
			}
			responses.Success(c, 200, events)
		} else {
			events, _ := h.s.ReadAllEventsWithState()
			if len(events) == 0 {
				responses.Failure(c, 400, errors.New("There is no Events"))
				return
			}
			responses.Success(c, 200, events)
		}
	}
}
