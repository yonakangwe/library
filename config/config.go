package config

import (
	"library/package/log"
	"library/package/util"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	WebServer   WebServerConfig
	Database    DatabaseConfig
	Subsystems  []*Subsystem
	Secret      SecretKey
	Mail        MailConfig
	PrivateKeys []Key `yaml:"privateKeys"`
	PublicKeys  []Key `yaml:"publicKeys"`
}

type Key struct {
	SystemName string `yaml:"systemName"`
	KeyPath    string `yaml:"keyPath"`
}

type MailConfig struct {
	UserName   string
	Secret     string
	ServerName string
}

type WebServerConfig struct {
	LocalHost  string
	PublicHost string
	Port       int
	Url        string
	Name       string
}
type Subsystem struct {
	Name   string
	Secret string
}
type DatabaseConfig struct {
	Name     string
	User     string
	Password string
	Port     int
}

type SecretKey struct {
	Secret string
}

func New() (*Config, error) {
	configFile := "config.yml"
	confPath, err := os.Getwd()
	if util.IsError(err) {
		log.Error("error getting a working directory:%v", err)
		return nil, err
	}
	configPath := fmt.Sprintf("%s/%s", confPath, configFile)

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error("error reading config file, %s", err)
		return nil, err
	}

	cfg := Config{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Error("Unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) GetSystemPrivateKey(systemName string) ([]byte, error) {
	for _, privKey := range c.PrivateKeys {
		if privKey.SystemName == systemName && privKey.KeyPath != "" {
			key, err := os.ReadFile(privKey.KeyPath)
			if util.IsError(err) {
				log.Errorf("error reading private key: %s", err)
				return nil, err
			}
			return key, nil
		}
	}
	log.Errorf("could not find the private key for: %s", systemName)
	return nil, errors.New("401 Unauthorized")
}

func (c *Config) GetSystemPublicKey(systemName string) ([]byte, error) {
	for _, pubKey := range c.PublicKeys {
		if pubKey.SystemName == systemName && pubKey.KeyPath != "" {
			key, err := os.ReadFile(pubKey.KeyPath)
			if util.IsError(err) {
				log.Errorf("error reading public key: %s", err)
				return nil, err
			}
			return key, nil
		}
	}
	log.Errorf("could not find the public key for: %s", systemName)
	return nil, errors.New("401 Unauthorized")
}

func (c *Config) GetSecret() string {
	return c.Secret.Secret
}

func (c *Config) MailConfig() MailConfig {
	return c.Mail
}

func GetDatabaseConnection() string {
	cfg, err := New()
	if err != nil {
		log.Errorf("error loading configuration file: %v", err)
		return ""
	}
	return cfg.GetDatabaseConnection()
}
func (c *Config) GetDatabaseConnection() string {
	conn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d timezone=%s options= ", c.WebServer.LocalHost, c.Database.Name, c.Database.User, c.Database.Password, c.Database.Port, "Africa/Dar_es_Salaam")
	return conn
}

func LoggerPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Errorf("error getting file path %s\n", err)
	}
	return path + "/.storage/.logs"
}

func TemplatePath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Errorf("Error getting file path %s\n", err)
	}
	path = fmt.Sprintf("%s/docs/templates/", path)
	return path
}

// LogoPath returns logo path
func LogoPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/webserver/public/images/tzlogo.png", path)
	return path, nil
}

// ReportDir returns .storage/reports path
func ReportDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/.storage/reports/", path)
	return path, nil
}

func DownloadDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/.storage/reports/downloads/", path)
	return path, nil
}

func UploadsDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/.storage/reports/uploads/", path)
	return path, nil
}
