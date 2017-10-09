mac:
	gox -osarch="darwin/amd64"

linux:
	gox -osarch="linux/amd64"

run:
	lywatch --cmd="go run *.go" --port="8965"

pro:
	nohup /root/sun/sun_linux_amd64 &
