package config

import (
	"fmt"
	"github.com/caarlos0/env"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ConfigurationEnvironment struct {
	DatabaseEnvironment
	ServerEnvironment
	JWTEnvironment
	LoggerEnvironment
	ApplicationEnvironment
	ProtoBuffEnvironment
}

type ProtoBuffEnvironment struct {
	PortServer string `env:"port_server"`
	Protocol   string `env:"protocol"`
}

type DatabaseEnvironment struct {
	MaxIdle    int    `env:"max_idle"`
	MaxPool    int    `env:"max_pool"`
	Driver     string `env:"driver"`
	PSN        string `env:"psn"`
	LogQueries bool   `env:"log_queries"`
	Timeout    int    `env:"timeout_seconds"`
}

type ServerEnvironment struct {
	Port         string `env:"port"`
	Debug        bool   `env:"debug"`
	ReadTimeout  int    `env:"read_timeout_seconds"`
	WriteTimeout int    `env:"write_timeout_seconds"`
}

type JWTEnvironment struct {
	Secret           string `env:"secret"`
	Duration         int    `env:"duration_minutes"`
	RefreshDuration  int    `env:"refresh_duration_minutes"`
	MaxRefresh       int    `env:"max_refresh_minutes"`
	SigningAlgorithm string `env:"signing_algorithm"`
}

type LoggerEnvironment struct {
	Path      string `env:"path_name"`
	LogName   string `env:"log_name"`
	Prod      bool   `env:"prod"`
	WithTrace bool   `env:"with_trace"`
}

type ApplicationEnvironment struct {
	MinPasswordStr int    `env:"min_password_strength"`
	SwaggerUIPath  string `env:"swagger_ui_path"`
}

var Env = ConfigurationEnvironment{}

func LoadEnv() {
	if err := env.Parse(&Env); err != nil {
		panic(err.Error())
	}
}

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server *Server      `yaml:"server,omitempty"`
	DB     *Database    `yaml:"database,omitempty"`
	JWT    *JWT         `yaml:"jwt,omitempty"`
	App    *Application `yaml:"application,omitempty"`
	Logger *Logger      `yaml:"logger"`
}

// Database holds data necessery for database configuration
type Database struct {
	MaxIdle    int    `yaml:"max_idle,omitempty"`
	MaxPool    int    `yaml:"max_pool,omitempty"`
	Driver     string `yaml:"driver,omitempty"`
	PSN        string `yaml:"psn,omitempty"`
	LogQueries bool   `yaml:"log_queries,omitempty"`
	Timeout    int    `yaml:"timeout_seconds,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

// JWT holds data necessery for JWT configuration
type JWT struct {
	Secret           string `yaml:"secret,omitempty"`
	Duration         int    `yaml:"duration_minutes,omitempty"`
	RefreshDuration  int    `yaml:"refresh_duration_minutes,omitempty"`
	MaxRefresh       int    `yaml:"max_refresh_minutes,omitempty"`
	SigningAlgorithm string `yaml:"signing_algorithm,omitempty"`
}

type Logger struct {
	Path      string `yaml:"path"`
	LogName   string `yaml:"log_name,omitempty"`
	Prod      bool   `yaml:"prod,omitempty"`
	WithTrace bool   `yaml:"with_trace,omitempty"`
}

// Application holds application configuration details
type Application struct {
	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
}
