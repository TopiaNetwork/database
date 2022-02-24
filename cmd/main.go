package main

import (
	"os"
	// "path"

	"github.com/urfave/cli"
	"database/cmd/topia"
)

func main() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "print only the executable's version",
	}

	app := cli.NewApp()

	app.Name = "topia"
	app.Author = "tao"
	app.Email = "tao@gmail.com"
	app.Usage = "database for mengmen"
	app.Version = "1.0"
	app.Commands = []cli.Command{
		// CreateCommand(),
		// SetCommand(),
		// GetCommand(),
		// RenameCommand(),
		// DelCommand(),
		// KeysCommand(),
		// ShowCommand(),
		// ExportCommand(),
		// ImportCommand(),
		// PushCommand(),
		// PullCommand(),
		// ListRecipientsCommand(),
		// AddRecipientCommand(),
		// RemoveRecipientCommand(),
		// MetaCommand(),
		// UpgradeCommand(),
		Testnima(),
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "good",
			Usage: "Filename of config file to override default lookup",
		},
		cli.StringFlag{
			Name:  "store, s",
			Usage: "Path to the trousseau data store to use",
		},
		cli.BoolFlag{
			Name:  "ask-passphrase",
			Usage: "Have trousseu prompt user for passphrase",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Set trousseau in verbose mode",
		},
		cli.StringFlag{
			Name:  "gnupg-home",
			Usage: "Provide an alternate gnupg home",
		},
	}

	app.EnableBashCompletion = true
	app.HideHelp = false
	app.HideVersion = false

	// app.Before = Before
	app.Run(os.Args)
}
