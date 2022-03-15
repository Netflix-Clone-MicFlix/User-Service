package v1

import (
	"net/http"

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
		h.POST("/register/", r.Register)
		h.POST("/login/", r.Login)
	}
}

type userCollectionResponse struct {
	Users []entity.User `json:"users"`
}
type UserRequest struct {
	User entity.User `json:"users"`
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

// @Summary     User Register
// @Description Registers User
// @ID          Registers
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body UserRequest true "Set up user"
// @Success     200 {object} entity.User
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /user/{UserRequest}[post]
func (r *UserRoutes) Register(c *gin.Context) {
	var request UserRequest

	if err := c.ShouldBindJSON(&request.User); err != nil {
		r.l.Error(err, "http - v1 - Register user")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.t.Register(c.Request.Context(), entity.User{
		Username: request.User.Username,
		Password: request.User.Password,
		Email:    request.User.Email,
	})

	if err != nil {
		r.l.Error(err, "http - v1 - Register")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, nil)
}

// @Summary     User Login
// @Description Login User
// @ID          Login
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body UserRequest true "Set up user"
// @Success     200 {object} entity.User
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /user/{UserRequest}[post]
func (r *UserRoutes) Login(c *gin.Context) {
	var request UserRequest

	if err := c.ShouldBindJSON(&request.User); err != nil {
		r.l.Error(err, "http - v1 - login reqeust")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.t.Login(c.Request.Context(), entity.User{
		Username: request.User.Username,
		Password: request.User.Password,
		Email:    request.User.Email,
	})

	if err != nil {
		r.l.Error(err, "http - v1 - login")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	//tempory
	succes := "Login succesfull!"

	c.JSON(http.StatusOK, succes)
}
