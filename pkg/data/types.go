package data

import (
	"time"

	"github.com/spf13/afero"
)

type DataClient struct {
	fs     *afero.Afero
	dbPath string
}

type Workout struct {
	Date        time.Time
	Type        string
	Repetitions string
}
