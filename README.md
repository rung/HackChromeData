# HackChromeData for Seccamp2022
#### What's this
- Simpler version of [HackBrowserData](https://github.com/moonD4rk/HackBrowserData) for training
  - Original source code came from [HackBrowserData](https://github.com/moonD4rk/HackBrowserData)

#### For Windows (tbd)
- Compire
```shell
brew install mingw-w64

CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build
```

- Usage

#### For Mac (tbd)
- Compile
````shell
brew install FiloSottile/musl-cross/musl-cross

CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static"
````
