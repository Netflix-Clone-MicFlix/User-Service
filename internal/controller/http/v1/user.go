package v1

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Netflix-Clone-MicFlix/User-Service/internal"
	"github.com/Netflix-Clone-MicFlix/User-Service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-Service/pkg/logger"
)

type UserRoutes struct {
	t internal.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, t internal.User, l logger.Interface) {
	r := &UserRoutes{t, l}

	h := handler.Group("/user")
	{
		h.GET("", r.GetAll)
		h.GET("/:user_id", r.GetById)
	}
}

type userResponse struct {
	Users []entity.User `json:"users"`
}

// @Summary     Show users
// @Description Show all users
// @ID          user
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Success     200 {object} userResponse
// @Failure     500 {object} response
// @Router      /user [get]
func (r *UserRoutes) GetAll(c *gin.Context) {
	users, err := r.t.GetAll(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - GetAll")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, userResponse{users})
}

type UserRequest struct {
	User entity.User `json:"users"`
}

// @Summary     User
// @Description User a text
// @ID          do-translate
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body doUserRequest true "Set up user"
// @Success     200 {object} entity.User
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /user/do-translate [post]
func (r *UserRoutes) GetById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		r.l.Error(err, "http - v1 - survey")
		errorResponse(c, http.StatusBadRequest, "SurveyId not an integer")
	}

	user, err := r.t.GetById(c.Request.Context(), userId)
	if err != nil {
		r.l.Error(err, "http - v1 - doUser")
		errorResponse(c, http.StatusInternalServerError, "user service problems")

		return
	}

	c.JSON(http.StatusOK, user)
}
