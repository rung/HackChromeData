package decrypter

import (
	"errors"
	"syscall"
	"unsafe"
)

func Chromium(key, encryptPass []byte) ([]byte, error) {
	// Ensure the encrypted data has at least 3 bytes for the version prefix.
	if len(encryptPass) < 3 {
		return nil, errors.New("encrypted data too short")
	}

	switch string(encryptPass[:3]) {
	case "v20":
		// For v20:
		// Remove the "v20" prefix. The next 12 bytes are the IV,
		// and the remainder is the ciphertext (with tag).
		if len(encryptPass) < 3+12 {
			return nil, errors.New("invalid v20 data")
		}
		iv := encryptPass[3:15]
		ciphertext := encryptPass[15:]
		decrypted, err := aesGCMDecrypt(ciphertext, key, iv)
		if err != nil {
			return nil, err
		}
		// Remove the first 32 bytes of the decrypted data (internal header).
		if len(decrypted) < 32 {
			return nil, errors.New("decrypted data too short")
		}
		return decrypted[32:], nil
	default:
		// For non-v20 (e.g., v10):
		// Remove the first 3 bytes as the prefix, then the next 12 bytes as the nonce,
		// and the remainder as the ciphertext (with tag).
		if len(encryptPass) <= 15 {
			return nil, errors.New("password is empty")
		}
		return aesGCMDecrypt(encryptPass[15:], key, encryptPass[3:15])
	}
}

type dataBlob struct {
	cbData uint32
	pbData *byte
}

func NewBlob(d []byte) *dataBlob {
	if len(d) == 0 {
		return &dataBlob{}
	}
	return &dataBlob{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

func (b *dataBlob) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

// DPApi
// chrome < 80 https://chromium.googlesource.com/chromium/src/+/76f496a7235c3432983421402951d73905c8be96/components/os_crypt/os_crypt_win.cc#82
func DPApi(data []byte) ([]byte, error) {
	dllCrypt := syscall.NewLazyDLL("Crypt32.dll")
	dllKernel := syscall.NewLazyDLL("Kernel32.dll")
	procDecryptData := dllCrypt.NewProc("CryptUnprotectData")
	procLocalFree := dllKernel.NewProc("LocalFree")
	var outBlob dataBlob
	r, _, err := procDecryptData.Call(uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&outBlob)))
	if r == 0 {
		return nil, err
	}
	defer procLocalFree.Call(uintptr(unsafe.Pointer(outBlob.pbData)))
	return outBlob.ToByteArray(), nil
}
