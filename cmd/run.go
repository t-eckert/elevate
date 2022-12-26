package cmd

import (
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an Elevate module: elevator, controller, or queuer.",
	Long: `The 'run' subcommand exposes the individual modules of Elevate:
- elevator: an elevator with an exposed API. Takes passengers and ferries them 
  to their destinations.
- controller: a module which controls a fleet of elevators. Passengers sent to 
  a controller will be assigned to an elevator that is most suited to minimize
  the length of their journey.
- queuer: a service which creates new passengers and sends them to a configured
  endpoint at regular intervals.`,
}

func init() {
	rootCmd.AddCommand(runCmd)
}
