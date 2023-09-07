windows:export GOOS=windows
windows:export GOARCH=amd64
windows:export CGO_ENABLED=0

linux:export GOOS=linux
linux:export GOARCH=amd64
linux:export GODEBUG=cgocheck=0

windows:
	go build -o build/cron_ftp.exe main.go

linux:
	go build -o build/cron_ftp main.go