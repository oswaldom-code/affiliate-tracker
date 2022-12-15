package main

import (
	app "github.com/oswaldom-code/affiliate-tracker/cmd/api"
	"github.com/oswaldom-code/affiliate-tracker/pkg/log"
)

func main() {
	log.Info("Starting application...")
	app.Execute()
}
