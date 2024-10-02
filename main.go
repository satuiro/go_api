package main

import (
	"context"
	"fmt"

	"github.com/satuiro/go_api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start application", err)
	}
}
