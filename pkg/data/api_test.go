package data_test

import (
	"path"
	"testing"
	"time"

	"github.com/deifyed/workout/pkg/data"
	"github.com/sebdah/goldie/v2"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestDataClient_Insert(t *testing.T) {
	testCases := []struct {
		name         string
		withWorkouts []data.Workout
	}{
		{
			name: "Should have an expected line in the file after inserting a workout",
			withWorkouts: []data.Workout{
				{
					Date:        time.Date(2022, time.November, 10, 0, 0, 0, 0, time.UTC),
					Type:        "pushups",
					Repetitions: "3",
				},
			},
		},
		{
			name: "Should support multiple days of workouts",
			withWorkouts: []data.Workout{
				{
					Date:        time.Date(2022, time.November, 10, 0, 0, 0, 0, time.UTC),
					Type:        "pushups",
					Repetitions: "3",
				},
				{
					Date:        time.Date(2022, time.November, 11, 0, 0, 0, 0, time.UTC),
					Type:        "pushups",
					Repetitions: "4",
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			fs := &afero.Afero{Fs: afero.NewMemMapFs()}
			dataFilepath := path.Join("/", "data.csv")

			client := data.NewClient(fs, dataFilepath)

			for _, workout := range tc.withWorkouts {
				err := client.Insert(workout)
				assert.NoError(t, err)
			}

			content, err := fs.ReadFile(dataFilepath)
			assert.NoError(t, err)

			g := goldie.New(t)
			g.Assert(t, tc.name, content)
		})
	}
}
