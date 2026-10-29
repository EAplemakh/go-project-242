// Package main is the entry point for the hexlet-path-size CLI utility.
package main

import (
	"code/pkg/convertunits"
	"code/pkg/pathsize"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var (
		filePath      string
		chooseUnits   bool
		includeHidden bool
		recursive     bool
	)

	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "human",
				Aliases:     []string{"H"},
				Value:       false,
				Usage:       "human-readable sizes (auto-select unit)",
				Destination: &chooseUnits,
			},
			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"a"},
				Value:       false,
				Usage:       "include hidden files and directories",
				Destination: &includeHidden,
			},
			&cli.BoolFlag{
				Name:        "recursive",
				Aliases:     []string{"r"},
				Value:       false,
				Usage:       "recursive size of directories",
				Destination: &recursive,
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "path",
				Destination: &filePath,
			},
		},
		HideHelpCommand: true,
		Action: func(_ context.Context, _ *cli.Command) error {
			if filePath == "" {
				return errors.New("ошибка: необходимо указать путь к файлу или директории")
			}

			res, err := pathsize.GetPathSize(filePath, includeHidden, recursive)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", convertunits.FormatUnitString(res, chooseUnits), filePath)

			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
