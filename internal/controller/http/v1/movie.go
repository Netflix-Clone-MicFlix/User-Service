package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Netflix-Clone-MicFlix/User-service/internal"
	"github.com/Netflix-Clone-MicFlix/User-service/internal/entity"
	"github.com/Netflix-Clone-MicFlix/User-service/pkg/logger"
)

type UserRoutes struct {
	t          internal.User
	l          logger.Interface
	corsConfig gin.HandlerFunc
}

func newUserRoutes(handler *gin.RouterGroup, t internal.User, l logger.Interface, corsConfig gin.HandlerFunc) {
	r := &UserRoutes{t, l, corsConfig}

	user := handler.Group("/user")
	{
		user.Use(corsConfig)
		user.GET("", r.GetAll)
		user.POST("", r.Create)
		user.GET("/:user_id", r.GetById)
		genres := user.Group("/profile")
		{
			genres.Use(corsConfig)
			genres.GET("/:user_id", r.GetAllProfilesById)
		}
	}

}

type userCollectionResponse struct {
	Users []entity.User `json:"users"`
}
type CreateUserRequest struct {
	Id         string   `json:"id"    example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	KeycloakId string   `json:"keycloak_id"    example:"6be244a7-25ac-34ce-31e3-04157d3d42e3"`
	ProfileIds []string `json:"profile_ids"    example:"[6be244a7-25ac-34ce-31e3-04157d3d42e3,6be244a7-25ac-34ce-31e3-04157d3d42e3]"`
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

	c.JSON(http.StatusOK, userCollectionResponse{users})
}

// @Summary     Show user with id
// @Description Show users with id
// @ID          user
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Success     200 {object} userResponse
// @Failure     500 {object} response
// @Router      /user [get]
func (r *UserRoutes) GetById(c *gin.Context) {
	userId := c.Param("user_id")

	user, err := r.t.GetById(c.Request.Context(), userId)
	if err != nil {
		r.l.Error(err, "http - v1 - doUser")
		errorResponse(c, http.StatusInternalServerError, "user service problems")

		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary     creates user
// @Description creates user with discription and title
// @ID          user
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Success     200 {object} CreateUserRequest
// @Failure     500 {object} response
// @Router      /user [post]
func (r *UserRoutes) Create(c *gin.Context) {
	var request CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - Register user")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.t.Create(c.Request.Context(), request.KeycloakId)
	if err != nil {
		r.l.Error(err, "http - v1 - doUser")
		errorResponse(c, http.StatusInternalServerError, "user service problems")

		return
	}

	c.JSON(http.StatusOK, nil)
}

// @Summary     Show profile with id
// @Description Shows profiles with id
// @ID          user
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Success     200 {object} profile
// @Failure     500 {object} response
// @Router      /user/profile [get]
func (r *UserRoutes) GetAllProfilesById(c *gin.Context) {
	userId := c.Param("user_id")

	profiles, err := r.t.GetAllProfilesById(c.Request.Context(), userId)
	if err != nil {
		r.l.Error(err, "http - v1 - doUser")
		errorResponse(c, http.StatusInternalServerError, "user service problems")

		return
	}

	c.JSON(http.StatusOK, profiles)
}
