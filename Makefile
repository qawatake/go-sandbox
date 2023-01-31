# https://github.com/golang/go/issues/50745
test:
	go list -f '{{.Dir}}' -m | xargs go test -short

long.test:
	go list -f '{{.Dir}}' -m | xargs go test -race
