package cmd

import (
	"fmt"
	"github.com/deifyed/workout/pkg/plot"
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Show a graph of workouts",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := plot.Open("/home/deifyed/life/Notes/notes/training/data.csv")
		if err != nil {
			return fmt.Errorf("opening plot: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}
