build:
	hugo  --gc --minify --buildFuture --enableGitInfo && \
	cd ./functions && \
	go get ./... &&\
	go build -o crawl-news main.go
