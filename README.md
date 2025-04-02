## Table Of Contents
- [Introduction](#introduction)
- [Synopsis](#synopsis)
- [Download](#download)
- [Building from source](#building-from-source)
- [License](#license)
- [Authors](#authors)

# Introduction

`mouse-mover.exe` is a simple command line tool for Microsoft Windows to 
move mouse to any corner when the system is idle for x seconds. This can
be used to prevent system to go to sleep without changing power plan.

# Synopsis
```
   Usage of mouse-mover.exe:
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

---
<sub>TOC is created by https://github.com/muquit/markdown-toc-go on Apr-02-2025</sub>
