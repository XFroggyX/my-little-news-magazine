package user

import (
	logging "github.com/XFroggyX/go-logger"
	"github.com/XFroggyX/my-little-news-magazine/internal/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *gin.Engine) {
	userAPI := router.Group("/api")
	{
		userAPI.GET("/user", h.GetList)
		userAPI.GET("/user/:uuid", h.GetUserByUUID)
		userAPI.POST("/user/:uuid", h.CreateUser)
		userAPI.PUT("/user/:uuid", h.UpdateUser)
		userAPI.PATCH("/user/:uuid", h.PartiallyUpdateUser)
		userAPI.DELETE("/user/:uuid", h.DeleteUser)
	}
}

func (h *handler) GetList(c *gin.Context) {
	c.String(http.StatusOK, "this is list of users")
}

func (h *handler) GetUserByUUID(c *gin.Context) {
	c.String(http.StatusOK, "this user by uuid")
}
func (h *handler) CreateUser(c *gin.Context) {
	c.String(http.StatusNoContent, "this is create user")
}
func (h *handler) UpdateUser(c *gin.Context) {
	c.String(http.StatusNoContent, "this is update user")
}
func (h *handler) PartiallyUpdateUser(c *gin.Context) {
	c.String(http.StatusNoContent, "this is user partially")
}
func (h *handler) DeleteUser(c *gin.Context) {
	c.String(http.StatusNoContent, "this is delete user")
}
