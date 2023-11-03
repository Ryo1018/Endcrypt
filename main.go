package main

import (
	"Endcrypt/decrypt"
	"Endcrypt/encrypt"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func main() {
	keyini, err := ini.Load("key.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	//key := []byte("64a7bb2bb8c86b8cc52079b37236af67")
	key := keyini.Section("").Key("key").String()
	fmt.Println(key)

	inputFilePath := "input.txt"
	encryptedFilePath := "encrypted.bin"
	decryptedFilePath := "decrypted.txt"

	// ファイルの暗号化
	err = encrypt.EncryptFile(inputFilePath, encryptedFilePath, []byte(key))
	if err != nil {
		log.Fatalf("ファイルの暗号化中にエラーが発生しました：%v", err)
	}

	err = decrypt.DecryptFile(encryptedFilePath, decryptedFilePath, []byte(key))
	if err != nil {
		log.Fatalf("ファイルの復号化中にエラーが発生しました：%v", err)
	}
}
