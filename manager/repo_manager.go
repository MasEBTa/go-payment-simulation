package manager

import (
	"fmt"
	"go-payment-simulation/repository"
)

type RepoManager interface {
	BaseRepo() repository.BasesRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infraManager InfraManager
}

// BaseRepo implements RepoManager.
func (rm *repoManager) BaseRepo() repository.BasesRepository {
	basesrepo, err := repository.NewBasesRepository(rm.infraManager.Conn().JSFile)
	if err!=nil {
		fmt.Println("Error reading configuration:", err)
	}
	return basesrepo
}

// UserRepo implements RepoManager.
func (rm *repoManager) UserRepo() repository.UserRepository {
	base, err := rm.BaseRepo().FindByName("user")
	if err!=nil {
		fmt.Println("Error finding data:", err)
	}
	userrepo, err :=  repository.NewUserRepository(base.File)
	if err!=nil {
		fmt.Println("Error reading configuration:", err)
	}
	return userrepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
