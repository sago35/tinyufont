package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	appName        = "tinyufontgen"
	appDescription = ""
)

type cli struct {
	outStream io.Writer
	errStream io.Writer
}

var (
	app = kingpin.New(appName, appDescription)
)

// Run ...
func (c *cli) Run(args []string) error {
	app.UsageWriter(c.errStream)

	if VERSION != "" {
		app.Version(fmt.Sprintf("%s version %s build %s", appName, VERSION, BUILDDATE))
	} else {
		app.Version(fmt.Sprintf("%s version - build -", appName))
	}
	app.HelpFlag.Short('h')

	str := app.Arg(`string`, `strings for font`).String()

	k, err := app.Parse(args[1:])
	if err != nil {
		return err
	}

	switch k {
	default:
		f := fontgen{
			pkgname:  `ayu20gothic`,
			fontname: `Ayu20gothic`,
			isofont:  `..\pyportal-private\tinyufont\tinyufont\ayu20gothic\10x20grkm.bdf`,
			font:     `..\pyportal-private\tinyufont\tinyufont\ayu20gothic\k20gm.bdf`,
		}

		w, err := os.Create(`.\ayu20gothic\ayu20gothic.go`)
		if err != nil {
			return nil
		}
		defer w.Close()
		f.generate(w, []rune(*str))
	}

	return nil
}
