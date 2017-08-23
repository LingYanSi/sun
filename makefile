mac:
	gox -osarch="darwin/amd64"

linux:
	gox -osarch="linux/amd64"

run:
	go run *.go
