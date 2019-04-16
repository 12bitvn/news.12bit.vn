build:
	cd ./functions && \
	go get ./... &&\
	go build -o crawl-news main.go
