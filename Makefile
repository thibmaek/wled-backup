.DEFAULT_GOAL := build

build_x64:
	GOOS=darwin go build -o bin/<app>_mac_x64 .
	GOOS=linux go build -o bin/<app>_linux_x64 .
	GOOS=windows go build -o bin/<app>_win_x64.exe .

build_arm:
	GOOS=linux GOARCH=arm GOARM=6 go build -o bin/<app>_linux_armv6 .
	GOOS=linux GOARCH=arm GOARM=7 go build -o bin/<app>_linux_armv7 .
	GOOS=darwin GOARCH=arm64 go build -o bin/<app>_mac_arm64 .

build: build_x64 build_arm
