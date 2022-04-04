package config

import (
	"fmt"
	"os"
	"path"
)

func GetDatabasePath() (string, error) {
	if dbPath := os.Getenv("WORKOUT_DATABASE_PATH"); dbPath != "" {
		return dbPath, nil
	}

	baseDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("acquiring config directory: %w", err)
	}

	return path.Join(baseDir, "workout", "workout_data.csv"), nil
}
