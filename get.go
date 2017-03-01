package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

var getCommands = []cli.Command{
	cli.Command{
		Name:      "apps",
		Aliases:   []string{"app"},
		Usage:     "Get application details",
		UsageText: "get apps [UUID]",
		Action:    GetApplication,
	},
}

// GetApplication retrieves a list of applications
func GetApplication(c *cli.Context) error {
	uuid := c.Args().Get(0)
	if uuid == "" {
		return errors.New(c.Command.UsageText)
	}

	type environment struct {
		Name string `json:"name"`
		UUID string `json:"uuid"`
	}

	type appDetails struct {
		Name         string        `json:"name"`
		UUID         string        `json:"uuid"`
		Environments []environment `json:"environments"`
	}

	body, err := HTTPGet("/api/v1/applications/" + uuid)
	if err != nil {
		return err
	}

	var app appDetails
	err = json.Unmarshal([]byte(body), &app)
	if err != nil {
		return errors.New("Unable to read JSON response.")
	}

	fmt.Println("Name: " + app.Name)
	fmt.Println("UUID: " + app.UUID)
	table := tablewriter.NewWriter(os.Stdout)
	fmt.Println("\nEnvironments")
	table.SetHeader([]string{"Name", "UUID"})
	for _, env := range app.Environments {
		table.Append([]string{env.Name, env.UUID})
	}
	table.Render()

	return nil
}
