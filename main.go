package main

import (
	"os"

	"blog.com/config"
	"blog.com/pkg/cmd"
	"blog.com/router"
)

func main() {
	r := router.SetupRouter()
	args := os.Args

	if len(args) > 1 {
		cmd.Execute()
		os.Exit(1)
	}

	r.Run(config.GetEnvWithKey("APP_DOMAIN", "localhost") + ":" + config.GetEnvWithKey("APP_PORT", "8080"))
}
