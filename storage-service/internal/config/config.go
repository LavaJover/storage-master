package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type StorageConfig struct {
	Env string 	`yaml:"env" env-required:"true"`
	GRPCServer 	`yaml:"grpc_server" env-required:"true"`
	Dsn string 	`yaml:"dsn" env-required:"true"`
}

type GRPCServer struct {
	Host string `yaml:"host" env-default:"localhost" env-required:"true"`
	Port string `yaml:"port" env-default:"50051" env-required:"true"`
}

func MustLoad() *StorageConfig {
	var cfg StorageConfig

	configPath := os.Getenv("STORAGE_CONFIG_PATH")

	if configPath == ""{
		panic("STORAGE_CONFIG_PATH required")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatal("storage-service config file was not found")
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatal("failed to read config")
	}

	return &cfg
}
