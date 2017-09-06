package main

import (
	"github.com/mohakkataria/redbubble_assignment/batchprocessor"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "RedBubble Batch Processor"
	app.Usage = "A batch processor to process image data"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Value: "http://take-home-test.herokuapp.com/api/v1/works.xml",
			Usage: "The endpoint of the API",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "./output_directory",
			Usage: "The directory where files will be generated.",
		},
	}

	app.Action = func(c *cli.Context) error {
		batchprocessor.Execute(c.String("url"), c.String("output"))
		return nil
	}
	app.Run(os.Args)

}
