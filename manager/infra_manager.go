package manager

import "go-payment-simulation/config"

type InfraManager interface {
	Conn() *config.Config
}

type infraManager struct {
	cfg *config.Config
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}
	return conn, nil
}

func (im *infraManager) Conn() *config.Config {
	return im.cfg
}