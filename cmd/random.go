/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Fetch a random joke",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("random called")
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	ID     string `json:id`
	JOKE   string `json:joke`
	Status int    `json:status`
}

func getRandomJoke() {

}

func getJokeData(baseAPI string) []byte {
	http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not get a dadjoke -%v", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/yrs147/go-cli)")

	response, err := http.DefailtClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request -%v", err)
	}

	responseBytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body -%v", err)
	}
}
