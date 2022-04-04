package cmd

import (
	"fmt"
	"github.com/deifyed/workout/pkg/config"
	"github.com/deifyed/workout/pkg/plot"
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Show a graph of workouts.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		dbPath, err := config.GetDatabasePath()
		if err != nil {
			return fmt.Errorf("acquiring database path: %w", err)
		}

		err = plot.Open(dbPath)
		if err != nil {
			return fmt.Errorf("opening plot: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}
