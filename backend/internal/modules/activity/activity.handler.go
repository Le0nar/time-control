package activity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/le0nar/time-control/util"
)

type ActivityHandler struct {
	activityService ActivityService
}

func NewActivityHandler(activityService ActivityService) *ActivityHandler {
	return &ActivityHandler{activityService: activityService}
}

func (h *ActivityHandler) CreateActivity(c *gin.Context) {
	var createActivityDto CreateActivityDto

	if err := c.BindJSON(&createActivityDto); err != nil {
		util.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	activity, err := h.activityService.CreateActivity(createActivityDto)
	if err != nil {
		util.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if activity.WasActive {
		return
		// return 201
	}

	// return 200 with json activity
}