package cmd

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/t-eckert/elevate/elevator"
)

// elevatorCmd represents the elevator command
var elevatorCmd = &cobra.Command{
	Use:   "elevator",
	Short: "TODO",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatal(err.Error())
		}

		config := *elevator.NewConfig().WithPort(port)
		e := elevator.NewElevator(config)

		elevator.Serve(cmd.Context(), e)

		log.Infof("Serving Elevator at http://localhost:%d\n", port)
		http.HandleFunc("/", elevator.NewIndexHandler(e))
		http.HandleFunc("/passengers", elevator.NewPassengerHandler(e))
		log.Error(http.ListenAndServe(fmt.Sprintf(":%d", port), nil).Error())
	},
}

func init() {
	runCmd.AddCommand(elevatorCmd)

	elevatorCmd.Flags().IntP("port", "p", 4000, "Port to serve the elevator on.")
}
