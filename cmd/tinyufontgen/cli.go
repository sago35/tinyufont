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

	pkgname := app.Flag("package", "package name").Default("main").String()
	fontname := app.Flag("fontname", "font name").Default("TinyUFont").String()
	isofont := app.Flag("ascii-font", "ascii font path (*.bdf)").ExistingFile()
	font := app.Flag("font", "font path (*.bdf) (jisx0208 only)").ExistingFile()
	str := app.Arg(`string`, `strings for font`).String()
	output := app.Flag("output", "output path").String()

	k, err := app.Parse(args[1:])
	if err != nil {
		return err
	}

	switch k {
	default:
		f := fontgen{
			pkgname:  *pkgname,
			fontname: *fontname,
			isofont:  *isofont,
			font:     *font,
		}

		w, err := os.Create(*output)
		if err != nil {
			return nil
		}
		defer w.Close()
		f.generate(w, []rune(*str))
	}

	return nil
}
