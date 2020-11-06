package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	defaultColor      = "black"
	defaultListenAddr = "127.0.0.1:3000"
)

func main() {
	app := fiber.New()

	// define our routes
	app.Get("/", func(c *fiber.Ctx) error {
		// our main route will return the defined color
		if os.Getenv("COLOR") != "" {
			log.Printf("color request recieved.. replying with %v", os.Getenv("COLOR"))
			return c.SendString(os.Getenv("COLOR"))
		}
		log.Printf("color request recieved.. no color defined.. replying with default color of %v", defaultColor)
		return c.SendString(defaultColor)
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		log.Println("ping request recieved.. replying with PONG")
		return c.SendString("PONG")
	})

	app.Get("/bad", func(c *fiber.Ctx) error {
		log.Println("bad request recieved")
		return c.SendStatus(400)
	})

	app.Get("/fail", func(c *fiber.Ctx) error {
		log.Println("failure request recieved")
		return c.SendStatus(500)
	})

	app.Get("/slow", func(c *fiber.Ctx) error {
		var delay int
		delay = rand.Intn(10)
		log.Printf("slow request recieved.. delay set to %v", delay)
		time.Sleep(time.Duration(delay) * time.Second)
		if os.Getenv("COLOR") != "" {
			return c.SendString(os.Getenv("COLOR"))
		}
		return c.SendString(defaultColor)
	})

	app.Get("/slow/2", func(c *fiber.Ctx) error {
		var delay int
		delay = rand.Intn(60)
		log.Printf("really slow request recieved.. delay set to %v", delay)
		time.Sleep(time.Duration(delay) * time.Second)
		if os.Getenv("COLOR") != "" {
			return c.SendString(os.Getenv("COLOR"))
		}
		return c.SendString(defaultColor)
	})

	app.Get("/random", func(c *fiber.Ctx) error {
		var delay int
		var state int
		state = rand.Intn(3)
		delay = rand.Intn(5)
		log.Printf("random request recieved.. delay set to %v", delay)
		log.Printf("state set to %v", state)
		time.Sleep(time.Duration(delay) * time.Second)

		switch state {
		case int(1):
			return c.SendStatus(500)
		case int(2):
			return c.SendStatus(400)
		default:
			if os.Getenv("COLOR") != "" {
				return c.SendString(os.Getenv("COLOR"))
			}
			return c.SendString(defaultColor)
		}

	})

	if os.Getenv("LISTEN_ADDR") != "" {
		log.Fatal(app.Listen(os.Getenv("LISTEN_ADDR")))
	}
	log.Fatal(app.Listen(defaultListenAddr))
}
