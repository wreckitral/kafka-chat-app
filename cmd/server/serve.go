package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/wreckitral/kafka-chat-app/pkg/config"
	"github.com/wreckitral/kafka-chat-app/pkg/middleware"
	"github.com/wreckitral/kafka-chat-app/pkg/route"
	"github.com/wreckitral/kafka-chat-app/platform/logger"
)

func Serve() {
    appCfg := config.AppCfg()

    logger.SetUpLogger()
    logr := logger.GetLogger()

    fiberCfg := config.FiberConfig()
    app := fiber.New(fiberCfg)

    middleware.FiberMiddleware(app)

    route.GeneralRoute(app)
    route.SwaggerRoute(app)
    route.NotFoundRoute(app)

    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

    go func() {
        <- sigCh
        logr.Infoln("Shutting down server...")
        _ = app.Shutdown()
    }()

    serverAddr := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
    if err := app.Listen(serverAddr); err != nil {
        logr.Errorf("Server is NOT running, error: %v", err)
    }
}
