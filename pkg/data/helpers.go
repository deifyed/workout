package data

import (
	"fmt"
	"github.com/spf13/afero"
	"os"
)

func ensureOpen(fs *afero.Afero, dbPath string) (afero.File, error) {
	f, err := fs.OpenFile(dbPath, os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		return nil, fmt.Errorf("opening: %w", err)
	}

	info, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("stating: %w", err)
	}

	if info.Size() == 0 {
		_, err = f.WriteString("date")
		if err != nil {
			return nil, fmt.Errorf("writing date header: %w", err)
		}
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return nil, fmt.Errorf("seeking after writing header: %w", err)
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
