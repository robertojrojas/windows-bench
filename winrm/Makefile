SOURCES := $(shell find . -name '*.go')
BINARY := WinRmPOC.exe
TARGET_OS := windows

build: $(BINARY)

$(BINARY): $(SOURCES)
	GOOS=$(TARGET_OS) GO111MODULE=on CGO_ENABLED=0  go build  -o $(BINARY) $(SOURCES)

tests:
	echo 'write tests....'

