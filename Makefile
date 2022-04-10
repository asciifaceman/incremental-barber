VERSION=$(shell cat VERSION)


.PHONY: build_all build_mac build_linux bundle

build_all: build_linux
	zip target/ib-linux-$(VERSION).zip target/ib-linux
	zip target/ib-mac-$(VERSION).zip target/ib-mac

build_mac:
	GOOS=darwin GOARCH=amd64 go build -o target/ib-mac -ldflags="-X 'main.version=$(VERSION)'"

build_linux:
	GOOS=linux GOARCH=amd64 go build -o target/ib-linux -ldflags="-X 'main.version=$(VERSION)'"

bundle:
	cd pkg/bundles && fyne bundle --pkg bundles --prefix Tools tools/ > tools.go
	cd pkg/bundles && fyne bundle --pkg bundles --prefix App app/ > app.go