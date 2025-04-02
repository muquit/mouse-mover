# Introduction

`mouse-mover.exe` is a simple command line tool for Microsoft Windows to 
move mouse to any corner when the system is idle for x seconds. This can
be used to prevent system to go to sleep without changing power plan.

# Synopsis
```
mouse-mover.exe v1.0.1
Compiled with go version: go1.22.5
A program to move mouse to any corner if the system is idle for x seconds
Usage mouse-mover.exe [options]
Where the options are:
  -corner string
        Corner to move mouse to (ulc, urc, blc, brc) (default "ulc")
  -idle int
        Idle time in seconds before moving mouse (default 60)
  -tick int
        Print idle time every X seconds (0 to disable)
  -version
        Show version info
```

# Download

Download pre-compiled binaries from
[Releases](https://github.com/muquit/mouse-mover/releases) page


# Building from source
Install [Go](https://go.dev/) first

```
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" .
or
make build
```
# License

MIT License - See LICENSE.txt file for details.

# Authors

Developed with Claude AI 3.7 Sonnet, working under my guidance and instructions
