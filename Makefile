# https://github.com/golang/go/issues/50745
test:
	go list -f '{{.Dir}}' -m | xargs go test -short

long.test:
	go list -f '{{.Dir}}' -m | xargs go test -race

tools:
	# pkgsiteを実行すれば、例えば、http://localhost:8080/github.com/qawatake/go-sandbox/recoveryでpackageのページを表示できる。
	go install golang.org/x/pkgsite/cmd/pkgsite@latest
