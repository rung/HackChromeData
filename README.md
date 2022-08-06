# HackChromeData for Seccamp2022
#### What's this
- Simpler version of [HackBrowserData](https://github.com/moonD4rk/HackBrowserData) for training
  - Original source code came from [HackBrowserData](https://github.com/moonD4rk/HackBrowserData)
- It's for Windows and maxOS

#### Build
- It uses github.com/crazy-max/xgo to build cgo binary on cross environment.
```bash
make build
```

#### Supported OS and Architecture
- Windows x64
- macOS x64
- macOS ARM64

#### Usage
- For Windows
```bash
# Cookie
hack-chrome-data.exe -kind cookie -targetpath "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Default\Network\Cookies" -localstate "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Local State"

# Password
hack-chrome-data.exe -kind logindata -targetpath "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Default\Login Data" -localstate "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Local State"
```

- For macOS (Normal)
````bash
# Cookie
$ ./hack-chrome-data -kind cookie -targetpath ~/Library/Application\ Support/Google/Chrome/Default/Cookies

# Password
$ ./hack-chrome-data -kind logindata -targetpath ~/Library/Application\ Support/Google/Chrome/Default/Login\ Data

````

- For macOS (Use Keychain Value)
```bash
# Get Chrome Session Storage's value
$ security find-generic-password -wa "Chrome"

# Cookie
$ ./hack-chrome-data -kind cookie -targetpath ~/Library/Application\ Support/Google/Chrome/Default/Cookies -sessionstorage <session storage value>

# Password
$ ./hack-chrome-data -kind logindata -targetpath ~/Library/Application\ Support/Google/Chrome/Default/Login\ Data -sessionstorage <session storage value>

```
