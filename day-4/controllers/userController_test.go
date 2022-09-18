package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	controllers UserControllers
)

var (
	userJSONValid = `{
		"name":"john",
		"email":"yy@mail.com",
		"password": "ggpass"
	}`
)

func TestCreateUsers_success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(userJSONValid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUserControllers(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		body := rec.Body.String()
		assert.Contains(t, body, "name")
		assert.Contains(t, body, "email")
		assert.Contains(t, body, "password")
	}
}

func TestGetAllUsersSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(controllers.GetUsersControllers(c)) {
		asserts.Equal(200, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "success to fetch data from server")
		asserts.Contains(body, "success")
	}
}
