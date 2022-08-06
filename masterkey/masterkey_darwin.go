package masterkey

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"golang.org/x/crypto/pbkdf2"
	"os/exec"
	"strings"
)

var (
	ErrWrongSecurityCommand   = errors.New("macOS wrong security command")
	ErrCouldNotFindInKeychain = errors.New("macOS could not find in keychain")
)

func GetMasterKey() ([]byte, error) {
	var (
		cmd            *exec.Cmd
		stdout, stderr bytes.Buffer
	)

	// defer os.Remove(item.TempChromiumKey)
	// Get the master key from the keychain
	// $ security find-generic-password -wa 'Chrome'
	cmd = exec.Command("security", "find-generic-password", "-wa", "Chrome Safe Storage")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	if stderr.Len() > 0 {
		if strings.Contains(stderr.String(), "could not be found") {
			return nil, ErrCouldNotFindInKeychain
		}
		return nil, errors.New(stderr.String())
	}
	chromeSecret := bytes.TrimSpace(stdout.Bytes())
	if chromeSecret == nil {
		return nil, ErrWrongSecurityCommand
	}
	chromeSalt := []byte("saltysalt")
	// @https://source.chromium.org/chromium/chromium/src/+/master:components/os_crypt/os_crypt_mac.mm;l=157
	key := pbkdf2.Key(chromeSecret, chromeSalt, 1003, 16, sha1.New)
	if key == nil {
		return nil, ErrWrongSecurityCommand
	}
	return key, nil
}
