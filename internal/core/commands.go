package core

import (
	"github.com/sakura-rip/sakurabot-cli/internal/commands/all"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
)

// ExecuteCommands execute all commands
func ExecuteCommands() error {
	root, err := all.BuildAllCommands()
	if err != nil {
		logger.Error().Err(err).Msg("failed to build commands")
	}
	return root.Execute()
}
