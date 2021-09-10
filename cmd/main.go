package main

import (
	"github.com/joho/godotenv"
	"github.com/sakura-rip/sakurabot-cli/internal/core"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		utils.Fatal().Err(err).Msg("")
	}
	if err := core.ExecuteCommands(); err != nil {
		os.Exit(1)
	}
}
