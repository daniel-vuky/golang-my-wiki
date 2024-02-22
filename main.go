package main

import (
	"context"
	"github.com/daniel-vuky/golang-my-wiki/application"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
)

func main() {
	loadEnvConfigErr := godotenv.Load()
	if loadEnvConfigErr != nil {
		log.Fatalln("Error when loading env config!")
	}
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		log.Fatalf("Can not start the server, %s", err.Error())
	}
}
