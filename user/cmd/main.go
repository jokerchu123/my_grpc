package main

import (
	"user/config"
	"user/internal/repository"
)

func main() {
	config.InitConfig()
	repository.InitDB()
}
