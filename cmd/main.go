package main

import (
	"fmt"
	"github.com/EldersJavas/qhash"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:        "Qhash cmd",
		Version:     "1.0",
		Description: "Quick hash tool",
		Flags: []cli.Flag{

			&cli.BoolFlag{
				Name:    "IsPure",
				Usage:   "Print hash purely",
				Value:   false,
				Aliases: []string{"p"},
			},
			&cli.BoolFlag{
				Name:    "Checksumfile",
				Usage:   "Generate checksum file",
				Value:   false,
				Aliases: []string{"g"},
			},
			&cli.BoolFlag{
				Name:    "String",
				Usage:   "get string hash",
				Value:   false,
				Aliases: []string{"s"},
			},
			&cli.BoolFlag{
				Name:    "File",
				Usage:   "get string hash",
				Value:   false,
				Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:        "Str",
				Usage:       "String content",
				Value:       "TEST",
				Destination: nil,
				Aliases:     []string{"str"},
			},
			&cli.StringFlag{
				Name:        "FilePath",
				Usage:       "FilePath",
				Value:       "TEST",
				Destination: nil,
				Aliases:     []string{"fp"},
			},
			&cli.BoolFlag{
				Name:  "MD5",
				Usage: "foo greeting",
				Value: false,
				Action: func(context *cli.Context, b bool) error {
					if b {
						if context.Bool("") {

						}
						fmt.Printf("MD5: %s", qhash.MD5([]byte("")))
					}
					return nil
				},
			},
		},
		EnableBashCompletion: true,
		Before:               nil,
		After:                nil,
		Action: func(cCtx *cli.Context) error {

			return nil
		},
		Authors:   []*cli.Author{{Name: "EldersJavas", Email: "eldersjavas@gmail.com"}},
		Copyright: "MIT License\n\nCopyright (c) 2022 EldersJavas",
		Suggest:   true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
func GetCliFlag() (Flags []cli.Flag) {
	flags1 := []cli.Flag{
		&cli.BoolFlag{
			Name:    "IsPure",
			Usage:   "Print hash purely",
			Value:   false,
			Aliases: []string{"p"},
		},
		&cli.BoolFlag{
			Name:    "Checksumfile",
			Usage:   "Generate checksum file",
			Value:   false,
			Aliases: []string{"g"},
		},
		&cli.BoolFlag{
			Name:    "String",
			Usage:   "get string hash",
			Value:   false,
			Aliases: []string{"s"},
		},
		&cli.BoolFlag{
			Name:    "File",
			Usage:   "get string hash",
			Value:   false,
			Aliases: []string{"f"},
		},
		&cli.StringFlag{
			Name:        "Str",
			Usage:       "String content",
			Value:       "TEST",
			Destination: nil,
			Aliases:     []string{"str"},
		},
		&cli.StringFlag{
			Name:        "FilePath",
			Usage:       "FilePath",
			Value:       "TEST",
			Destination: nil,
			Aliases:     []string{"fp"},
		},
		&cli.BoolFlag{
			Name:  "MD5",
			Usage: "foo greeting",
			Value: false,
			Action: func(context *cli.Context, b bool) error {
				if b {
					if context.Bool("") {

					}
					fmt.Printf("MD5: %s", qhash.MD5([]byte("")))
				}
				return nil
			},
		},
	}

	hs := map[string]func([]byte) string{
		"MD4": qhash.MD4,
	}
	var Hsflag []cli.Flag

	for s, f := range hs {
		c := &cli.BoolFlag{
			Name:  s,
			Usage: "Get " + s,
			Value: false,
			Action: func(context *cli.Context, b bool) error {
				f([]byte(""))
				return nil
			},
		}
		Hsflag = append(Hsflag, c)

	}

	Flags = append(Flags, flags1...)
	Flags = append(Flags, Hsflag...)
	return
}

func GetHash(tp string) string {

	return ""
}
