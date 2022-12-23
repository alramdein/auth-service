package http

import (
	"context"

	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/alramdein/auth-service/config"
	userPb "github.com/alramdein/user-service/pb"
	log "github.com/sirupsen/logrus"
)

type JWTCustomClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthInfo struct {
	Token string `json:"token"`
}

type AuthHandler struct {
	userClient userPb.UserServiceClient
}

func NewAuthHandler(e *echo.Echo, userClient userPb.UserServiceClient) {
	handler := &AuthHandler{
		userClient: userClient,
	}

	e.POST("/login", handler.Login)
}

func (u *AuthHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid username or password",
		})
	}

	user, err := u.userClient.FindUserByUsernameAndPassword(context.Background(), &userPb.FindUserByUsernameAndPasswordRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "something went wrong",
		})
	}
	if user == nil {
		return c.JSON(http.StatusUnauthorized, &Response{
			Message: "invalid username or password",
		})
	}

	claims := &JWTCustomClaims{
		user.Id,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.JWTSecret()))
	if err != nil {
		log.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "login succesfully",
		Data: AuthInfo{
			Token: t,
		},
	})
}

func WithAuth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(config.JWTSecret()),
	})
}

func ExtractJWTToken(c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*JWTCustomClaims)
}
