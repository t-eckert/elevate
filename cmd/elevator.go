/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/t-eckert/elevate/elevator"
)

// elevatorCmd represents the elevator command
var elevatorCmd = &cobra.Command{
	Use:   "elevator",
	Short: "TODO",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		e := elevator.NewElevator("127.0.0.1")

		fmt.Printf("Serving Elevator at http://localhost%s\n", port)
		http.HandleFunc("/passengers", elevator.NewPassengerHandler(e))
		fmt.Println(http.ListenAndServe(port, nil).Error())
	},
}

func init() {
	runCmd.AddCommand(elevatorCmd)

	elevatorCmd.Flags().StringP("port", "p", ":3000", "Port to serve the elevator on.")
}
