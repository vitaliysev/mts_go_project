package main

import (
	"context"
	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/hotel/app"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app", err)
	}
}
