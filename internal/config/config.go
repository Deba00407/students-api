package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost" env-required:"true"`
	Port int    `yaml:"port" env:"PORT" env-default:"5001" env-required:"true"`
	ServerPath string `yaml:"serverPath" env:"SERVER_PATH" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"production" env-required:"true"`
	StoragePath string `yaml:"storage" env-required:"true"`
	HTTPServer  `yaml:"http_server" env-required:"true"`
}

func MustLoadConfig() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	// In case if user provides config file path in cli flags
	if configPath == ""{
		flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == ""{
			log.Fatal("Config path was not set")
		}
	}

	// check if there is a file in the configPath
	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("Config file was not found: %s", configPath)
	}

	// Read the config file
	var config Config

	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil{
		log.Fatalf("Could not read config file: %s", err.Error())
	}

	fmt.Println("Config loaded successfully")
	return &config
}