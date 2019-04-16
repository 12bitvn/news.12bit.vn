build:
	hugo  --gc --minify --buildFuture --enableGitInfo && \
	cd ./functions && \
	go mod download &&\
	go build -o ../bin/crawl-news main.go
