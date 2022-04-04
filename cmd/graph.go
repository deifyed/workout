package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Show a graph of workouts",
	RunE: func(cmd *cobra.Command, args []string) error {
		gnucmd := exec.Command("gnuplot", "-p", "/home/deifyed/life/Notes/notes/training/cfg.plot")

		err := gnucmd.Run()
		if err != nil {
			return fmt.Errorf("running command: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}
