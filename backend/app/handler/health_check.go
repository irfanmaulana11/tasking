package handler

import (
	"be-tasking/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ApiHandler) Check(c *gin.Context) {
	helper.RespondJSON(c, http.StatusOK, helper.ResponseWrapper{
		Data:    h.healthCheckService.Check(),
		Meta:    nil,
		Message: http.StatusText(http.StatusOK),
		Error:   nil,
	})
}
