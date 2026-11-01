package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"xor-operation/internal/config"
	"xor-operation/internal/service"
	"xor-operation/internal/api"
)

func main(){
	conf := config.Get()
	app := fiber.New()

	app.Use(cors.New())

	encryptService := service.NewEncrypt()

	
	v1 := fiber.New()
	api.NewEncrypt(v1,encryptService)

	app.Mount("/v1", v1)

	log.Fatal(app.Listen(conf.Server.Host + ":" + conf.Server.Port))
}