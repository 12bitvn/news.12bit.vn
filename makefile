build:
	rm -rf ./public && hugo  --gc --minify --buildFuture --enableGitInfo

preview: build
	netlify deploy

deploy: build
	netlify deploy --prod
