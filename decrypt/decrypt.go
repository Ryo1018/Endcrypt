package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"gopkg.in/ini.v1"
	"io"
	"os"
)

func DecryptFile(inputPath string, outputPath string, key []byte) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	keyini, err := ini.Load("key.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	//key := []byte("64a7bb2bb8c86b8cc52079b37236af67")
	iv := []byte(keyini.Section("").Key("endkey").String())
	_, err = io.ReadFull(inputFile, iv)
	if err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)
	reader := &cipher.StreamReader{S: stream, R: inputFile}

	_, err = io.Copy(outputFile, reader)
	if err != nil {
		return err
	}

	return nil
}
