package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	bindPort      int
	fileDirectory string
	enableBrowse  bool
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	flag.IntVar(&bindPort, "bindPort", 3000, "The port that this application will listen on")
	flag.StringVar(&fileDirectory, "fileDirectory", "./", "")
	flag.BoolVar(&enableBrowse, "enableBrowse", true, "Enable directory browsing")

	app.Use(filesystem.New(filesystem.Config{
		Root:   http.Dir(fileDirectory),
		Browse: enableBrowse,
		MaxAge: 3600,
	}))

	app.Listen(fmt.Sprintf(":%d", bindPort))
}
