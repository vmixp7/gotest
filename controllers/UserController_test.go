package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest/core"
	"gotest/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// mockDB is a mock for core.DB
type mockDB struct {
	shouldFind bool
}

func (m *mockDB) Where(query interface{}, args ...interface{}) *mockDB {
	if len(args) > 0 && args[0] == "founduser" {
		m.shouldFind = true
	} else {
		m.shouldFind = false
	}
	return m
}

func (m *mockDB) First(dest interface{}, conds ...interface{}) error {
	if m.shouldFind {
		user := dest.(*models.User)
		user.ID = 1
		user.Username = "founduser"
		user.Email = "found@example.com"
		return nil
	}
	return errors.New("record not found")
}

// patch core.DB for testing
func setMockDB(m *mockDB) func() {
	orig := core.DB
	core.DB = m
	return func() { core.DB = orig }
}

func setupRouterForGetUserByUsername() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/users/name", GetUserByUsername)
	return r
}

func TestGetUserByUsername_Found(t *testing.T) {
	mock := &mockDB{}
	restore := setMockDB(mock)
	defer restore()

	router := setupRouterForGetUserByUsername()
	req, _ := http.NewRequest("GET", "/users/name?username=founduser", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"username":"founduser"`)
	assert.Contains(t, w.Body.String(), `"email":"found@example.com"`)
}

func TestGetUserByUsername_NotFound(t *testing.T) {
	mock := &mockDB{}
	restore := setMockDB(mock)
	defer restore()

	router := setupRouterForGetUserByUsername()
	req, _ := http.NewRequest("GET", "/users/name?username=notfound", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "User not found")
}

func TestGetUserByUsername_NoUsernameParam(t *testing.T) {
	mock := &mockDB{}
	restore := setMockDB(mock)
	defer restore()

	router := setupRouterForGetUserByUsername()
	req, _ := http.NewRequest("GET", "/users/name", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "User not found")
}
