package cmd

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/t-eckert/elevate/controller"
)

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "TODO",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatal(err.Error())
		}

		c := controller.NewController()

		log.Infof("Serving Controller at http://localhost:%d\n", port)
		http.HandleFunc("/", controller.NewIndexHandler(c))
		http.HandleFunc("/elevators", controller.NewElevatorHandler(c))
		http.HandleFunc("/passengers", controller.NewPassengerHandler(c))
		log.Error(http.ListenAndServe(fmt.Sprintf(":%d", port), nil).Error())
	},
}

func init() {
	runCmd.AddCommand(controllerCmd)

	controllerCmd.Flags().IntP("port", "p", 3000, "Port to serve the controller on.")
}
