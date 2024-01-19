package http

import (
	"gotest/internal/core/common/router"
	"gotest/internal/core/entity/error_code"
	"gotest/internal/core/model/request"
	"gotest/internal/core/model/response"
	"gotest/internal/core/port/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	invalidRequestResponse = response.HttpResponse{
		ErrorCode:    error_code.InvalidRequest,
		ErrorMessage: error_code.InvalidRequestErrMsg,
		Status:       false,
	}
)

type PersonController struct {
	gin           *gin.Engine
	personService service.PersonService
}

func NewUserController(gin *gin.Engine, personService service.PersonService) PersonController {
	return PersonController{
		gin:           gin,
		personService: personService,
	}
}

func (p PersonController) InitRouter() {
	api := p.gin.Group("/api/v1")
	router.Post(api, "/signup", p.signUp)
	router.Get(api, "/persons", p.listPersons)
	router.Delete(api, "/persons/:PersonsID", p.DeletePerson)
}

func (p PersonController) signUp(c *gin.Context) {
	req, err := p.parseRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &invalidRequestResponse)
	}

	resp := p.personService.AddPerson(req)
	c.JSON(http.StatusOK, resp)
}

func (p PersonController) DeletePerson(c *gin.Context) {
	req, err := p.parseDeleteRequest(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &invalidRequestResponse)
	}

	resp := p.personService.DeletePerson(req)

	c.JSON(http.StatusOK, resp)
}

func (p PersonController) listPersons(c *gin.Context) {
	resp := p.personService.GetListPersons()

	c.JSON(http.StatusOK, resp)
}

func (p PersonController) parseRequest(ctx *gin.Context) (*request.AddPersonRequest, error) {
	var req request.AddPersonRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func (p PersonController) parseDeleteRequest(ctx *gin.Context) (*request.DeletePersonRequest, error) {
	var req request.DeletePersonRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
