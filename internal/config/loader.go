package config

import(
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get() * Config{
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error when loading file configuration", err.Error())
	}

	return &Config{
		Server:Server{
			Host:os.Getenv("SERVER_HOST"),
			Port:os.Getenv("SERVER_PORT"),
		},
	}
}