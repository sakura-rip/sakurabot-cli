package utils

import (
	"github.com/spf13/cobra"
	"os"
)

func GetHomeDir() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return home
}
