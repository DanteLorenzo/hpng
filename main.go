package main

import (
	"log"
	"os"
	"github.com/DanteLorenzo/hpng/hidepng"
	"github.com/urfave/cli/v2"
)

func main() {

	var sourceImage, hiddeFile string
	var reverse bool

	app := &cli.App{
		Name:           "hpng",
		Usage:          "Hide FILE to PNG",
		Version:        "v1.0.0",
		Description:    "",
		UsageText: ("hpng [global options] [arguments...]\n" +
					"Example: hpng -s IMAGE -f FILE // Hide file\n" +
					"Example: hpng -r -s IMAGE -f NEW_FILE // Find file"),
		Flags: []cli.Flag{

			// "Reverse" Flag - take file from png
			&cli.BoolFlag{
				Name:        "reverse",
				Aliases:     []string{"r"},
				Required:    false,
				Destination: &reverse,
				Usage:       "take file from image"},

			// "Source" Flag to choose image where will be hide file
			&cli.StringFlag{
				Name:        "source",
				Aliases:     []string{"s"},
				Required:    true,
				Destination: &sourceImage,
				Usage:       "hide or take file from this image"},

			// "File" Flag - choose file to hide into image
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Required:    false,
				Destination: &hiddeFile,
				Usage:       "file to hide or output"},
			
		},
		EnableBashCompletion: true,
		HideHelp:             false,
		HideHelpCommand:      true,
		HideVersion:          true,
		UseShortOptionHandling:    false,
		SkipFlagParsing:           false,
		Action: func(cCtx *cli.Context) error {
			if !reverse {
				hpng.CreateHPNG(sourceImage, hiddeFile)
			} else {
				hpng.ReverseHPNG(sourceImage, hiddeFile)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

