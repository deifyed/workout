package data

import (
	"fmt"
	"github.com/spf13/afero"
)

func ensureOpen(fs *afero.Afero, dbPath string) (afero.File, error) {
	exists, err := fs.Exists(dbPath)
	if err != nil {
		return nil, fmt.Errorf("checking existence: %w", err)
	}

	var f afero.File

	if !exists {
		f, err = fs.Create(dbPath)
		if err != nil {
			return nil, fmt.Errorf("creating database: %w", err)
		}

		_, err = f.WriteString("date")
		if err != nil {
			return nil, fmt.Errorf("writing date header: %w", err)
		}
	} else {
		f, err = fs.Open(dbPath)
		if err != nil {
			return nil, fmt.Errorf("opening file: %w", err)
		}
	}

	return f, nil
}

func contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}
