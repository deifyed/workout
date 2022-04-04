package cmd

import (
	"fmt"
	"github.com/deifyed/workout/pkg/config"
	"github.com/deifyed/workout/pkg/data"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"time"
)

var addCmd = &cobra.Command{
	Use:   "add [workout type] [repetitions]",
	Short: "Add workout data.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		dbPath, err := config.GetDatabasePath()
		if err != nil {
			return fmt.Errorf("acquiring database path: %w", err)
		}

		dataClient := data.NewClient(&afero.Afero{Fs: afero.NewOsFs()}, dbPath)

		workout := data.Workout{
			Date:        time.Now(),
			Type:        args[0],
			Repetitions: args[1],
		}

		err = dataClient.Insert(workout)
		if err != nil {
			return fmt.Errorf("saving workout: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
