package main

import (
	"os"

	"github.com/dlacreme/did-events/api"
	"github.com/dlacreme/did-events/didgrpc"
	"github.com/dlacreme/did-events/env"
	"github.com/labstack/echo/v4"
)

func main() {
	env.LoadAndCheck()
	e := echo.New()
	api.RegisterMiddlewares(e)
	api.RegisterEndpoints(e)
	go didgrpc.StartServer()
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
