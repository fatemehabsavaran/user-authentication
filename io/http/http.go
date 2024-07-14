package http

import (
	"fmt"
	"github.com/fatemehabsavaran/user-authentication.git/service"
	"github.com/gin-gonic/gin"
)

type GinConfig struct {
	Port int `json:"port"`
}
type UserController struct {
	engine      *gin.Engine
	config      GinConfig
	userService *service.UserService
}

func NewUserController(config GinConfig, user service.UserService) *UserController {
	engine := gin.Default()
	return &UserController{
		engine:      engine,
		config:      config,
		userService: &user,
	}
}
func (s *UserController) Start() error {
	return s.engine.Run(fmt.Sprintf("127.0.0.1:%v", s.config.Port))
}
func (s *UserController) RegisterRoutes() {
	s.engine.POST("/auth/signup", s.Signup)
}
func (s *UserController) Signup(c *gin.Context) {
	signup := SignupRequest{}
	if err := c.BindJSON(&signup); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := s.userService.SignUp(signup.Email, signup.Name, signup.Pass)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := SignupResponse{
		Name:  user.Name,
		ID:    user.ID,
		Email: user.Email,
		Token: user.Tokens[0].Token,
	}

	c.JSON(200, response)
}
