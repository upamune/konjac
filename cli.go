package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"io"
	"log"
	"github.com/theplant/bingtranslator/translator"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

type Config struct {
	Client Client `toml:"client"`
}

func (c *Config) loadConfig(fileName string) error {
	_, err := toml.DecodeFile(fileName, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) validate() error {
	errorMessage := ""
	if c.Client.Id == "" {
		errorMessage += "Id not found\n"
	}
	if c.Client.Secret == "" {
		errorMessage += "Secret not found\n"
	}

	if errorMessage == "" {
		return nil
	} else {
		return errors.New(errorMessage)
	}
}

type Client struct {
	Id     string `toml:"id"`
	Secret string `toml:"secret"`
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		c string

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&c, "config", "", "")
	flags.StringVar(&c, "c", "", "(Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if c == "" {
		log.Fatal("please set a setting file with c option")
		return ExitCodeError
	}

	_ = c



	if flags.NArg() == 0 {
		log.Fatal("please set translate text")
		return ExitCodeError
	}

	var text string

	for _, arg := range flags.Args() {
		text += " " + arg
	}

	config := Config{}
	err := config.loadConfig(c)
	if err != nil {
		log.Fatal(err)
		return ExitCodeError
	}
	err = config.validate()
	if err != nil {
		log.Fatal(err)
		return ExitCodeError
	}

	fmt.Println(text)

	bingtranslator.SetCredentials(config.Client.Id, config.Client.Secret)
	translations, err := bingtranslator.Translate("ja", "en", text, "INPUT_TEXT")
	if err != nil {
		log.Fatal(err)
		return ExitCodeError
	}

	fmt.Println(translations)

	return ExitCodeOK
}
