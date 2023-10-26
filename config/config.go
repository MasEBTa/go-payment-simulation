package config

import (
	"fmt"
	"go-payment-simulation/utils/common"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

// db config
type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

// json data file config
type JSData struct {
	JSFile string
}

// api config
type ApiConfig struct {
	ApiHost string
	ApiPort string
}

// token config
type TokenConfig struct {
	ApplicationName  string
	JwtSignatureKey  []byte
	JwtSigningMethod *jwt.SigningMethodHMAC
	ExpirationToken  int
}

// file config (logger)
type FileConfig struct {
	FilePath string
}

type Config struct {
	DbConfig
	JSData
	ApiConfig
	TokenConfig
	FileConfig
}

func (c *Config) ReadConfig() error {
	// env
	err := common.LoadEnv()
	if err != nil {
		return err
	}

	// environtment Variable
	// c.DbConfig = DbConfig{
		// Host:     os.Getenv("DB_HOST"),
		// Port:     os.Getenv("DB_PORT"),
		// Name:     os.Getenv("DB_NAME"),
		// User:     os.Getenv("DB_USER"),
		// Password: os.Getenv("DB_PASSWORD"),
		// Driver:   os.Getenv("DB_DRIVER"),
	// }

	// if c.DbConfig.Host == "" ||
	// 	c.DbConfig.Port == "" ||
	// 	c.DbConfig.Name == "" ||
	// 	c.DbConfig.User == "" ||
	// 	c.DbConfig.Password == "" ||
	// 	c.DbConfig.Driver == "" {
	// 	return fmt.Errorf("missing requirenment variable")
	// }

	// file data json
	c.JSData = JSData{
		JSFile: os.Getenv("JS_FILE"),
	}
	if c.JSData.JSFile == "" {
		return fmt.Errorf("missing requirenment variable")
	}

	// api config env
	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"),
		ApiPort: os.Getenv("API_PORT"),
	}
	if c.ApiConfig.ApiHost == "" ||
		c.ApiConfig.ApiPort == ""  {
		return fmt.Errorf("missing requirenment variable")
	}

	// token
	expiration, err := strconv.Atoi(os.Getenv("APP_EXPIRATION_TOKEN"))
	if err != nil {
		return err
	}
	c.TokenConfig = TokenConfig{
		ApplicationName:  os.Getenv("APP_TOKEN_NAME"),
		JwtSignatureKey:  []byte(os.Getenv("APP_TOKEN_KEY")),
		JwtSigningMethod: jwt.SigningMethodHS256,
		ExpirationToken:  expiration,
	}

	// logged file path
	c.FileConfig = FileConfig{
		FilePath: os.Getenv("FILE_PATH"),
	}

	return nil
}

// constructor
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
