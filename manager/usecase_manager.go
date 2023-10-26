package manager

import (
	"go-payment-simulation/usecase"

	"github.com/gin-gonic/gin"
)

type UseCaseManager interface {
	UserUsecase() usecase.UserUsecase
}

type useCaseManager struct {
	repoManager RepoManager
	ctx         *gin.Context
}

// UserUsecase implements UseCaseManager.
func (ucm *useCaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUseCase(ucm.repoManager.UserRepo())
}

func NewUsecaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
		ctx:         &gin.Context{},
	}
}
