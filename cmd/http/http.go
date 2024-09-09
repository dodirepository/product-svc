package http

import (
	"context"
	"os"

	"github.com/dodirepository/product-svc/infrastructure/database"
	"github.com/dodirepository/product-svc/infrastructure/http"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
}

// Start :nodoc:
func Start(ctx context.Context) {
	httpServer := http.NewHTTPServer()
	defer httpServer.Done()
	database.GetConnection()

	logrus.Infof("http server start on port %s", os.Getenv("APP_PORT"))
	if err := httpServer.Run(ctx); err != nil {
		logrus.Info("http server stopped")
	}
}
