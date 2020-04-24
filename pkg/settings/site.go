package settings

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

// Config app configuration
type Config struct {
	Server struct {
		PrefixURL       string
		RunMode         string
		HTTPPort        int
		RuntimeRootPath string
		ReadTimeout     time.Duration
		WriteTimeout    time.Duration
		TimeFormat      string
	}
	Database struct {
		Type        string
		Connections map[string]string
	}
	Redis map[string]struct {
		Addr     string
		Password string
		DB       int
	}
	Logger struct {
		LogSavePath string
		LogSaveName string
		LogFileExt  string
	}
	Auth struct {
		JwtSecret  string
		ExpireTime time.Duration
	}
	Pagger struct {
		PageSize int
	}
	Media map[string]struct {
		Access string
		Secret string
		Bucket string
	}
}

// AppConfig app config
var AppConfig = &Config{}

// Setup initialize the configuration instance
func Setup() {
	read, readErr := ioutil.ReadFile("config/site.yml")
	if readErr != nil {
		log.Fatalf("config setup fail: %v", readErr)
	}
	err := yaml.Unmarshal(read, AppConfig)
	if err != nil {
		log.Fatalf("config setup fail: %v", err)
	}
	AppConfig.Server.ReadTimeout = AppConfig.Server.ReadTimeout * time.Second
	AppConfig.Server.WriteTimeout = AppConfig.Server.WriteTimeout * time.Second
	AppConfig.Auth.ExpireTime = AppConfig.Auth.ExpireTime * time.Hour
}

// GetDomain  get domain
func GetDomain(id string) string {
	/*if strings.TrimSpace(id) == "" ||
		strings.Contains(id, "localhost") ||
		strings.Contains(id, "127.0.0.1") {
		id = "default"
	}
	return id*/
	return "default"
}
