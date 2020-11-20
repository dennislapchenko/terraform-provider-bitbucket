VER=`cat ./VERSION`

test:
	go test ./...

build: test
	go build -o terraform-provider-bitbucket_$(VER)
