build:
	rm -rf dist
	mkdir dist
	go build -o dist ./cmd/gitmerge

build-linux:
	rm -rf dist
	mkdir dist
	GOOS=linux GOARCH=386 go build -o dist ./cmd/gitmerge

upload:
	scp dist/gitmerge root@45.131.40.57:/root/gitmerge
