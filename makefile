build:
	hugo  --gc --minify --buildFuture --enableGitInfo && \
	go get ./... &&\
	go build -o ./bin/crawl-news main.go
