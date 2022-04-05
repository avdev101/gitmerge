build:
	rm -rf dist
	mkdir dist
	go build -o dist ./cmd/gitmerge

build-linux:
	rm -rf dist
	mkdir dist
	GOOS=linux GOARCH=386 go build -o dist ./cmd/gitmerge
