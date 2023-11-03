package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"gopkg.in/ini.v1"
	"io"
	"os"
)

func EncryptFile(inputPath string, outputPath string, key []byte) error {
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
	fmt.Println(block)

	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)
	writer := &cipher.StreamWriter{S: stream, W: outputFile}

	_, err = io.Copy(writer, inputFile)
	if err != nil {
		return err
	}

	keyini, err := ini.Load("key.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	keyini.Section("").Key("endkey").SetValue(string(iv))
	keyini.SaveTo("key.ini")

	return nil
}
