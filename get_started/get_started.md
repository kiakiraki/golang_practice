# Get Started Golang

## Install Golang

### Windows

1. Download and Install https://dl.google.com/go/go1.13.4.windows-amd64.msi
2. Add user `%PATH%` to `%GOPATH%\bin`
3. Install `gore` (Golang REPL) and some packages
   1. `go get github.com/motemen/gore/cmd/gore`
   2. `go get github.com/mdempsky/gocode`
   3. `go get github.com/motemen/ghq`
      * 以下のようなエラーを吐いたので、`v0.12.8` をcheckoutして手動で `go install` した
      ```
        # github.com/motemen/ghq
        src\github.com\motemen\ghq\commands.go:26:14: cannot use cli.BoolFlag literal (type cli.BoolFlag) as type cli.Flag in array or slice literal:
                cli.BoolFlag does not implement cli.Flag (Apply method has pointer receiver)
        src\github.com\motemen\ghq\commands.go:27:14: cannot use cli.BoolFlag literal (type cli.BoolFlag) as type cli.Flag in array or slice literal:
                cli.BoolFlag does not implement cli.Flag (Apply method has pointer receiver)
        src\github.com\motemen\ghq\commands.go:28:14: cannot use cli.BoolFlag literal (type cli.BoolFlag) as type cli.Flag in array or slice literal:
                cli.BoolFlag does not implement cli.Flag (Apply method has pointer receiver)
                cli.BoolFlag does not implement cli.Flag (Apply method has pointer receiver)
        src\github.com\motemen\ghq\commands.go:30:16: cannot use cli.StringFlag literal (type cli.StringFlag) as type cli.Flag in array or slice literal:
                cli.StringFlag does not implement cli.Flag (Apply method has pointer receiver)
        src\github.com\motemen\ghq\commands.go:31:14: cannot use cli.BoolFlag literal (type cli.BoolFlag) as type cli.Flag in array or slice literal:
                cli.BoolFlag does not implement cli.Flag (Apply method has pointer receiver)
        src\github.com\motemen\ghq\commands.go:59:15: cannot use cli.BoolFlag literal (type cli.BoolFlag) as type cli.Flag in array or slice literal:
                cli.BoolFlag does not implement cli.Flag (Apply method has pointer receiver)
        src\github.com\motemen\ghq\commands.go:60:17: cannot use cli.StringFlag literal (type cli.StringFlag) as type cli.Flag in array or slice literal:
                cli.StringFlag does not implement cli.Flag (Apply method has pointer receiver)
        src\github.com\motemen\ghq\commands.go:61:15: cannot use cli.BoolFlag literal (type cli.BoolFlag) as type cli.Flag in array or slice literal:
                cli.BoolFlag does not implement cli.Flag (Apply method has pointer receiver)
        src\github.com\motemen\ghq\commands.go:62:15: cannot use cli.BoolFlag literal (type cli.BoolFlag) as type cli.Flag in array or slice literal:
      ```
   4. Install `peco`
       * Download and add to `%PATH%`  https://github.com/peco/peco/releases/download/v0.5.3/peco_windows_amd64.zip

### Ubuntu

[公式Wiki](https://github.com/golang/go/wiki/Ubuntu) を参照

```text
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```
