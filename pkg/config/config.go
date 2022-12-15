package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/oswaldom-code/api-template-gin/pkg/log"

	"github.com/spf13/viper"
)

const (
	CONFIG_FILE = "api"
)

var environment string

type EnvironmentConfig struct {
	Environment string
}

type DBConfig struct {
	User               string
	Password           string
	Host               string
	Port               int
	Database           string
	MaxOpenConnections int
	SSLMode            string
	LogMode            string
	Engine             string
}

type ServerConfig struct {
	Host              string
	Port              string
	Scheme            string
	Mode              string
	PathToSSLKeyFile  string
	PathToSSLCertFile string
	Static            string
}

func SetEnvironment(env string) {
	environment = env
}

// Validate ServerConfig value validation
func (s *ServerConfig) Validate() error {
	if s.Host == "" || s.Port == "0" || s.Scheme == "" || s.Mode == "" {
		return fmt.Errorf(`ServerConfig is invalid: \n
		env: %s
		host: %s
		port: %s
		scheme: %s
		mode: %s`,
			environment, s.Host, s.Port, s.Scheme, s.Mode)
	}
	return nil
}

// AsUri returns the host:port
func (s ServerConfig) AsUri() string {
	return s.Host + ":" + s.Port
}

type LoggingConfig struct {
	Level        string
	ErrorLogFile string
}

type AuthenticateKeyConfig struct {
	Secret string
}

// GetProjectPath returns the current project path
func GetProjectPath() string {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Warn("Warning, cannot get current path")
		return ""
	}
	// Traverse back from current directory until service base dir is reach and add to config path
	for !strings.HasSuffix(dir, "") && dir != "/" {
		dir, err = filepath.Abs(dir + "/..")
		if err != nil {
			break
		}
	}
	return dir
}

func LoadConfigurationFile() {
	viper.SetConfigName(CONFIG_FILE)
	viper.AddConfigPath(GetProjectPath() + "/config")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_", ".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config error:", err.Error())
	}
	log.SetLogLevel(GetLogConfig().Level)
}

func GetDBConfig() DBConfig {
	// Get current environment
	config := DBConfig{
		User:               viper.GetString(environment + ".db.user"),
		Password:           viper.GetString(environment + ".db.password"),
		Host:               viper.GetString(environment + ".db.host"),
		Port:               viper.GetInt(environment + ".db.port"),
		Database:           viper.GetString(environment + ".db.database"),
		MaxOpenConnections: viper.GetInt(environment + ".db.max_connections"),
		SSLMode:            viper.GetString(environment + ".db.ssl_mode"),
		LogMode:            viper.GetString(environment + ".db.log_mode"),
		Engine:             viper.GetString(environment + ".db.engine"),
	}
	log.DebugWithFields("DBConfig", log.Fields{"config": config})
	return config
}

func GetServerConfig() ServerConfig {
	log.DebugWithFields("GetServerConfig", log.Fields{"environment": environment})
	config := ServerConfig{
		Host:              viper.GetString(environment + ".server.host"),
		Port:              viper.GetString(environment + ".server.port"),
		Scheme:            viper.GetString(environment + ".server.scheme"),
		Mode:              viper.GetString(environment + ".server.mode"),
		PathToSSLKeyFile:  viper.GetString(environment + ".server.ssl.key"),
		PathToSSLCertFile: viper.GetString(environment + ".server.ssl.cert"),
		Static:            viper.GetString(environment + ".server.static"),
	}
	log.DebugWithFields("ServerConfig", log.Fields{"config": config})
	return config
}

func GetPathFromWhereStaticFilesWillBeServed() string {
	return viper.GetString(environment + ".server.static")
}

func GetLogConfig() LoggingConfig {
	return LoggingConfig{
		Level:        viper.GetString(environment + ".log.level"),
		ErrorLogFile: viper.GetString(environment + ".log.errorLogFile"),
	}
}

func GetEnvironmentConfig() EnvironmentConfig {
	return EnvironmentConfig{
		Environment: viper.GetString("environment"),
	}
}

func GetAuthenticationKey() AuthenticateKeyConfig {
	return AuthenticateKeyConfig{
		Secret: viper.GetString(environment + ".auth.secret"),
	}
}
