package topia

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

//func Testnima() cli.Command {
//	fmt.Println("nima")
//}
func CreateCommand() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "Create asfddass",
		Description: "The create command will generate an encrypted data store " +
			"placed at $HOME/.trousseau.tr or at the location described by " +
			"the $TROUSSEAU_HOME environment variable if you provided it.\n\n" +
			"   Encryption is made using your GPG main identity, and targets the " +
			"GPG recipients you provide as the command arguments.\n\n" +
			"   Examples:\n\n" +
			"     trousseau create 16DB4F3\n" +
			"     trousseau create tcrevon@gmail.com\n" +
			"     export TROUSSEAU_STORE=/tmp/test_trousseau.tr && trousseau create 16DB4F3\n",
		Action: func(c *cli.Context) error {
			//var encryptionType string = c.String("encryption-type")
			//
			//if encryptionType == trousseau.SYMMETRIC_ENCRYPTION_REPR {
			//	err := trousseau.CreateAction(trousseau.SYMMETRIC_ENCRYPTION, trousseau.AES_256_ENCRYPTION, nil)
			//	if err != nil {
			//		trousseau.ErrorLogger.Fatal(err)
			//	}
			//} else {
			//	if len(c.Args()) > 0 {
			//		err := trousseau.CreateAction(trousseau.ASYMMETRIC_ENCRYPTION, trousseau.GPG_ENCRYPTION, c.Args())
			//		if err != nil {
			//			trousseau.ErrorLogger.Fatal(err)
			//		}
			//	} else {
			//		trousseau.ErrorLogger.Fatal("invalid number of arguments provided to " +
			//			"the create command. At least one recipient to encrypt the " +
			//			"data store for is needed.")
			//	}
			//}
			//
			//trousseau.InfoLogger.Println("Trousseau data store succesfully created")
			fmt.Println("caonima")
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name: "encryption-type",
				Usage: "Define the encryption type to be used for store encryption. " +
					"Whether symmetric or asymmetric.",
				Value: trousseau.ASYMMETRIC_ENCRYPTION_REPR,
			},
			cli.StringFlag{
				Name: "encryption-algorithm",
				Usage: "Define the algorithm to be used for store encryption. " +
					"Whether gpg or aes.",
				Value: trousseau.GPG_ENCRYPTION_REPR,
			},
		},
	}
}
