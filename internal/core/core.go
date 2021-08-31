package core

import (
	"github.com/sakura-rip/sakurabot-cli/internal/commands/all"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
)

func ExecuteCommands() error {
	root, err := all.BuildCommands()
	if err != nil {
		utils.Logger.Error().Err(err).Msg("failed to build commands")
	}
	return root.Execute()
}
