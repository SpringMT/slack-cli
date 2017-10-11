package main

import(
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "slack"
	app.Version = Version

	flags := []cli.Flag {
		cli.StringFlag{
			Name:  "token, t",
			Usage: "Slack Token",
			EnvVar: "TOKEN",
		},
	}
	app.Flags = flags

	app.Commands = []cli.Command{
		{
			Name:    "file-upload",
			Usage:   "Upload a file",
			Action:  fileUpload,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "channel, c",
					Usage: "Channel",
				},
			},
		},
	}
	app.Run(os.Args)
}