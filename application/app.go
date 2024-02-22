package application

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"time"
)

type App struct {
	router *gin.Engine
	rdb    *sql.DB
}

// New
// Init new application
// Init connection to database
// Load all the routes
func New() *App {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER_NAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE_NAME"),
	)
	db, connectedErr := sql.Open("mysql", connectionString)
	if connectedErr != nil {
		panic(
			fmt.Sprintf(
				"Can not connect to mysql server, %s",
				connectedErr.Error(),
			),
		)
	}
	app := &App{
		rdb: db,
	}

	app.LoadRoutes()

	return app
}

// Start
// Launching the server
// Init http server with port from env
// Ping to database and check if connection still opened
// Add defer for checking if the connection is closed successfully
// Listen and serve the server
// Listen event done from ctx and shutdown the server
func (app *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APPLICATION_PORT")),
		Handler: app.router,
	}

	pingErr := app.rdb.Ping()
	if pingErr != nil {
		return fmt.Errorf(
			"can not ping to database, %s",
			pingErr.Error(),
		)
	}
	defer func() {
		if err := app.rdb.Close(); err != nil {
			fmt.Printf(
				"can not close the database connection, %s",
				err.Error(),
			)
		}
	}()

	fmt.Println("Starting server")

	serverErr := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			serverErr <- fmt.Errorf("fail to start the server, %s", err.Error())
		}
	}()

	select {
	case err := <-serverErr:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

	return nil

}
