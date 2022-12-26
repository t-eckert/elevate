/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/t-eckert/elevate/passenger"
)

const (
	localhost = "http://localhost"
)

// queueCmd represents the queue command
var queueCmd = &cobra.Command{
	Use:   "enqueue",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		p := passenger.NewRandomPassenger()

		b, err := json.Marshal(p)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		resp, err := http.Post(fmt.Sprintf("%s%s/passengers", localhost, port), "application/json", bytes.NewBuffer(b))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.Status)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(queueCmd)

	queueCmd.Flags().StringP("port", "p", ":3000", "The port where the passenger will be sent.")
}
