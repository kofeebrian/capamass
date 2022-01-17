package config

type ServerConfig struct {
	Environment string
	Port        uint
}

type ServiceConfig struct {
	Port uint
}

type ClientConfig struct {
	ServerPort uint
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		Port: 50051,
	}
}
