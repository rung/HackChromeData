package decrypter

import (
	"errors"
	"syscall"
	"unsafe"
)

func Chromium(key, encryptPass []byte) ([]byte, error) {
	if len(encryptPass) > 15 {
		// remove Prefix 'v10'
		return aesGCMDecrypt(encryptPass[15:], key, encryptPass[3:15])
	} else {
		return nil, errors.New("password is empty")
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
