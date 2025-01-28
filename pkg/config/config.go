package config

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LoadAllConfigs(envFilePath string) {
    if err := godotenv.Load(envFilePath); err != nil {
        log.Fatalf("can't load .env file. error: %v", err)
    }

    LoadApp()
}

func FiberConfig() fiber.Config {
    return fiber.Config {
        ReadTimeout: time.Second * time.Duration(AppCfg().ReadTimeout),
    }
}
