package file

import (
	"bufio"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/spf13/cobra"
	"os"
)

func ReadFileLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Fatal().Err(err).Msg("scanner error")
	}
	return lines
}

func ReadAll(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed open file")
	}
	return file
}

func GetHomeDir() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return home
}
