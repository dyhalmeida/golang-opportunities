package main

import (
	"github.com/dyhalmeida/golang-opportunities/internal/config"
	"github.com/dyhalmeida/golang-opportunities/internal/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
