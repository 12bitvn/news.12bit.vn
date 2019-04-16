hugo:
	hugo --gc --minify
build: hugo
	mkdir -p functions
	go get ./...
	go build -o functions/hello-lambda ./...
