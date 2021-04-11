validate:
	swagger validate node-api.yaml
	swagger validate public-api.yaml
	swagger validate sd-api.yaml

client: validate
	swagger generate client --default-scheme https -f node-api.yaml -A api -t ./node
	swagger generate client --default-scheme https -f public-api.yaml -A api -t ./public
	swagger generate client --default-scheme https -f sd-api.yaml -A api -t ./sd

all: client
