package main

import (
	"context"
	"fmt"
	"os"
	"github.com/urfave/cli/v3"
)

func main() {
	(&cli.Command{}).Run(context.Background(), os.Args)
	fmt.Println("Hello from Hexlet!")
}
