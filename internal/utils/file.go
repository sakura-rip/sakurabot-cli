package utils

import (
	"bufio"
	"os"
)

func ReadFileLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		Logger.Fatal().Err(err).Msg("failed open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		Logger.Fatal().Err(err).Msg("scanner error")
	}
	return lines
}
