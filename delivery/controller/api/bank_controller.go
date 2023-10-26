package api

import (
	"go-payment-simulation/model/dto"
	"go-payment-simulation/usecase"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	userUC usecase.UserUsecase
	rg     *gin.RouterGroup
}

func (b *BankController) listHandler(c *gin.Context) {
	bank, err := b.userUC.FindByRole("bank")
	if err!=nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": err.Error(),
		})
	}
	response := dto.ResponseData{
		Message: "successfully logged-in",
		Data: bank,
	}

	c.JSON(200, response)
}

func (b *BankController) Route() {
	b.rg.GET("/bank/list", b.listHandler)
}

func NewBankController(userUC usecase.UserUsecase, rg *gin.RouterGroup) *BankController {
	return &BankController{
		userUC: userUC,
		rg:     rg,
	}
}
