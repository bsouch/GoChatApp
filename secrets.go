package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getSecret(name string) (string, error) {
	secretPath := os.Getenv("SECRET_PATH")
	path := filepath.Join(secretPath, name)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Path to secrets missing: %v", path)
		return os.Getenv(strings.ToUpper(name)), err
	}

	return strings.TrimSpace(string(data)), nil
}
