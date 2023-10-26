package api

import (
	"go-payment-simulation/model/dto"
	"go-payment-simulation/usecase"

	"github.com/gin-gonic/gin"
)

type MarchandController struct {
	userUC usecase.UserUsecase
	rg     *gin.RouterGroup
}

func (mc *MarchandController) listHandler(c *gin.Context) {
	march, err := mc.userUC.FindByRole("marchand")
	if err!=nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
	}
	response := dto.ResponseData{
		Message: "successfully logged-in",
		Data: march,
	}

	c.JSON(200, response)
}

func (mc *MarchandController) Route() {
	mc.rg.GET("/marchand/list", mc.listHandler)
}

func NewMarchandController(userUC usecase.UserUsecase, rg *gin.RouterGroup) *MarchandController {
	return &MarchandController{
		userUC: userUC,
		rg:     rg,
	}
}
