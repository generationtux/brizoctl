package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

var listCommands = []cli.Command{
	cli.Command{
		Name:   "apps",
		Usage:  "list all applications",
		Action: ListApplications,
	},
}

// ListApplications retrieves a list of applications
func ListApplications(c *cli.Context) error {
	type appList struct {
		Name string `json:"name"`
		UUID string `json:"uuid"`
	}

	body, err := HTTPGet("/api/v1/applications")
	if err != nil {
		return err
	}

	var apps []appList
	err = json.Unmarshal([]byte(body), &apps)
	if err != nil {
		return errors.New("Unable to read JSON response.")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "UUID"})
	for _, app := range apps {
		table.Append([]string{app.Name, app.UUID})
	}
	table.Render()

	return nil
}
