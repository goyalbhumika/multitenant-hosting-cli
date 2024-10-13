package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

type App struct {
	Name       string `json:"name"`
	DeployType string `json:"deploy_type"`
}

var (
	name       string
	deployType string
)

func init() {
	CreateAppCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the app")
	CreateAppCmd.Flags().StringVarP(&deployType, "deploy_type", "t", "", "Name of the app")

}

type response struct {
	Name string `json:"name"`
	Port int    `json:"port"`
	DNS  string `json:"dns"`
}

var CreateAppCmd = &cobra.Command{
	Use:   "create-app",
	Short: "Create a new app",
	Run: func(cmd *cobra.Command, args []string) {
		err := createApp(name)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func createApp(name string) error {
	fmt.Println("Creating app", name)
	// Host of multitenant hosting platform
	url := "http://localhost:8080/v1/apps"

	app := &App{
		Name:       name,
		DeployType: deployType,
	}
	jsonData, err := json.Marshal(app)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	requestBody := bytes.NewBuffer(jsonData)
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		return fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Error deploying app, status code %d", resp.StatusCode)
	}

	respBody := response{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	fmt.Printf("App created successfully: %+v\n", respBody)
	return nil
}
