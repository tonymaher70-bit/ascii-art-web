package main

import (
	"os"
	"strings"
)

const charHeight = 8

func loadBanner(path string) (map[rune][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(
		strings.ReplaceAll(string(data), "\r\n", "\n"),
		"\n",
	)

	banner := make(map[rune][]string)

	ch := 32

	for i := 0; i < len(lines); {
		if lines[i] == "" {
			i++
			continue
		}

		if i+charHeight > len(lines) {
			break
		}

		banner[rune(ch)] = lines[i : i+charHeight]

		i += charHeight
		ch++
	}

	return banner, nil
}

func generateASCII(text string, banner map[rune][]string) string {
	var result strings.Builder

	for line := 0; line < charHeight; line++ {
		for _, ch := range text {
			if art, ok := banner[ch]; ok {
				result.WriteString(art[line])
			}
		}
		result.WriteString("\n")
	}

	return result.String()
}

func GenerateASCII(text, bannerName string) (string, error) {
	bannerFile := bannerName + ".txt"

	banner, err := loadBanner(bannerFile)
	if err != nil {
		return "", err
	}

	var result strings.Builder

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		result.WriteString(generateASCII(line, banner))
		result.WriteString("\n")
	}

	return result.String(), nil
}
