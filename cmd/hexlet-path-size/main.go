// Package main is the entry point for the hexlet-path-size CLI utility.
package main

import (
	"code/pkg"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var filePath string

	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "path",
				Destination: &filePath,
			},
		},
		HideHelpCommand: true,
		Action: func(_ context.Context, _ *cli.Command) error {
			if filePath == "" {
				return fmt.Errorf("ошибка: необходимо указать путь к файлу или директории")
			}

			res, err := pkg.GetPathSize(filePath)
			if err != nil {
				return err
			}

			fmt.Printf("%dB\t%s\n", res, filePath)

			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
