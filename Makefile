validate:
	swagger validate node-api.yaml
	swagger validate public-api.yaml

client: validate
	swagger generate client --default-scheme https -f node-api.yaml -A api -t ./node
	swagger generate client --default-scheme https -f public-api.yaml -A api -t ./public

all: client
