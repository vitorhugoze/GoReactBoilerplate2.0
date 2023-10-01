package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"server/utils"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func App() {

	if os.Getenv("APP_ENV") != "PROD" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}

	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	interruptionHandler()
	utils.CreateRedisClient()
	utils.CreateMongoClient()

	initializeFiber()
}

func initializeFiber() {

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Server",
		AppName:       "Go server",
	})

	app.Use(logger.New(logger.Config{
		Format: "[${time}] [${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Group("/").Route("/", MainRouter)

	log.Fatal(app.Listen(fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))))
}

func interruptionHandler() {
	c := make(chan os.Signal, 3)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c

		if err := utils.CloseRedisClient(); err != nil {
			log.Println(err)
		}

		if err := utils.CloseMongoClient(); err != nil {
			log.Println(err)
		}

		log.Fatal("system interruption!")
	}()
}
