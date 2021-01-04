package webserver

import (
	"fmt"
	"github.com/Samb8104/Restore/internal/api"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Web struct {
	app            *fiber.App
	servedPages    int
	startTime      time.Time
	servedRequests int
	port           int
}

func Initialise(port int) *Web {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})

	//Load API routers
	api.Load(app)
	app.Listen(fmt.Sprintf(":%d", port))

	return &Web{
		app:            app,
		servedPages:    0,
		startTime:      time.Now(),
		servedRequests: 0,
		port:           port,
	}
}
