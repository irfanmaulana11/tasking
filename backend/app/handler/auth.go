package handler

import (
	"be-tasking/app/service/dto"
	"be-tasking/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ApiHandler) UserRegister(c *gin.Context) {
	reqBody := dto.Register{}
	if err := c.BindJSON(&reqBody); err != nil {
		helper.RespondJSON(c, http.StatusInternalServerError, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: "can't bind request to struct",
			Error:   err.Error(),
		})
		return
	}

	code, err := h.authService.Register(c, reqBody)
	if err != nil {
		helper.RespondJSON(c, code, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: "register failed!",
			Error:   err.Error(),
		})
		return
	}

	helper.RespondJSON(c, code, helper.ResponseWrapper{
		Data:    reqBody,
		Meta:    nil,
		Message: "regis success!",
		Error:   nil,
	})
}

func (h *ApiHandler) UserLogin(c *gin.Context) {
	var (
		reqBody = dto.Login{}
	)

	if err := c.BindJSON(&reqBody); err != nil {
		helper.RespondJSON(c, http.StatusInternalServerError, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: "can't bind request to struct",
			Error:   err.Error(),
		})
		return
	}

	code, respBody, err := h.authService.Login(c, reqBody)
	if err != nil {
		helper.RespondJSON(c, code, helper.ResponseWrapper{
			Data:    nil,
			Meta:    nil,
			Message: http.StatusText(code),
			Error:   err.Error(),
		})
		return
	}

	helper.RespondJSON(c, code, helper.ResponseWrapper{
		Data:    respBody,
		Meta:    nil,
		Message: http.StatusText(code),
		Error:   nil,
	})
}
