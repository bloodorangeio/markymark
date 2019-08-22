package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli"

	"github.com/bloodorangeio/markymark/pkg/markymark"
)

var (
	Version  string
	Revision string

	ErrorNoPipe        = errors.New("this command is intended to work with pipes (|)")
	ErrorMissingRefArg = errors.New("missing arguments: [ref]")
)

func main() {
	app := cli.NewApp()
	app.Name = "mm (markymark)"
	app.Version = fmt.Sprintf("%s (build %s)", Version, Revision)
	app.Usage = "markdown storage over oci"
	app.Commands = []cli.Command{
		{
			Name:  "push",
			Usage: "upload a markdown file to registry from stdin",
			Action: func(c *cli.Context) error {
				ref := c.Args().Get(0)
				if ref == "" {
					return ErrorMissingRefArg
				}
				content, err := getStdin()
				if err != nil {
					return err
				}
				return markymark.Push(content, ref)
			},
		},
		{
			Name:  "pull",
			Usage: "download a markdown file from registry and print to stdout",
			Action: func(c *cli.Context) error {
				ref := c.Args().Get(0)
				if ref == "" {
					return ErrorMissingRefArg
				}
				return markymark.Pull(ref)
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getStdin() ([]byte, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return nil, ErrorNoPipe
	}
	reader := bufio.NewReader(os.Stdin)
	var output []rune
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	return []byte(string(output)), nil
}
