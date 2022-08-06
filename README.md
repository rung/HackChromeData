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

#### Usage
- For Windows
```bash
# Cookie
hack-chrome-data.exe -kind cookie -targetpath "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Default\Network\Cookies" -localstate "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Local State"

# Password
hack-chrome-data.exe -kind logindata -targetpath "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Default\Login Data" -localstate "%HOMEPATH%\AppData\Local\Google\Chrome\User Data\Local State"
```

- For macOS
````bash

````
