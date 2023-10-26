package api

import (
	"go-payment-simulation/model"
	"go-payment-simulation/model/dto"
	"go-payment-simulation/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userUC usecase.UserUsecase
	rg     *gin.RouterGroup
}

func (a *AuthController) registerHandler(c *gin.Context) {
	var auth model.UserCredential
	
	// ambil data
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Faild getting data: "+err.Error(),
		})
		return
	}
	// pengecekan input
	if auth.Username == "" {
		c.AbortWithStatusJSON(422, gin.H{
			"message": "Username is required.",
		})
		return
	}
	if auth.Email == "" {
		c.AbortWithStatusJSON(422, gin.H{
			"message": "Email is required.",
		})
		return
	}
	if auth.Password == "" {
		c.AbortWithStatusJSON(422, gin.H{
			"message": "Password is required.",
		})
		return
	}
	if auth.UserRole == "" {
    c.AbortWithStatusJSON(422, gin.H{
        "message": "Role is required.",
    })
    return
	}
	if auth.UserRole != "user" && auth.UserRole != "marchand" && auth.UserRole != "bank" {
    c.AbortWithStatusJSON(422, gin.H{
        "message": "Role is invalid.",
    })
    return
	}

	// kirim ke usecase
	user, err := a.userUC.NewUser(auth)
	if err != nil {
		// cek apakah karena email | username sudah digunakan
		if err.Error() == "email is used" || err.Error() == "username is used" {
			c.AbortWithStatusJSON(422, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := dto.ResponseData{
		Message: "successfully register",
		Data: user,
	}

	c.JSON(200, gin.H{
		"message": response.Message,
	})
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var auth model.UserCredential
	// ambil data
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Faild getting data: "+err.Error(),
		})
		return
	}
	// pengecekan input
	if auth.Email == "" && auth.Username == "" {
		c.AbortWithStatusJSON(422, gin.H{
			"message": "Email or Username is required",
		})
		return
	}
	if auth.Password == "" {
		c.AbortWithStatusJSON(422, gin.H{
			"message": "Password is required",
		})
		return
	}
	// kirim ke usecase
	user, err := a.userUC.Login(auth)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Infalid Credential",
		})
		return
	}

	response := dto.ResponseData{
		Message: "successfully logged-in",
		Data: user,
	}

	c.JSON(200, response)
}

func (a *AuthController) Route() {
	a.rg.POST("/auth/register", a.registerHandler)
	a.rg.GET("/auth/login", a.loginHandler)
}

func NewAuthController(userUC usecase.UserUsecase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{
		userUC: userUC,
		rg:     rg,
	}
}
