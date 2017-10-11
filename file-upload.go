package main

import(
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"log"
	"github.com/nlopes/slack"
	"os"
)

func fileUpload(c *cli.Context) error {
	sourceFile := c.Args().First()
	token := c.GlobalString("token")
	channels := []string{c.String("channel")}

	// Listening slack event and response
	log.Printf("[INFO] Start slack event listening")
	api := slack.New(token)
	params := slack.FileUploadParameters{
		Title: sourceFile,
		File: sourceFile,
		Filename: sourceFile,
		Channels: channels,
	}

	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivate)
	return nil
}
