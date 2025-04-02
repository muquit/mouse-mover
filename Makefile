#====================================================================
# Reqquires https://github.com/muquit/go-xbuild-go for cross compiling
# for other platforms.
# Mar-29-2025 muquit@muquit.com 
#====================================================================
README_ORIG=./docs/README.md
README=./README.md
BINARY=./mouse-mover.exe
GEN_TOC_PROG=markdown-toc-go

all: build build_all doc

build:
	@echo "*** Compiling $(BINARY) ..."
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o $(BINARY)

build_all:
	@echo "*** Cross Compiling $(BINARY)...."
	go-xbuild-go -pi=false

doc:
	@echo "*** Generating README.md with TOC ..."
	@if [ -f $(README) ]; then chmod 600 $(README); fi
	$(GEN_TOC_PROG) -i $(README_ORIG) -o $(README) -f
	chmod 444 $(README)

clean:
	/bin/rm -f $(BINARY)
	/bin/rm -rf ./bin
