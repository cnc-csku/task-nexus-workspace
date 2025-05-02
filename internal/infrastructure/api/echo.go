package api

import (
	"context"
	"log"

	"github.com/cnc-csku/task-nexus-go-lib/jsonvalidator"
	"github.com/cnc-csku/task-nexus-go-lib/logging"
	"github.com/cnc-csku/task-nexus-go-lib/utils/errutils"
	"github.com/cnc-csku/task-nexus-workspace/internal/config"
	"github.com/cnc-csku/task-nexus-workspace/internal/router"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EchoAPI struct {
	echo        *echo.Echo
	ctx         context.Context
	config      *config.Config
	mongoClient *mongo.Client
	router      *router.Router
}

func NewEchoAPI(
	ctx context.Context,
	config *config.Config,
	mongoClient *mongo.Client,
	router *router.Router,
) *EchoAPI {
	return &EchoAPI{
		echo:        echo.New(),
		ctx:         ctx,
		config:      config,
		mongoClient: mongoClient,
		router:      router,
	}
}

func (a *EchoAPI) Start(logger *logrus.Logger) error {
	e := echo.New()

	// Set up logger
	var formatter logrus.Formatter
	if a.config.LogFormat == "TEXT" {
		formatter = &logrus.TextFormatter{}
	} else {
		formatter = &logging.CustomFormatter{}
	}

	// Set up logging middleware
	e.Use(logging.EchoLoggingMiddleware(logger, formatter))

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: a.config.AllowOrigins,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		AllowMethods: []string{
			echo.GET,
			echo.PUT,
			echo.PATCH,
			echo.POST,
			echo.DELETE,
			echo.OPTIONS,
		},
	}))

	// Set up JSON validator
	e.Validator = jsonvalidator.NewValidator()

	// Custom error handler
	e.HTTPErrorHandler = errutils.CustomHTTPErrorHandler

	a.router.RegisterAPIRouter(e)

	err := e.Start(":" + a.config.RestPort)
	if err != nil {

		return err
	}

	defer func() {
		if err := a.mongoClient.Disconnect(a.ctx); err != nil {
			log.Printf("❌ Error disconnecting from MongoDB: %v\n", err)
		}
	}()

	return nil
}
