package main

import (
	"github.com/kofeebrian/capamass/config"
	"github.com/kofeebrian/capamass/service"
)

func main() {
	config := config.NewServiceConfig()
	config.Port = 3000

	service.Init(config)
}
