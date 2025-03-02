# HackChromeData for Development Environment Security Training (Security Camp 2022)
## What's this
- The training material is here: https://github.com/rung/training-devenv-security. This repository is a part of the training.
- This software decrypt your Chrome's cookie and password, then send them to standard output.
  - This software **does not** upload any credential to the internet.

### Referenced source code
- Original source code is [HackBrowserData](https://github.com/moonD4rk/HackBrowserData).
- This repository contains the necessary part only for PoC.

## Disclaimer
- This tool is limited to education and security research only

## Build
- It uses github.com/crazy-max/xgo to build cgo binary on cross environment.
```bash
make build
```

## Supported OS and Architecture
- Windows x64
- macOS x64
- macOS ARM64

## Usage
- For Windows
  - (When your profile name is `Default`)
  - (For Powershell user) Please replace `%HOMEPATH%` to `$HOME`

  - For Chrome versions before v127 (before application-bound credentials):
```bash
# Cookie
hack-chrome-data.exe -kind cookie -targetpath "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Default\Network\Cookies" -localstate "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Local State"

# Password
hack-chrome-data.exe -kind logindata -targetpath "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Default\Login Data" -localstate "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Local State"
```

- For Chrome v127+ (with application-bound credentials):
  - You need to extract the master key using [Chrome-App-Bound-Encryption-Decryption](https://github.com/xaitax/Chrome-App-Bound-Encryption-Decryption)
  - Then use the extracted key with the `-sessionstorage` option:
```bash
# Cookie
hack-chrome-data.exe -kind cookie -targetpath "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Default\Network\Cookies" -sessionstorage <your-extracted-master-key>

# Password　(This program doesn't support Password extraction atm)
```

- For macOS (Normal)
  - (When your profile name is `Default`)
  - HackChromeData asks to access keychain
    - (`security find-generic-password -wa "Chrome"` is called internally)
```bash
# Cookie
$ ./hack-chrome-data -kind cookie -targetpath ~/Library/Application\ Support/Google/Chrome/Default/Cookies

# Password
$ ./hack-chrome-data -kind logindata -targetpath ~/Library/Application\ Support/Google/Chrome/Default/Login\ Data
````

- For macOS (Use Keychain Value)
  - (When your profile name is `Default`)
  1. Get `Chrome Sesssion Storage` value on Keychain
      - `security find-generic-password -wa "Chrome"`
      - or you can get the value through forensic tool like [chainbreaker](https://github.com/n0fate/chainbreaker).
  2. Decrypt cookies and passwords 
```
# Cookie
$ ./hack-chrome-data -kind cookie -targetpath ~/Library/Application\ Support/Google/Chrome/Default/Cookies -sessionstorage <session storage value>

# Password
$ ./hack-chrome-data -kind logindata -targetpath ~/Library/Application\ Support/Google/Chrome/Default/Login\ Data -sessionstorage <session storage value>
```