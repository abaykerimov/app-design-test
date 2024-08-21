package internal

import (
	"applicationDesignTest/handlers"
	"applicationDesignTest/packages/log"
	"net/http"
)

type Application struct {
	HttpServer *http.Server
	container  *Container
	Logger     *log.Logger
}

func InitializeApplication() (*Application, error) {
	app, err := BuildApplication()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func BuildApplication() (*Application, error) {
	mux := http.NewServeMux()

	srv := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	logger := log.NewLogger()

	container := NewContainer()

	handler := handlers.NewHandler(container.service, logger)
	handler.LoadRoutes(mux)

	application := &Application{
		HttpServer: srv,
		container:  container,
		Logger:     logger,
	}

	return application, nil
}

func (a *Application) Start() error {
	a.Logger.LogInfoF("Server listening on localhost:8080")

	if err := a.HttpServer.ListenAndServe(); err != nil {
		a.Logger.LogFatalF("failed to listen: %v", err)
	}

	return nil
}
