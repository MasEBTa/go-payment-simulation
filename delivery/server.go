package delivery

import (
	"fmt"
	"go-payment-simulation/config"
	"go-payment-simulation/delivery/controller/api"
	"go-payment-simulation/delivery/middleware"
	"go-payment-simulation/manager"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
	log       *logrus.Logger
}

func (s *Server) Run() {
	s.initMiddlewares()
	s.initControllers()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initMiddlewares() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
}

func (s *Server) initControllers() {
	rg := s.engine.Group("/api/v1")
	api.NewAuthController(s.ucManager.UserUsecase(), rg).Route()
	api.NewBankController(s.ucManager.UserUsecase(), rg).Route()
	api.NewMarchandController(s.ucManager.UserUsecase(), rg).Route()
}

func NewServer() *Server {
	// config
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	// infrastruktur manager
	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Println(err)
	}
	// repo manager
	repoManager := manager.NewRepoManager(infraManager)
	// usecase manager
	useCaseManager := manager.NewUsecaseManager(repoManager)
	// config
	engine := gin.Default()
	// host
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	// logger
	log := logrus.New()

	return &Server{
		ucManager: useCaseManager,
		engine:    engine,
		host:      host,
		log:       log,
	}
}