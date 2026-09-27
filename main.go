package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("shit is is getting started")

	allArgs := os.Args

	singleArg := allArgs[1]

	fmt.Println("the args is ", singleArg)

	result, err := os.ReadDir(singleArg)

	if err != nil {
		fmt.Println("error occured while reading", err)
		return
	}

	for _, val := range result {
		// if err != nil {
		// 	fmt.Println("err", err)
		// 	return
		// }

		if val.Type().IsRegular() {

			hash, err := fileHash(val.Name())
			if err != nil {
				fmt.Println("error in hash", err)
				return
			}

			fmt.Println("hash of the file is ", hash)
		}

		fmt.Println("the result is ", val.Type())
		fmt.Println("res", val.Name())
	}
}

func fileHash(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
