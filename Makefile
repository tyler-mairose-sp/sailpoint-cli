clean:
	go clean ./...

mocks:
	# Ref: https://github.com/golang/mock
	mockgen -source=client/client.go -destination=mocks/client.go -package=mocks

test:
	go test -v -count=1 ./...

install:
	go build -o /usr/local/bin/sail -buildvcs=false

vhs:
	find assets -name "*.tape" | xargs -n 1 vhs

.PHONY: clean mocks test install vhs .docker/login .docker/build .docker/push
