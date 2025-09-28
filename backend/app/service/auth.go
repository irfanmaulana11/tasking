package service

import (
	"be-tasking/app/model"
	"be-tasking/app/service/dto"
	"be-tasking/helper"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *authService) Register(c *gin.Context, req dto.Register) (int, error) {
	hashPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user := model.User{
		Name:      req.Name,
		Username:  req.UserName,
		Password:  hashPassword,
		Role:      req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.MySQL.CreateUser(c, user); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *authService) Login(c *gin.Context, login dto.Login) (int, *dto.LoginResponse, error) {
	var (
		res = dto.LoginResponse{}
		exp = time.Now().Add(time.Hour * 20).Unix()
	)

	user, err := s.MySQL.GetUserByUserName(c, login.UserName)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	valid := helper.CheckPassword(login.Password, user.Password)
	if !valid {
		return http.StatusUnauthorized, nil, errors.New("user name and password invalid")
	}

	claim := helper.TokenClaims{
		Sub:         int(user.ID),
		UserName:    user.Username,
		DisplayName: user.Name,
		Role:        user.Role,
		Exp:         exp,
	}

	token, err := helper.CreateToken(claim)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed generating token")
	}

	res.DisplayName = claim.DisplayName
	res.Username = claim.UserName
	res.Role = claim.Role
	res.Expired = claim.Exp
	res.Token = &token

	return http.StatusOK, &res, nil
}
