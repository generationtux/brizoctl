package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/urfave/cli"
)

// PromoteApplication moves an application from one environment to another
func PromoteApplication(c *cli.Context) error {
	type environment struct {
		Name string `json:"name"`
		UUID string `json:"uuid"`
	}

	type appDetails struct {
		Name         string        `json:"name"`
		UUID         string        `json:"uuid"`
		Environments []environment `json:"environments"`
	}

	appName := c.Args().Get(0)
	envSrcName := c.Args().Get(1)
	envDstName := c.Args().Get(2)
	var envSrc environment
	var envDst environment

	body, err := HTTPGet("/api/v1/applications/" + appName)
	if err != nil {
		return err
	}

	var app appDetails
	err = json.Unmarshal([]byte(body), &app)
	if err != nil {
		return errors.New("Unable to read JSON response.")
	}

	// get environment
	for _, v := range app.Environments {
		if v.Name == envSrcName {
			envSrc = v
		}

		if v.Name == envDstName {
			envDst = v
		}
	}

	_, err = HTTPPost("/api/v1/environments/promote/" + envSrc.UUID + "/" + envDst.UUID)
	if err != nil {
		return err
	}

	fmt.Println("Promotion successfully")

	return nil
}
