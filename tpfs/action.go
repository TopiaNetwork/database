package tpfs

import (
	// "encoding/json"
	"fmt"
	// "io"
	// "io/ioutil"
	"os"
	// "strings"

)
func CreateAction(ct CryptoType, ca CryptoAlgorithm, recipients []string) error {
// func CreateAction(ct CryptoType, ca CryptoAlgorithm) error {
	// fmt.Println("asdf")
	config, err := config.Load("")
	if err != nil {
		fmt.Errorf("unable to load configuration; reason: %s", err.Error())
	}

	secretStore := NewSecretStore()
	vault := Vault{
		CryptoType:      ct,
		CryptoAlgorithm: ca,
	}

	if vault.CryptoType == SYMMETRIC_ENCRYPTION {
		passphrase, err := GetPassphrase(config)
		if err != nil {
			if !AskPassphraseFlagCheck() {
				AskPassphrase(true)
			}
		} else {
			SetPassphrase(passphrase)
		}
	}

	err = vault.Lock(config, secretStore)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(InferStorePath(config), os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	err = vault.Dump(f)
	if err != nil {
		return err
	}

	return nil
}

func GetAction(key string, filepath string) error {
	// config, err := config.Load("")
	// if err != nil {
	// 	fmt.Errorf("unable to load configuration; reason: %s", err.Error())
	// }

	// vault, err := OpenVault(InferStorePath(config))
	// if err != nil {
	// 	return err
	// }

	// secretStore, err := vault.Unlock(config)
	// if err != nil {
	// 	return err
	// }

	// value, found := secretStore.Data[key]
	// if !found {
	// 	return fmt.Errorf("key %s not found", key)
	// }

	// // If the --file flag is provided
	// if filepath != "" {
	// 	err := ioutil.WriteFile(filepath, []byte(value), os.FileMode(0600))
	// 	if err != nil {
	// 		return err
	// 	}
	// } else {
	// 	// Use fmt.Print to support patch processes,
	// 	// to get the value "as is" without any appended newlines
	// 	if isPipe(os.Stdout) {
	// 		fmt.Print(value)
	// 	} else {
	// 		InfoLogger.Println(value)
	// 	}
	// }

	return nil
}

func SetAction(key, value, file string) error {
	// config, err := config.Load("")
	// if err != nil {
	// 	fmt.Errorf("unable to load configuration; reason: %s", err.Error())
	// }

	// // If the --file flag is provided
	// if file != "" {
	// 	// And the file actually exists on file system
	// 	if PathExists(file) {
	// 		// Then load it's content
	// 		fileContent, err := ioutil.ReadFile(file)
	// 		if err != nil {
	// 			return err
	// 		}

	// 		value = string(fileContent)
	// 	} else {
	// 		ErrorLogger.Fatalf("Cannot open %s because it doesn't exist", file)
	// 	}
	// }

	// vault, err := OpenVault(InferStorePath(config))
	// if err != nil {
	// 	return err
	// }

	// secretStore, err := vault.Unlock(config)
	// if err != nil {
	// 	return err
	// }
	// secretStore.Data[key] = value

	// err = vault.Lock(config, secretStore)
	// if err != nil {
	// 	return err
	// }

	// f, err := os.Open(InferStorePath(config))
	// if err != nil {
	// 	return err
	// }

	// err = vault.Dump(f)
	// if err != nil {
	// 	return err
	// }

	return nil
}
