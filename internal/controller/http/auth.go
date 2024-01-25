package http

import (
	"gotest/internal/core/common/router"
	"gotest/internal/core/model/request"
	"gotest/internal/core/port/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	gin         *gin.Engine
	authService service.AuthService
}

func NewAuthController(gin *gin.Engine, authService service.AuthService) AuthController {
	return AuthController{
		gin:         gin,
		authService: authService,
	}
}

func (a AuthController) InitRouter() {
	var token service.TokenService
	api := a.gin.Group("api/v1").Use(authMiddleware(token))
	{
		router.PostWithMiddleware(api, "/login", a.login)
	}
}

func (a AuthController) login(c *gin.Context) {
	var req request.GetPersonsByFirstName

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &invalidRequestResponse)
	}

	resp := a.authService.Login(&req)

	c.JSON(http.StatusOK, resp)
}

func (a AuthController) parseRequest(ctx *gin.Context) (*request.GetPersonsByFirstName, error) {
	var req request.GetPersonsByFirstName

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
