package main

import (
	app "github.com/oswaldom-code/api-template-gin/cmd/api"
	"github.com/oswaldom-code/api-template-gin/pkg/log"
)

func main() {
	log.Info("Starting application...")
	app.Execute()
}
